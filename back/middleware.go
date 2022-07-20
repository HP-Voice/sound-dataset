package main

import "net/http"

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:5173" || origin == "http://localhost:8080" || origin == "https://uquark.me" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Admin-Password")
		if r.Method == http.MethodOptions {
			return
		}
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
