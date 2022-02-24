package main

import (
	"os"

	eventcal "github.com/gohandson/toybox-ja/solution/section07/step03"
)

func main() {
	cli := eventcal.CLI{
		Calendar: eventcal.NewCalendar(),
	}
	os.Exit(cli.Main())
}
