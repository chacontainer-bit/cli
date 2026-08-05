package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
	"github.com/cli/cli/v2/chacontainer/internal/config"
	"github.com/cli/cli/v2/chacontainer/internal/store/postgres"
)

// NewRouter wires all routes against real PostgreSQL-backed stores, so the
// whole API runs with genuine persistence on a single local machine. The
// only network calls this server ever makes are to the database given in
// cfg.DatabaseURL - Airtable/ERP/Make.com are only contacted if their own
// client code is explicitly invoked, which nothing here does. Webhook
// receivers stay wired to a no-op processor: this server accepts inbound
// webhooks but never calls out to those third-party services itself.
func NewRouter(cfg *config.Config, db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "chacontainer"})
	})

	assetStore := postgres.NewAssetStore(db)
	shipmentStore := postgres.NewShipmentStore(db)
	clientStore := postgres.NewClientStore(db)
	plantStore := postgres.NewPlantStore(db)
	statsStore := postgres.NewStatsStore(db)
	userStore := postgres.NewUserStore(db)
	webhookProcessor := &stubWebhookProcessor{}

	assetsH := handlers.NewAssetsHandler(assetStore)
	shipmentsH := handlers.NewShipmentsHandler(shipmentStore)
	clientsH := handlers.NewClientsHandler(clientStore)
	plantsH := handlers.NewPlantsHandler(plantStore)
	dashboardH := handlers.NewDashboardHandler(statsStore)
	webhooksH := handlers.NewWebhooksHandler(webhookProcessor, cfg.MakeWebhookSecret)
	authH := handlers.NewAuthHandler(userStore, cfg.JWTSecret)

	// Public webhook receivers: called server-to-server, never from a
	// browser, so they don't need CORS.
	mux.HandleFunc("POST /webhooks/make", webhooksH.Make)
	mux.HandleFunc("POST /webhooks/erp", webhooksH.ERP)
	mux.HandleFunc("POST /webhooks/airtable", webhooksH.Airtable)

	// Login is public (no JWT yet - that's the whole point) but is still
	// called from the browser, so it needs the same CORS header handling as
	// every authenticated route below.
	public := chain(middleware.Logger, middleware.CORS)
	mux.Handle("POST /api/v1/auth/login", public(http.HandlerFunc(authH.Login)))

	// Auth middleware
	authed := chain(
		middleware.Logger,
		middleware.CORS,
		middleware.RateLimit,
		middleware.Auth(cfg.JWTSecret),
	)

	// Dashboard
	mux.Handle("GET /api/v1/dashboard/summary", authed(http.HandlerFunc(dashboardH.Summary)))
	mux.Handle("GET /api/v1/dashboard/assets/trend", authed(http.HandlerFunc(dashboardH.AssetTrend)))
	mux.Handle("GET /api/v1/dashboard/shipments/trend", authed(http.HandlerFunc(dashboardH.ShipmentTrend)))
	mux.Handle("GET /api/v1/dashboard/plants/occupancy", authed(http.HandlerFunc(dashboardH.PlantOccupancy)))

	// Assets
	mux.Handle("GET /api/v1/assets", authed(http.HandlerFunc(assetsH.List)))
	mux.Handle("POST /api/v1/assets", authed(http.HandlerFunc(assetsH.Create)))
	mux.Handle("GET /api/v1/assets/{id}", authed(http.HandlerFunc(assetsH.Get)))
	mux.Handle("PATCH /api/v1/assets/{id}", authed(http.HandlerFunc(assetsH.Update)))
	mux.Handle("DELETE /api/v1/assets/{id}", authed(http.HandlerFunc(assetsH.Delete)))
	mux.Handle("GET /api/v1/assets/{id}/events", authed(http.HandlerFunc(assetsH.GetEvents)))
	mux.Handle("POST /api/v1/assets/scan", authed(http.HandlerFunc(assetsH.ScanQR)))

	// Shipments
	mux.Handle("GET /api/v1/shipments", authed(http.HandlerFunc(shipmentsH.List)))
	mux.Handle("POST /api/v1/shipments", authed(http.HandlerFunc(shipmentsH.Create)))
	mux.Handle("GET /api/v1/shipments/{id}", authed(http.HandlerFunc(shipmentsH.Get)))
	mux.Handle("POST /api/v1/shipments/{id}/transition", authed(http.HandlerFunc(shipmentsH.Transition)))
	mux.Handle("GET /api/v1/shipments/{id}/lines", authed(http.HandlerFunc(shipmentsH.GetLines)))
	mux.Handle("POST /api/v1/shipments/{id}/lines", authed(http.HandlerFunc(shipmentsH.AddLine)))
	mux.Handle("GET /api/v1/shipments/{id}/timeline", authed(http.HandlerFunc(shipmentsH.GetTimeline)))

	// Clients / CRM
	mux.Handle("GET /api/v1/clients", authed(http.HandlerFunc(clientsH.List)))
	mux.Handle("POST /api/v1/clients", authed(http.HandlerFunc(clientsH.Create)))
	mux.Handle("GET /api/v1/clients/{id}", authed(http.HandlerFunc(clientsH.Get)))
	mux.Handle("PATCH /api/v1/clients/{id}", authed(http.HandlerFunc(clientsH.Update)))
	mux.Handle("GET /api/v1/clients/{id}/contacts", authed(http.HandlerFunc(clientsH.ListContacts)))
	mux.Handle("POST /api/v1/clients/{id}/contacts", authed(http.HandlerFunc(clientsH.CreateContact)))
	mux.Handle("GET /api/v1/clients/{id}/interactions", authed(http.HandlerFunc(clientsH.ListInteractions)))
	mux.Handle("POST /api/v1/clients/{id}/interactions", authed(http.HandlerFunc(clientsH.CreateInteraction)))

	// Plants
	mux.Handle("GET /api/v1/plants", authed(http.HandlerFunc(plantsH.List)))
	mux.Handle("POST /api/v1/plants", authed(http.HandlerFunc(plantsH.Create)))
	mux.Handle("GET /api/v1/plants/{id}", authed(http.HandlerFunc(plantsH.Get)))
	mux.Handle("PATCH /api/v1/plants/{id}", authed(http.HandlerFunc(plantsH.Update)))
	mux.Handle("GET /api/v1/plants/{id}/zones", authed(http.HandlerFunc(plantsH.ListZones)))
	mux.Handle("POST /api/v1/plants/{id}/zones", authed(http.HandlerFunc(plantsH.CreateZone)))

	return withCORSPreflight(mux)
}

func chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}

// withCORSPreflight answers every OPTIONS request directly, before it ever
// reaches the mux. net/http's ServeMux only matches the exact HTTP method a
// route was registered with (e.g. "GET /api/v1/assets"), so a browser's
// preflight OPTIONS request - triggered by cross-origin calls that send an
// Authorization header, like the dev workflow in docs/local-deployment.md
// where the web UI and API run on different ports - would otherwise get a
// 405 from the mux itself, never reaching middleware.CORS on the real route.
func withCORSPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			middleware.CORS(next).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
