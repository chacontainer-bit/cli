// Package symphony talks to Symphony (Wix "individuals-chat" agent) — the
// user's team of AI business agents that can schedule, reach contacts,
// follow up, and build reports on the account holder's behalf.
package symphony

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultBaseURL = "https://symphony.wix.com/individuals-chat/poc/agent"

// ErrTimeout is returned by Ask when Symphony accepted the request but did
// not produce a reply within the caller's poll budget.
var ErrTimeout = errors.New("symphony: timed out waiting for a reply")

// Client wraps the Symphony ask/reply HTTP API.
type Client struct {
	baseURL   string
	apiToken  string
	sessionID string
	http      *http.Client
}

// Option customizes a Client returned by NewClient.
type Option func(*Client)

// WithBaseURL overrides the Symphony API base URL (mainly for tests).
func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = url }
}

// WithSessionID sets the sessionId used to keep replies in one thread.
// Defaults to "chacontainer".
func WithSessionID(id string) Option {
	return func(c *Client) { c.sessionID = id }
}

// WithHTTPClient overrides the underlying *http.Client.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.http = h }
}

// NewClient builds a Symphony client. apiToken is the bearer token issued by
// Symphony → Settings → Connections; it must come from configuration
// (SYMPHONY_API_TOKEN), never be hardcoded, since it authorizes acting on
// the account holder's behalf with full authority.
func NewClient(apiToken string, opts ...Option) *Client {
	c := &Client{
		baseURL:   defaultBaseURL,
		apiToken:  apiToken,
		sessionID: "chacontainer",
		http:      &http.Client{Timeout: 30 * time.Second},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type askRequest struct {
	Message   string `json:"message"`
	SessionID string `json:"sessionId"`
}

type askResponse struct {
	Status         string `json:"status"`
	Reply          string `json:"reply"`
	ConversationID string `json:"conversationId"`
}

type replyRequest struct {
	ConversationID string `json:"conversationId"`
}

type replyResponse struct {
	Status string `json:"status"`
	Reply  string `json:"reply"`
}

// Ask sends message to Symphony. If Symphony answers immediately it returns
// the reply. Otherwise it polls Reply every pollInterval, up to maxWait,
// until Symphony finishes answering — mirroring the ask/reply contract
// documented for the Symphony agent endpoint.
func (c *Client) Ask(message string, maxWait, pollInterval time.Duration) (string, error) {
	resp, err := c.ask(message)
	if err != nil {
		return "", err
	}
	if resp.Reply != "" {
		return resp.Reply, nil
	}
	if resp.Status != "accepted" {
		return "", fmt.Errorf("symphony: unexpected status %q", resp.Status)
	}
	if resp.ConversationID == "" {
		return "", errors.New("symphony: accepted response missing conversationId")
	}

	deadline := time.Now().Add(maxWait)
	for {
		time.Sleep(pollInterval)

		reply, answered, err := c.Reply(resp.ConversationID)
		if err != nil {
			return "", err
		}
		if answered {
			return reply, nil
		}
		if time.Now().After(deadline) {
			return "", ErrTimeout
		}
	}
}

// ask performs a single POST to /ask.
func (c *Client) ask(message string) (*askResponse, error) {
	payload, err := json.Marshal(askRequest{Message: message, SessionID: c.sessionID})
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	body, err := c.do(http.MethodPost, c.baseURL+"/ask", payload)
	if err != nil {
		return nil, err
	}

	var resp askResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal ask response: %w", err)
	}
	return &resp, nil
}

// Reply performs a single POST to /reply and reports whether Symphony has
// finished answering conversationID.
func (c *Client) Reply(conversationID string) (reply string, answered bool, err error) {
	payload, err := json.Marshal(replyRequest{ConversationID: conversationID})
	if err != nil {
		return "", false, fmt.Errorf("marshal: %w", err)
	}

	body, err := c.do(http.MethodPost, c.baseURL+"/reply", payload)
	if err != nil {
		return "", false, err
	}

	var resp replyResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", false, fmt.Errorf("unmarshal reply response: %w", err)
	}
	return resp.Reply, resp.Status == "answered" && resp.Reply != "", nil
}

func (c *Client) do(method, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("symphony http: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("symphony %d: %s", resp.StatusCode, respBody)
	}
	return respBody, nil
}
