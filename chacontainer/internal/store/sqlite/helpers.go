package sqlite

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func newID() string {
	return uuid.NewString()
}

func nowStr() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

func fmtTime(t time.Time) string {
	if t.IsZero() {
		return nowStr()
	}
	return t.UTC().Format(time.RFC3339Nano)
}

func parseTime(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return time.Time{}
	}
	return t
}

func nullTimePtr(t *time.Time) interface{} {
	if t == nil || t.IsZero() {
		return nil
	}
	return t.UTC().Format(time.RFC3339Nano)
}

func parseTimePtr(ns sql.NullString) *time.Time {
	if !ns.Valid || ns.String == "" {
		return nil
	}
	t := parseTime(ns.String)
	return &t
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func fromJSON(s string, v interface{}) {
	if s == "" {
		return
	}
	_ = json.Unmarshal([]byte(s), v)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
