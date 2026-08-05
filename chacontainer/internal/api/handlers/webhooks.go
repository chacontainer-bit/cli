package handlers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// WebhookEvent is received from Make.com or external integrations.
type WebhookEvent struct {
	Source  string                 `json:"source"`
	Event   string                 `json:"event"`
	Payload map[string]interface{} `json:"payload"`
}

type WebhookProcessor interface {
	ProcessMakeEvent(event WebhookEvent) error
	ProcessERPEvent(event WebhookEvent) error
	ProcessAirtableEvent(event WebhookEvent) error
}

type WebhooksHandler struct {
	processor WebhookProcessor
	secret    string
}

func NewWebhooksHandler(processor WebhookProcessor, secret string) *WebhooksHandler {
	return &WebhooksHandler{processor: processor, secret: secret}
}

func (h *WebhooksHandler) Make(w http.ResponseWriter, r *http.Request) {
	if !h.verifySignature(r, "X-Make-Signature") {
		writeError(w, http.StatusUnauthorized, "invalid signature")
		return
	}

	var event WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	event.Source = "make"

	go func() {
		if err := h.processor.ProcessMakeEvent(event); err != nil {
			log.Printf("make webhook error: %v", err)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
}

func (h *WebhooksHandler) ERP(w http.ResponseWriter, r *http.Request) {
	if !h.verifySignature(r, "X-ERP-Signature") {
		writeError(w, http.StatusUnauthorized, "invalid signature")
		return
	}

	var event WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	event.Source = "erp"

	go func() {
		if err := h.processor.ProcessERPEvent(event); err != nil {
			log.Printf("erp webhook error: %v", err)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
}

func (h *WebhooksHandler) Airtable(w http.ResponseWriter, r *http.Request) {
	if !h.verifySignature(r, "X-Airtable-Signature") {
		writeError(w, http.StatusUnauthorized, "invalid signature")
		return
	}

	var event WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	event.Source = "airtable"

	go func() {
		if err := h.processor.ProcessAirtableEvent(event); err != nil {
			log.Printf("airtable webhook error: %v", err)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
}

// verifySignature reads the body, verifies HMAC, then resets r.Body for later reading.
// A webhook is rejected outright if no secret is configured; an unsigned
// endpoint should never be reachable, not silently unauthenticated.
func (h *WebhooksHandler) verifySignature(r *http.Request, header string) bool {
	if h.secret == "" {
		log.Printf("webhook %s rejected: no signing secret configured", header)
		return false
	}
	sig := r.Header.Get(header)
	if sig == "" {
		return false
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}
	r.Body = io.NopCloser(bytes.NewReader(body))
	mac := hmac.New(sha256.New, []byte(h.secret))
	mac.Write(body)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(sig), []byte(expected))
}
