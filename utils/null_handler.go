package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
// func (ns *NullString) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &ns.String)
// 	ns.Valid = (err == nil)
// 	return err
// }
func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println("errorwkwk", err)
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
	return nil
}