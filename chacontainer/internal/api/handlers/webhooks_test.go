package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http/httptest"
	"strings"
	"testing"
)

type noopProcessor struct{}

func (noopProcessor) ProcessMakeEvent(WebhookEvent) error     { return nil }
func (noopProcessor) ProcessERPEvent(WebhookEvent) error      { return nil }
func (noopProcessor) ProcessAirtableEvent(WebhookEvent) error { return nil }

func sign(secret, body string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

func TestVerifySignature_RejectsWhenSecretMissing(t *testing.T) {
	h := NewWebhooksHandler(noopProcessor{}, "")
	body := `{"event":"x"}`
	r := httptest.NewRequest("POST", "/webhooks/make", strings.NewReader(body))
	r.Header.Set("X-Make-Signature", sign("whatever", body))

	if h.verifySignature(r, "X-Make-Signature") {
		t.Error("verifySignature() = true, want false when no secret is configured")
	}
}

func TestVerifySignature_RejectsBadSignature(t *testing.T) {
	h := NewWebhooksHandler(noopProcessor{}, "topsecret")
	body := `{"event":"x"}`
	r := httptest.NewRequest("POST", "/webhooks/make", strings.NewReader(body))
	r.Header.Set("X-Make-Signature", "not-the-right-signature")

	if h.verifySignature(r, "X-Make-Signature") {
		t.Error("verifySignature() = true, want false for a bad signature")
	}
}

func TestVerifySignature_AcceptsValidSignature(t *testing.T) {
	h := NewWebhooksHandler(noopProcessor{}, "topsecret")
	body := `{"event":"x"}`
	r := httptest.NewRequest("POST", "/webhooks/make", strings.NewReader(body))
	r.Header.Set("X-Make-Signature", sign("topsecret", body))

	if !h.verifySignature(r, "X-Make-Signature") {
		t.Error("verifySignature() = false, want true for a valid signature")
	}
}

func TestAirtableWebhook_RequiresSignature(t *testing.T) {
	h := NewWebhooksHandler(noopProcessor{}, "topsecret")
	body := `{"event":"record.updated"}`
	r := httptest.NewRequest("POST", "/webhooks/airtable", strings.NewReader(body))
	w := httptest.NewRecorder()

	h.Airtable(w, r)

	if w.Code != 401 {
		t.Errorf("Airtable() status = %d, want 401 for an unsigned request", w.Code)
	}
}
