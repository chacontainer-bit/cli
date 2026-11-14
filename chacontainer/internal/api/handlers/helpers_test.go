package handlers

import (
	"net/http/httptest"
	"testing"
)

// pathParam must read the named wildcard set by http.ServeMux, not the last
// URL segment — a route like "/assets/{id}/events" has "events" as its last
// segment, which is not the asset id.
func TestPathParam(t *testing.T) {
	tests := []struct {
		name string
		path string
		id   string
		want string
	}{
		{name: "simple resource", path: "/api/v1/assets/abc123", id: "abc123", want: "abc123"},
		{name: "nested sub-resource", path: "/api/v1/assets/abc123/events", id: "abc123", want: "abc123"},
		{name: "nested sub-resource with trailing segment", path: "/api/v1/shipments/xyz/lines", id: "xyz", want: "xyz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.path, nil)
			r.SetPathValue("id", tt.id)

			if got := pathParam(r, "id"); got != tt.want {
				t.Errorf("pathParam() = %q, want %q", got, tt.want)
			}
		})
	}
}
