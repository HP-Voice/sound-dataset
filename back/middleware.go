package main

import "net/http"

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.Api.AdminPassword == "" || r.Header.Get("Admin-Password") != config.Api.AdminPassword {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("no admin password"))
			return
		}
		next(w, r)
	}
}
