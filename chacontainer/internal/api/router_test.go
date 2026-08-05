package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chacontainer/backend/internal/config"
	"github.com/chacontainer/backend/internal/store/sqlite"
)

func newTestRouter(t *testing.T) http.Handler {
	t.Helper()
	db, err := sqlite.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	cfg := &config.Config{JWTSecret: "test-secret", MakeWebhookSecret: ""}
	return NewRouter(cfg, db)
}

func doJSON(t *testing.T, router http.Handler, method, path, token string, body interface{}) (*httptest.ResponseRecorder, map[string]interface{}) {
	t.Helper()
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("encode body: %v", err)
		}
	}
	req := httptest.NewRequest(method, path, &buf)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	var out map[string]interface{}
	if rec.Body.Len() > 0 {
		_ = json.Unmarshal(rec.Body.Bytes(), &out)
	}
	return rec, out
}

func registerTenant(t *testing.T, router http.Handler, company, email, password string) (string, map[string]interface{}) {
	t.Helper()
	rec, out := doJSON(t, router, http.MethodPost, "/api/v1/auth/register", "", map[string]string{
		"company_name": company,
		"name":         "Admin " + company,
		"email":        email,
		"password":     password,
	})
	if rec.Code != http.StatusCreated {
		t.Fatalf("register %s: status %d, body %v", email, rec.Code, out)
	}
	return out["token"].(string), out
}

func TestRegisterAndLogin(t *testing.T) {
	router := newTestRouter(t)

	// Successful registration issues a usable token.
	token, out := registerTenant(t, router, "Acme SA", "admin@acme.mx", "Password123")
	if token == "" {
		t.Fatal("expected non-empty token")
	}
	if out["user"] == nil || out["tenant"] == nil {
		t.Fatalf("expected user and tenant in response, got %v", out)
	}

	// Password too short is rejected.
	rec, _ := doJSON(t, router, http.MethodPost, "/api/v1/auth/register", "", map[string]string{
		"company_name": "X", "name": "Y", "email": "short@x.com", "password": "short",
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for short password, got %d", rec.Code)
	}

	// Duplicate email is rejected.
	rec, _ = doJSON(t, router, http.MethodPost, "/api/v1/auth/register", "", map[string]string{
		"company_name": "Other", "name": "Other Admin", "email": "admin@acme.mx", "password": "Password123",
	})
	if rec.Code != http.StatusConflict {
		t.Fatalf("expected 409 for duplicate email, got %d", rec.Code)
	}

	// Correct login succeeds.
	rec, loginOut := doJSON(t, router, http.MethodPost, "/api/v1/auth/login", "", map[string]string{
		"email": "admin@acme.mx", "password": "Password123",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 login, got %d: %v", rec.Code, loginOut)
	}
	if loginOut["token"] == "" || loginOut["token"] == nil {
		t.Fatal("expected token on login")
	}

	// Wrong password is rejected.
	rec, _ = doJSON(t, router, http.MethodPost, "/api/v1/auth/login", "", map[string]string{
		"email": "admin@acme.mx", "password": "wrong-password",
	})
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for wrong password, got %d", rec.Code)
	}

	// Protected routes reject missing/invalid tokens.
	rec, _ = doJSON(t, router, http.MethodGet, "/api/v1/assets", "", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 without token, got %d", rec.Code)
	}
	rec, _ = doJSON(t, router, http.MethodGet, "/api/v1/assets", "not-a-real-token", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 with bogus token, got %d", rec.Code)
	}
}

func TestTenantIsolation(t *testing.T) {
	router := newTestRouter(t)

	tokenA, _ := registerTenant(t, router, "Tenant A", "a@tenanta.mx", "PasswordA1")
	tokenB, _ := registerTenant(t, router, "Tenant B", "b@tenantb.mx", "PasswordB1")

	plantRec, plantOut := doJSON(t, router, http.MethodPost, "/api/v1/plants", tokenA, map[string]string{
		"code": "PLT-A", "name": "Planta A",
	})
	if plantRec.Code != http.StatusCreated {
		t.Fatalf("create plant: status %d body %v", plantRec.Code, plantOut)
	}
	plantID := plantOut["id"].(string)

	assetRec, assetOut := doJSON(t, router, http.MethodPost, "/api/v1/assets", tokenA, map[string]interface{}{
		"plant_id": plantID, "name": "Contenedor A", "code": "AST-A1", "type": "container",
	})
	if assetRec.Code != http.StatusCreated {
		t.Fatalf("create asset: status %d body %v", assetRec.Code, assetOut)
	}
	assetID := assetOut["id"].(string)

	// Tenant A sees its own asset.
	_, listA := doJSON(t, router, http.MethodGet, "/api/v1/assets", tokenA, nil)
	if int(listA["total"].(float64)) != 1 {
		t.Fatalf("expected tenant A to see 1 asset, got %v", listA["total"])
	}

	// Tenant B sees nothing.
	_, listB := doJSON(t, router, http.MethodGet, "/api/v1/assets", tokenB, nil)
	if int(listB["total"].(float64)) != 0 {
		t.Fatalf("expected tenant B to see 0 assets, got %v", listB["total"])
	}

	// Tenant B cannot fetch tenant A's asset by ID.
	rec, _ := doJSON(t, router, http.MethodGet, "/api/v1/assets/"+assetID, tokenB, nil)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for cross-tenant asset access, got %d", rec.Code)
	}
}

func TestAssetLifecycleAndEvents(t *testing.T) {
	router := newTestRouter(t)
	token, _ := registerTenant(t, router, "Logistics Co", "ops@logistics.mx", "Password123")

	_, plantOut := doJSON(t, router, http.MethodPost, "/api/v1/plants", token, map[string]string{"code": "PLT-1", "name": "Planta Uno"})
	plantID := plantOut["id"].(string)

	rec, assetOut := doJSON(t, router, http.MethodPost, "/api/v1/assets", token, map[string]interface{}{
		"plant_id": plantID, "name": "Pallet Euro", "code": "AST-100", "type": "pallet",
	})
	if rec.Code != http.StatusCreated {
		t.Fatalf("create asset: %d %v", rec.Code, assetOut)
	}
	assetID := assetOut["id"].(string)
	qrCode := assetOut["qr_code"].(string)
	if assetOut["status"] != "available" {
		t.Fatalf("expected default status available, got %v", assetOut["status"])
	}

	// Creating an asset must record a check_in event (regression guard for the
	// nested-route bug where GET /assets/{id}/events used to 404/misparse id).
	rec, eventsOut := doJSON(t, router, http.MethodGet, "/api/v1/assets/"+assetID+"/events", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get events: status %d body %v", rec.Code, eventsOut)
	}
	events := eventsOut["data"].([]interface{})
	if len(events) != 1 {
		t.Fatalf("expected 1 event after create, got %d", len(events))
	}

	// Scan by QR code updates status/location and appends an event.
	rec, scanOut := doJSON(t, router, http.MethodPost, "/api/v1/assets/scan", token, map[string]interface{}{
		"qr_code": qrCode, "action": "check_out", "plant_id": plantID,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("scan: status %d body %v", rec.Code, scanOut)
	}

	_, eventsOut2 := doJSON(t, router, http.MethodGet, "/api/v1/assets/"+assetID+"/events", token, nil)
	events2 := eventsOut2["data"].([]interface{})
	if len(events2) != 2 {
		t.Fatalf("expected 2 events after scan, got %d", len(events2))
	}

	// Scanning an unknown QR code 404s.
	rec, _ = doJSON(t, router, http.MethodPost, "/api/v1/assets/scan", token, map[string]interface{}{
		"qr_code": "does-not-exist", "action": "check_out",
	})
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for unknown QR, got %d", rec.Code)
	}
}

func TestShipmentLifecycle(t *testing.T) {
	router := newTestRouter(t)
	token, _ := registerTenant(t, router, "Freight Co", "ops@freight.mx", "Password123")

	_, plantOut := doJSON(t, router, http.MethodPost, "/api/v1/plants", token, map[string]string{"code": "PLT-1", "name": "Origen"})
	plantID := plantOut["id"].(string)

	_, clientOut := doJSON(t, router, http.MethodPost, "/api/v1/clients", token, map[string]string{"code": "CLI-1", "name": "Cliente Uno"})
	clientID := clientOut["id"].(string)

	rec, shipOut := doJSON(t, router, http.MethodPost, "/api/v1/shipments", token, map[string]interface{}{
		"client_id": clientID, "origin_plant_id": plantID,
	})
	if rec.Code != http.StatusCreated {
		t.Fatalf("create shipment: %d %v", rec.Code, shipOut)
	}
	if shipOut["status"] != "draft" {
		t.Fatalf("expected new shipment status draft, got %v", shipOut["status"])
	}
	shipID := shipOut["id"].(string)

	// Nested-route regression guard: POST /shipments/{id}/transition must
	// resolve {id} correctly, not the literal "transition" segment.
	rec, transOut := doJSON(t, router, http.MethodPost, "/api/v1/shipments/"+shipID+"/transition", token, map[string]string{
		"status": "in_transit", "description": "salió de planta",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("transition: status %d body %v", rec.Code, transOut)
	}
	if transOut["status"] != "in_transit" {
		t.Fatalf("expected status in_transit, got %v", transOut["status"])
	}

	rec, timelineOut := doJSON(t, router, http.MethodGet, "/api/v1/shipments/"+shipID+"/timeline", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("timeline: status %d body %v", rec.Code, timelineOut)
	}
	timeline := timelineOut["data"].([]interface{})
	if len(timeline) != 2 { // created + transitioned
		t.Fatalf("expected 2 timeline events, got %d", len(timeline))
	}
}

// TestCORSPreflight guards against a real bug found while testing the web
// app in a browser: JSON POST bodies and the Authorization header both force
// a CORS preflight OPTIONS request, which needs Access-Control-Allow-Origin
// on both the public auth routes and every authenticated /api/v1/* route.
func TestCORSPreflight(t *testing.T) {
	router := newTestRouter(t)

	for _, path := range []string{"/api/v1/auth/login", "/api/v1/auth/register", "/api/v1/assets", "/api/v1/shipments/abc/transition"} {
		req := httptest.NewRequest(http.MethodOptions, path, nil)
		req.Header.Set("Origin", "http://localhost:5173")
		req.Header.Set("Access-Control-Request-Method", "POST")
		req.Header.Set("Access-Control-Request-Headers", "content-type,authorization")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusNoContent {
			t.Fatalf("OPTIONS %s: expected 204, got %d", path, rec.Code)
		}
		if rec.Header().Get("Access-Control-Allow-Origin") == "" {
			t.Fatalf("OPTIONS %s: missing Access-Control-Allow-Origin header", path)
		}
	}
}

func TestHealth(t *testing.T) {
	router := newTestRouter(t)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
