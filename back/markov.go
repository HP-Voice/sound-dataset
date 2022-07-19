package main

import "os/exec"

func initMarkov() error {
	return exec.Command("python", "markov.py", "book.txt").Start()
}
