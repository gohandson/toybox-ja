package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	const url = "https://connpass.com/api/v1/event/?keyword=golang"
	// TODO: GETメソッドのリクエストを生成する

	if err != nil {
		return err
	}

	// TODO: http.DefaultClientでリクエストを送る

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 標準出力にダンプする
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		return err
	}

	return nil
}
