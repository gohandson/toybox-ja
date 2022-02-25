package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"regexp"

	imgcheck "github.com/gohandson/toybox-ja/solution/section08/step04"
)

var (
	flagFormat        string
	flagFormatPattern string
	flagMaxHeight     int
	flagMaxWidth      int
)

func init() {
	flag.StringVar(&flagFormat, "format", "", "allow image format")
	flag.StringVar(&flagFormatPattern, "format-pattern", "", "allow image format")
	flag.IntVar(&flagMaxHeight, "height", -1, "max height")
	flag.IntVar(&flagMaxWidth, "width", -1, "max width")
}

func main() {
	// TODO: 不要な場合は消す
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "エラー:", r)
			os.Exit(1)
		}
	}()

	flag.Parse()

	var rules []imgcheck.Rule
	if flagFormat != "" {
		rules = append(rules, imgcheck.Format(flagFormat))
	}
	if flagFormatPattern != "" {
		// TODO: MustCompileを使わないように修正する
		pattern := regexp.MustCompile(flagFormatPattern)
		rules = append(rules, imgcheck.FormatPattern(pattern))
	}
	if flagMaxHeight > 0 {
		rules = append(rules, imgcheck.MaxHeight(flagMaxHeight))
	}
	if flagMaxWidth > 0 {
		rules = append(rules, imgcheck.MaxWidth(flagMaxWidth))
	}

	if err := imgcheck.ValidateDir(flag.Arg(0), rules...); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
