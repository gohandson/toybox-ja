package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gohandson/toybox-ja/skeleton/section10/step03/eventwatcher"
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

	ew, err := eventwatcher.New(context.Background(), addr)
	if err != nil {
		return err
	}

	if err := ew.Start(); err != nil {
		return err
	}

	return nil
}
