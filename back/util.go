package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/pgtype"
)

type UUID struct {
	pgtype.UUID
}

func (u *UUID) MarshalJSON() ([]byte, error) {
	if u.Status == pgtype.Null || u.Status == pgtype.Undefined {
		return json.Marshal(nil)
	}
	return json.Marshal(u.String())
}

func (u *UUID) UnmarshalJSON(data []byte) error {
	var s *string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == nil {
		u.Status = pgtype.Null
	}
	bytes, err := hex.DecodeString(*s)
	if err != nil {
		return err
	}
	u.Bytes = *(*[16]byte)(bytes)
	u.Status = pgtype.Present
	return nil
}

func (u *UUID) String() string {
	return fmt.Sprintf("%x", u.Bytes)
}
