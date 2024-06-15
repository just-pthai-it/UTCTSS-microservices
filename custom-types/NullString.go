package custom_types

import (
	"database/sql"
	"encoding/json"
)

// swagger:type string
// nullable: true
type NullString struct {
	sql.NullString
}

func (nullString NullString) MarshalJSON() ([]byte, error) {
	if nullString.Valid {
		return json.Marshal(nullString.String)
	}

	return json.Marshal(nil)
}

func (nullString *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		nullString.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &nullString.String)
	if err == nil {
		nullString.Valid = true
	}

	return err
}
