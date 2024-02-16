package main

import (
	"fmt"
	"os"

	"github.com/gohandson/toybox-ja/solution/section10/step05/eventwatcher"
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
