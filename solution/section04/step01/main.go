package main

import "fmt"

func main() {
	// 関数startを呼び出す
	start()
}

// startという関数を定義
func start() {
	fmt.Println("摂氏[°C] -> 華氏[°F]")
	var from float64
	fmt.Print("変換する値[°C]>")
	fmt.Scanln(&from)

	to := from*1.8 + 32
	fmt.Printf("結果: %.2f[°F]\n", to)
}
