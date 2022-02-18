package main

import "fmt"

func main() {
	start()
}

func start() {
	// TODO: 値を入力する関数を呼び出し変数fromに代入

	to := from*1.8 + 32
	fmt.Printf("結果: %.2f[°F]\n", to)
}

// TODO: 値を入力する関数を定義する
// 引数：name（値の名前）, fromUnit（変換元の単位）
// 戻り値：入力された値
