package main

import (
	"flag"
	"fmt"
	"os"

	imgconv "github.com/gohandson/toybox-ja/solution/section06/step04"
)

var (
	flagTo   = imgconv.PNG
	flagFrom = imgconv.JPEG
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	if err := imgconv.ConvertAll(os.Args[1], flagFrom, flagTo); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
