package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func intParam(s string, fallback int) int {
	if s == "" {
		return fallback
	}
	n, err := strconv.Atoi(s)
	if err != nil || n < 1 {
		return fallback
	}
	return n
}

// pathParam extracts a named parameter from the route pattern, e.g. "id"
// from "GET /api/v1/shipments/{id}/lines". Go 1.22+'s net/http ServeMux
// resolves these via r.PathValue; every route registered in router.go
// uses the {name} pattern syntax so this always matches.
func pathParam(r *http.Request, name string) string {
	return r.PathValue(name)
}
