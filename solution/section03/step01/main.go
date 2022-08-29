package main

import (
	"fmt"
)

func main() {
	// 山札を作る
	// 長さが0で容量が13のintのスライスの変数allを作成する
	all := make([]int, 0, 13)
	for n := 2; n <= 14; n++ { // 14がAを表す
		// スライスallの末尾に変数nの値を追加する
		all = append(all, n)
	}

	// 山札を表示させる
	for i, n := range all {
		// " 1番目: "のように出力する
		fmt.Printf("%2d番目: ", i+1)
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
