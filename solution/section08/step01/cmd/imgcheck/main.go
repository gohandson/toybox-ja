package main

import (
	"flag"
	"fmt"
	"os"
	_ "image/png"
	_ "image/jpeg"

	imgcheck "github.com/gohandson/toybox-ja/solution/section08/step01"
)

var (
	flagFormat string
)

func init() {
	flag.StringVar(&flagFormat, "format", "", "allow image format")
}

func main() {
	flag.Parse()

	var rules []imgcheck.Rule
	if flagFormat != "" {
		rules = append(rules, imgcheck.Format(flagFormat))
	}

	if err := imgcheck.ValidateDir(flag.Arg(0), rules...); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
