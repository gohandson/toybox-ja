package main

import (
	"os"

	eventcal "github.com/gohandson/toybox-ja/solution/section07/step01"
)

func main() {
	cli := eventcal.CLI{
		Calendar: eventcal.NewCalendar(),
	}
	os.Exit(cli.Main())
}
