package openai

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func NewString(s string, valid bool) NullString {
	return NullString{
		sql.NullString{
			String: s,
			Valid:  valid,
		},
	}
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		ns.String, ns.Valid = "", false

		return nil
	}

	if err := json.Unmarshal(b, &ns.String); err != nil {
		return err
	}

	ns.Valid = true

	return nil
}
