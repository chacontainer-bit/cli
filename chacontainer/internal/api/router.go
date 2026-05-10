package api

import (
	"encoding/json"
	"net/http"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
	"github.com/cli/cli/v2/chacontainer/internal/config"
)

// NewRouter wires all routes. Dependencies (stores, integrations) are injected
// via the config at startup; replace stub implementations with postgres stores
// before shipping.
func NewRouter(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "chacontainer"})
	})

	// Stub stores (replace with postgres implementations)
	assetStore := &stubAssetStore{}
	shipmentStore := &stubShipmentStore{}
	clientStore := &stubClientStore{}
	plantStore := &stubPlantStore{}
	statsStore := &stubStatsStore{}
	webhookProcessor := &stubWebhookProcessor{}

	assetsH := handlers.NewAssetsHandler(assetStore)
	shipmentsH := handlers.NewShipmentsHandler(shipmentStore)
	clientsH := handlers.NewClientsHandler(clientStore)
	plantsH := handlers.NewPlantsHandler(plantStore)
	dashboardH := handlers.NewDashboardHandler(statsStore)
	webhooksH := handlers.NewWebhooksHandler(webhookProcessor, cfg.MakeWebhookSecret)

	// Public webhook receivers (no JWT)
	mux.HandleFunc("POST /webhooks/make", webhooksH.Make)
	mux.HandleFunc("POST /webhooks/erp", webhooksH.ERP)
	mux.HandleFunc("POST /webhooks/airtable", webhooksH.Airtable)

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

	return mux
}

func chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
