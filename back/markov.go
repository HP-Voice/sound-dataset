package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
)

func initMarkov() error {
	return exec.Command("python", "markov.py", "book.txt").Start()
}

func getSentence() ([]byte, error) {
	resp, err := http.Get("http://localhost:5000/sentence")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
