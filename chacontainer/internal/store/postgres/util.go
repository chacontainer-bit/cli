package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// nullStr converts an empty Go string to SQL NULL, otherwise passes it
// through. Used for optional foreign keys / text columns.
func nullStr(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

func strOrEmpty(s sql.NullString) string {
	if !s.Valid {
		return ""
	}
	return s.String
}

func nullTime(t *time.Time) sql.NullTime {
	if t == nil || t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func timeOrNil(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	tm := t.Time
	return &tm
}

// marshalJSON encodes any Go value (struct, map, slice) as a JSON string
// suitable for a jsonb column parameter. A nil/zero value still marshals to
// "null"/"{}"/"[]", never to a Go nil, so jsonb columns declared NOT NULL are
// always satisfied.
func marshalJSON(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// unmarshalJSON decodes a jsonb column (scanned as []byte) into dst. A NULL
// or empty column is treated as a no-op so callers can zero-initialize dst
// first.
func unmarshalJSON(raw []byte, dst interface{}) error {
	if len(raw) == 0 {
		return nil
	}
	return json.Unmarshal(raw, dst)
}

// nullDate renders an optional date as "YYYY-MM-DD" for a DATE column, or
// SQL NULL when unset.
func nullDate(t *time.Time) interface{} {
	if t == nil || t.IsZero() {
		return nil
	}
	return t.Format("2006-01-02")
}

// nullFloat treats the zero value as "not set" for optional NUMERIC columns
// such as purchase cost, where 0 and "unknown" are indistinguishable anyway.
func nullFloat(f float64) interface{} {
	if f == 0 {
		return nil
	}
	return f
}

func intOrZero(n sql.NullInt64) int {
	if !n.Valid {
		return 0
	}
	return int(n.Int64)
}

// columnKind tells buildUpdate how to convert a JSON-decoded patch value
// (string, float64, bool, map[string]interface{}, []interface{}, nil) into
// the right driver argument for a given column.
type columnKind int

const (
	kindText columnKind = iota
	kindNumber
	kindBool
	kindJSON
)

// buildUpdate turns a client-supplied PATCH body into a "col = $n, ..." SET
// clause plus matching args, restricted to an allow-list of columns so a
// request body can never write to a column it wasn't meant to touch.
// Unknown keys in patch are silently ignored rather than rejected, since the
// same PATCH endpoints double as "send only what changed".
func buildUpdate(allowed map[string]columnKind, patch map[string]interface{}) (setClause string, args []interface{}, err error) {
	var sets []string
	for key, val := range patch {
		kind, ok := allowed[key]
		if !ok {
			continue
		}
		arg, convErr := convertColumnValue(kind, val)
		if convErr != nil {
			return "", nil, fmt.Errorf("field %s: %w", key, convErr)
		}
		args = append(args, arg)
		sets = append(sets, fmt.Sprintf("%s = $%d", key, len(args)))
	}
	if len(sets) == 0 {
		return "", nil, fmt.Errorf("no valid fields to update")
	}
	return strings.Join(sets, ", "), args, nil
}

func convertColumnValue(kind columnKind, val interface{}) (interface{}, error) {
	if val == nil {
		return nil, nil
	}
	if kind == kindJSON {
		return marshalJSON(val)
	}
	return val, nil
}

// nullTimeStr renders a zero time.Time as "" (so a COALESCE(NULLIF($n,”)...
// NOW()) SQL pattern can default it), otherwise as RFC3339.
func nullTimeStr(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

// pagination normalizes page/per_page query params to sane bounds.
func pagination(page, perPage int) (int, int) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 200 {
		perPage = 50
	}
	return page, perPage
}
