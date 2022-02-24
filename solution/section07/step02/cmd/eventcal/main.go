package main

import (
	"os"

	eventcal "github.com/gohandson/toybox-ja/solution/section07/step02"
)

func main() {
	cli := eventcal.CLI{
		Calendar: eventcal.NewCalendar(),
	}
	os.Exit(cli.Main())
}
