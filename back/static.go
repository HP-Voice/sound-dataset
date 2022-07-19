package main

import "net/http"

func initStatic() error {
	handler := http.FileServer(http.Dir(config.Static.Root))
	if config.Tls == nil {
		return http.ListenAndServe(config.Static.Address, handler)
	} else {
		return http.ListenAndServeTLS(config.Static.Address, config.Tls.Cert, config.Tls.Key, handler)
	}
}
