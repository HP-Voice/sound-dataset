package main

import "flag"

var flags = struct {
	config *string
	clean  *bool
}{}

func initFlags() {
	flags.config = flag.String("config", "config.json", "configuration file path")
	flags.clean = flag.Bool("clean", false, "clean files")
	flag.Parse()
}
