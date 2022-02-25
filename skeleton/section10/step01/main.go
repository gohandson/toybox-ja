package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gohandson/toybox-ja/skeleton/section10/step01/eventwatcher"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	// TODO: 環境変数PORTからポート番号を取得する

	if port == "" {
		port = "8080"
	}
	// TODO: ホストを空、ポートを変数portとしてnet.JoinHostPort関数を呼ぶ


	ew, err := eventwatcher.New(addr)
	if err != nil {
		return err
	}

	if err := ew.Start(); err != nil {
		return err
	}

	return nil
}
