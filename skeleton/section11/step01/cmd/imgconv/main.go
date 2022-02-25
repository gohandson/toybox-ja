package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime/trace"

	imgconv "github.com/gohandson/toybox-ja/skeleton/section11/step01"
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

	// TODO: 出力先はファイルfとしtrace.Startを呼ぶ
	
	// TODO: deferでtrace.Stopを呼ぶ


	ctx, task := trace.NewTask(context.Background(), "imgconv")
	// TODO: deferでtask.Endを呼ぶ

	if err := imgconv.ConvertAll(ctx, os.Args[1], flagFrom, flagTo); err != nil {
		return err
	}

	return nil
}
