package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime/trace"
	"time"

	imgconv "github.com/gohandson/toybox-ja/solution/section11/step05"
)

var (
	flagTo   = imgconv.PNG
	flagFrom = imgconv.TIFF
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}

func run() (rerr error) {
	f, err := os.Create("trace.out")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil && rerr != nil {
			rerr = err
		}
	}()

	if err := trace.Start(f); err != nil {
		return err
	}
	defer trace.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ctx, task := trace.NewTask(ctx, "imgconv")
	defer task.End()
	if err := imgconv.ConvertAll(ctx, os.Args[1], flagFrom, flagTo); err != nil {
		return err
	}

	return nil
}
