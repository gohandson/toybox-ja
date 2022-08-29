package main

import (
	"fmt"
)

func main() {
	// 山札を作る
	// TODO: 長さが0で容量が13のintスライスの変数allを作成する

	for n := 2; n <= 14; n++ { // 14がAを表す
		// TODO: スライスallの末尾に変数nの値を追加する
	}

	// 山札を表示させる
	for i, n := range all {
		// TODO: " 1番目: "のように出力する

		switch n {
		case 11:
			fmt.Println("J")
		case 12:
			fmt.Println("Q")
		case 13:
			fmt.Println("K")
		case 14:
			fmt.Println("A")
		default:
			fmt.Println(n)
		}
	}
}
