package main

import (
	"math/rand"
	"time"
)

func initRandom() {
	rand.Seed(time.Now().UnixNano())
}
