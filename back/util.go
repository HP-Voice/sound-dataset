package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/pgtype"
	"math/rand"
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

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(result)
}
