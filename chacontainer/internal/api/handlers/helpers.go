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

// pathParam reads a {name} wildcard registered on the route pattern (e.g.
// "GET /api/v1/assets/{id}/events") via net/http's ServeMux path values.
// This must not fall back to "last URL segment" - nested routes like
// .../{id}/events or .../{id}/transition have the id in the middle of the
// path, not at the end.
func pathParam(r *http.Request, name string) string {
	return r.PathValue(name)
}
