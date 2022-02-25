package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gohandson/toybox-ja/solution/section10/step01/eventwatcher"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := net.JoinHostPort("", port)

	ew, err := eventwatcher.New(addr)
	if err != nil {
		return err
	}

	if err := ew.Start(); err != nil {
		return err
	}

	return nil
}
