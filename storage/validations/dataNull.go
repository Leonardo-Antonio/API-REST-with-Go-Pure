package validations

import (
	"database/sql"
)

// StringNull valid nulls
func StringNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// IntNull valid nulls
func IntNull(n int32) sql.NullInt32 {
	null := sql.NullInt32{Int32: n}
	if null.Int32 != 0 {
		null.Valid = true
	}
	return null
}
