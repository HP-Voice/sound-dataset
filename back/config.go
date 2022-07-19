package main

import (
	"encoding/json"
	"github.com/jackc/pgx"
	"os"
)

var config = struct {
	Db pgx.ConnConfig
	Fs struct {
		Path      string
		Extension string
		BlockSize int
	}
	Api struct {
		Address string
	}
	Static struct {
		Address string
		Root    string
	}
	Tls *struct {
		Cert string
		Key  string
	}
}{}

func initConfig(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(&config)
}
