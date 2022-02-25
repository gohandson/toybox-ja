package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tenntenn/connpass"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	params, err := connpass.SearchParam(connpass.Keyword("golang"))
	if err != nil {
		return err
	}

	// TODO: connpassのAPIクライアントを生成する

	ctx := context.Background()
	result, err := cli.Search(ctx, params)
	if err != nil {
		return err
	}

	for _, e := range result.Events {
		// TODO: イベントのタイトルを出力する

	}

	return nil
}
