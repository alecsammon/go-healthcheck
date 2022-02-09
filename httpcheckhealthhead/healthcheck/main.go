package main

import (
	"net/http"
	"syscall"
)

func main() {
	c := http.Client{}
	r, err := c.Head("http://127.0.0.1/check/health")

	if err != nil || r.StatusCode != 200 {
		syscall.Exit(1)
	}
}
