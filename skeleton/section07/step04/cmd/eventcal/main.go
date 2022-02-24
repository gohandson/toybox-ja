package main

import (
	"os"

	eventcal "github.com/gohandson/toybox-ja/skeleton/section07/step04"
)

func main() {
	cli := eventcal.CLI{
		Calendar: eventcal.NewCalendar(),
		Stdout:   os.Stdout,
		Stderr:   os.Stderr,
		Stdin:    os.Stdin,
	}
	os.Exit(cli.Main())
}
