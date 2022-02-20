// TODO: cmd/imgconv/main.goに移動する
package main

import (
	"flag"
	"fmt"
	"os"

	// TODO: imgconvパッケージをインポートする
)

var (
	flagTo   = PNG   // TODO: パッケージ名をつける
	flagFrom = JPEG  // TODO: パッケージ名をつける
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	// TODO: convertAllはパッケージ名をつけてエクスポートされたものに変える
	if err := convertAll(os.Args[1], flagFrom, flagTo); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
