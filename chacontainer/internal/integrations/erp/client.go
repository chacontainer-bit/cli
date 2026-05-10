package erp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a generic ERP REST adapter.
// Swap the base URL and auth method for SAP, Oracle, Odoo, or custom ERP.
type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		http:    &http.Client{Timeout: 30 * time.Second},
	}
}

// ── Purchase Orders ───────────────────────────────────────────────────────────

type PurchaseOrder struct {
	ERPID       string     `json:"id"`
	Reference   string     `json:"reference"`
	VendorID    string     `json:"vendor_id"`
	Status      string     `json:"status"`
	TotalAmount float64    `json:"total_amount"`
	Currency    string     `json:"currency"`
	Lines       []POLine   `json:"lines"`
	IssuedAt    time.Time  `json:"issued_at"`
	ExpectedAt  time.Time  `json:"expected_at"`
}

type POLine struct {
	SKU       string  `json:"sku"`
	Desc      string  `json:"description"`
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

func (c *Client) GetPurchaseOrder(erpID string) (*PurchaseOrder, error) {
	body, err := c.get("/api/purchase-orders/" + erpID)
	if err != nil {
		return nil, err
	}
	var po PurchaseOrder
	return &po, json.Unmarshal(body, &po)
}

// ── Inventory ─────────────────────────────────────────────────────────────────

type StockLevel struct {
	SKU         string  `json:"sku"`
	LocationID  string  `json:"location_id"`
	Quantity    float64 `json:"quantity"`
	UnitOfMeasure string `json:"uom"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Client) GetStockLevel(sku, locationID string) (*StockLevel, error) {
	url := fmt.Sprintf("/api/inventory?sku=%s&location=%s", sku, locationID)
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}
	var sl StockLevel
	return &sl, json.Unmarshal(body, &sl)
}

func (c *Client) AdjustStock(sku, locationID string, delta float64, reason string) error {
	payload, _ := json.Marshal(map[string]interface{}{
		"sku":        sku,
		"location":   locationID,
		"delta":      delta,
		"reason":     reason,
		"timestamp":  time.Now().UTC(),
	})
	_, err := c.post("/api/inventory/adjustment", payload)
	return err
}

// ── GR / Goods Receipt ───────────────────────────────────────────────────────

type GoodsReceipt struct {
	ERPID     string    `json:"id"`
	POID      string    `json:"po_id"`
	PlantCode string    `json:"plant_code"`
	ReceivedAt time.Time `json:"received_at"`
	Lines     []GRLine  `json:"lines"`
}

type GRLine struct {
	SKU      string  `json:"sku"`
	Received float64 `json:"received_qty"`
	Accepted float64 `json:"accepted_qty"`
	Rejected float64 `json:"rejected_qty"`
}

func (c *Client) PostGoodsReceipt(gr *GoodsReceipt) (string, error) {
	payload, _ := json.Marshal(gr)
	body, err := c.post("/api/goods-receipts", payload)
	if err != nil {
		return "", err
	}
	var resp struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}
	return resp.ID, nil
}

// ── Delivery Notification ────────────────────────────────────────────────────

type DeliveryNotification struct {
	ShipmentRef string    `json:"shipment_ref"`
	ClientID    string    `json:"client_id"`
	DeliveredAt time.Time `json:"delivered_at"`
	Items       int       `json:"items"`
	WeightKg    float64   `json:"weight_kg"`
}

func (c *Client) NotifyDelivery(dn *DeliveryNotification) error {
	payload, _ := json.Marshal(dn)
	_, err := c.post("/api/deliveries/notify", payload)
	return err
}

// ── Transport ─────────────────────────────────────────────────────────────────

func (c *Client) get(path string) ([]byte, error) {
	return c.do(http.MethodGet, path, nil)
}

func (c *Client) post(path string, body []byte) ([]byte, error) {
	return c.do(http.MethodPost, path, body)
}

func (c *Client) do(method, path string, body []byte) ([]byte, error) {
	var req *http.Request
	var err error
	url := c.baseURL + path
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
		return nil, fmt.Errorf("erp http: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("erp %d: %s", resp.StatusCode, respBody)
	}
	return respBody, nil
}
