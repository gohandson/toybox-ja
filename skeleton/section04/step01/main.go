package main

import "fmt"

func main() {
	// TODO: 以下の処理をstart関数を定義し移動させる
	fmt.Println("摂氏[°C] -> 華氏[°F]")
	var from float64
	fmt.Print("変換する値[°C]>")
	fmt.Scanln(&from)

	to := from*1.8 + 32
	fmt.Printf("結果: %.2f[°F]\n", to)

	// TODO: 関数startを呼び出す
}

// TODO: startという関数を定義
