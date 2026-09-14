package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/cli/cli/v2/chacontainer/internal/integrations/symphony"
)

// defaultSymphonyMaxWait bounds how long an HTTP request will block waiting
// for Symphony to reply before the caller should poll a background job
// instead. Symphony's own docs note replies can take a minute or two.
const (
	defaultSymphonyMaxWait      = 110 * time.Second
	defaultSymphonyPollInterval = 3 * time.Second
)

// SymphonyAgent sends a message to Symphony and waits for its reply.
type SymphonyAgent interface {
	Ask(message string, maxWait, pollInterval time.Duration) (string, error)
}

// SymphonyHandler exposes CHACONTAINER's Symphony integration: any
// authenticated user of this tenant can ask Symphony to act on the
// business's behalf (follow up with a client, schedule, build a report).
type SymphonyHandler struct {
	agent SymphonyAgent
}

func NewSymphonyHandler(agent SymphonyAgent) *SymphonyHandler {
	return &SymphonyHandler{agent: agent}
}

type symphonyAskRequest struct {
	Message string `json:"message"`
}

type symphonyAskResponse struct {
	Reply string `json:"reply"`
}

// Ask relays a message to Symphony and returns its reply once available.
// Symphony's own turns can take a minute or two, so this blocks up to
// defaultSymphonyMaxWait before giving up.
func (h *SymphonyHandler) Ask(w http.ResponseWriter, r *http.Request) {
	if h.agent == nil {
		writeError(w, http.StatusServiceUnavailable, "symphony integration not configured (missing SYMPHONY_API_TOKEN)")
		return
	}

	var req symphonyAskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if req.Message == "" {
		writeError(w, http.StatusBadRequest, "message is required")
		return
	}

	reply, err := h.agent.Ask(req.Message, defaultSymphonyMaxWait, defaultSymphonyPollInterval)
	if err != nil {
		if errors.Is(err, symphony.ErrTimeout) {
			writeError(w, http.StatusGatewayTimeout, "symphony did not reply in time")
			return
		}
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, symphonyAskResponse{Reply: reply})
}
