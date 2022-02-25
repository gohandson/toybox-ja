package main

import (
	"fmt"
	"os"

	"github.com/gohandson/toybox-ja/skeleton/section09/step04/eventwatcher"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	ew, err := eventwatcher.New(":8080")
	if err != nil {
		return err
	}

	if err := ew.Start(); err != nil {
		return err
	}

	return nil
}
