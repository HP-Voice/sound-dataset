package main

import "net/http"

func initStatic() error {
	handler := http.FileServer(http.Dir(config.Static.Root))
	return http.ListenAndServe(config.Static.Address, handler)
}
