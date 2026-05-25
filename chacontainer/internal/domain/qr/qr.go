package qr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// ScanEvent is recorded every time a QR code is scanned.
type ScanEvent struct {
	ID         string     `json:"id" db:"id"`
	TenantID   string     `json:"tenant_id" db:"tenant_id"`
	AssetID    string     `json:"asset_id" db:"asset_id"`
	QRCode     string     `json:"qr_code" db:"qr_code"`
	ScannedBy  string     `json:"scanned_by" db:"scanned_by"`
	PlantID    string     `json:"plant_id,omitempty" db:"plant_id"`
	ZoneID     string     `json:"zone_id,omitempty" db:"zone_id"`
	Lat        float64    `json:"lat,omitempty" db:"lat"`
	Lng        float64    `json:"lng,omitempty" db:"lng"`
	DeviceID   string     `json:"device_id,omitempty" db:"device_id"`
	AppVersion string     `json:"app_version,omitempty" db:"app_version"`
	Action     ScanAction `json:"action" db:"action"`
	Notes      string     `json:"notes,omitempty" db:"notes"`
	ScannedAt  time.Time  `json:"scanned_at" db:"scanned_at"`
}

type ScanAction string

const (
	ActionCheckIn     ScanAction = "check_in"
	ActionCheckOut    ScanAction = "check_out"
	ActionInventory   ScanAction = "inventory"
	ActionTransfer    ScanAction = "transfer"
	ActionInspect     ScanAction = "inspect"
	ActionMaintenance ScanAction = "maintenance"
)

// QRPayload is the data embedded in a QR code.
type QRPayload struct {
	TenantID  string `json:"t"`
	AssetID   string `json:"a"`
	AssetCode string `json:"c"`
	Version   int    `json:"v"`
	Signature string `json:"s"`
}

// GenerateCode produces a signed QR payload string.
func GenerateCode(tenantID, assetID, assetCode, secret string) string {
	payload := fmt.Sprintf("%s|%s|%s|1", tenantID, assetID, assetCode)
	sig := sign(payload, secret)
	return fmt.Sprintf("%s|%s", payload, sig[:12])
}

// VerifyCode validates the QR signature.
func VerifyCode(code, secret string) bool {
	if len(code) < 14 {
		return false
	}
	parts := splitLast(code, "|")
	if len(parts) != 2 {
		return false
	}
	sig := sign(parts[0], secret)
	return sig[:12] == parts[1]
}

func sign(data, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func splitLast(s, sep string) []string {
	idx := -1
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == sep {
			idx = i
			break
		}
	}
	if idx == -1 {
		return []string{s}
	}
	return []string{s[:idx], s[idx+1:]}
}
