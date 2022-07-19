package main

import "flag"

var flags = struct {
	config *string
}{}

func initFlags() {
	flags.config = flag.String("config", "config.json", "configuration file path")
	flag.Parse()
}
