package airtable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.airtable.com/v0"

// Client wraps the Airtable REST API.
type Client struct {
	apiKey string
	baseID string
	http   *http.Client
}

func NewClient(apiKey, baseID string) *Client {
	return &Client{
		apiKey: apiKey,
		baseID: baseID,
		http:   &http.Client{Timeout: 15 * time.Second},
	}
}

type Record struct {
	ID     string                 `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields"`
}

type listResponse struct {
	Records []*Record `json:"records"`
	Offset  string    `json:"offset,omitempty"`
}

type createRequest struct {
	Records []*Record `json:"records"`
}

type updateRequest struct {
	Records []*Record `json:"records"`
}

// List fetches all records from a table (handles pagination).
func (c *Client) List(table string, filterFormula string) ([]*Record, error) {
	var all []*Record
	offset := ""
	for {
		url := fmt.Sprintf("%s/%s/%s?pageSize=100", baseURL, c.baseID, table)
		if filterFormula != "" {
			url += "&filterByFormula=" + filterFormula
		}
		if offset != "" {
			url += "&offset=" + offset
		}

		body, err := c.get(url)
		if err != nil {
			return nil, err
		}

		var resp listResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, fmt.Errorf("unmarshal: %w", err)
		}
		all = append(all, resp.Records...)
		if resp.Offset == "" {
			break
		}
		offset = resp.Offset
	}
	return all, nil
}

// Get fetches a single record by ID.
func (c *Client) Get(table, recordID string) (*Record, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", baseURL, c.baseID, table, recordID)
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}
	var r Record
	return &r, json.Unmarshal(body, &r)
}

// Create inserts up to 10 records at once.
func (c *Client) Create(table string, records []*Record) ([]*Record, error) {
	url := fmt.Sprintf("%s/%s/%s", baseURL, c.baseID, table)
	payload, _ := json.Marshal(createRequest{Records: records})
	body, err := c.do(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}
	var resp listResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return resp.Records, nil
}

// Upsert updates existing records or creates new ones (10 max per call).
func (c *Client) Upsert(table string, records []*Record) ([]*Record, error) {
	url := fmt.Sprintf("%s/%s/%s", baseURL, c.baseID, table)
	payload, _ := json.Marshal(updateRequest{Records: records})
	body, err := c.do(http.MethodPatch, url, payload)
	if err != nil {
		return nil, err
	}
	var resp listResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return resp.Records, nil
}

// Delete removes a record.
func (c *Client) Delete(table, recordID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", baseURL, c.baseID, table, recordID)
	_, err := c.do(http.MethodDelete, url, nil)
	return err
}

func (c *Client) get(url string) ([]byte, error) {
	return c.do(http.MethodGet, url, nil)
}

func (c *Client) do(method, url string, body []byte) ([]byte, error) {
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("airtable %d: %s", resp.StatusCode, respBody)
	}
	return respBody, nil
}
