package main

import (
	"encoding/json"
	"os"
)

var config = struct {
	Db string
	Fs struct {
		Path      string
		Extension string
		BlockSize int
	}
	Api struct {
		Address       string
		AdminPassword string
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
