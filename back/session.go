package main

import (
	"errors"
	"sync"
)

var sessions = sync.Map{}

func createSession(password string) (string, error) {
	if config.Api.AdminPassword == "" {
		return "", errors.New("no admin password")
	}
	if password != config.Api.AdminPassword {
		return "", errors.New("invalid admin password")
	}
	session := randomString(16)
	sessions.Store(session, struct{}{})
	return session, nil
}

func checkSession(session string) bool {
	_, ok := sessions.Load(session)
	return ok
}
