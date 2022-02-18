package main

import "fmt"

func main() {

	fmt.Println("Q. 以下のコードコンパイルして実行するとどうなるか？")
	fmt.Println("package main")
	fmt.Println("func main() {")
	fmt.Println("	true := false")
	fmt.Println("	println(true == false)")
	fmt.Println("}")

	fmt.Println("1: コンパイルエラー")
	fmt.Println("2: trueと表示される")
	fmt.Println("3: falseと表示される")
	fmt.Println("4: パニックが起きる")

// TODO: 繰り返しにLOOPというラベルをつける

	for count := 1; count <= 2; count++ {
		var answer int
		for {
			fmt.Print("回答>")
			fmt.Scanln(&answer)
			if answer >= 1 && answer <= 4 {
				break
			}
			fmt.Println("1から4で入力してえください")
		}

		switch {
		case answer == 2:
			fmt.Println("正解!")
			// TODO: ラベルLOOPのついた繰り返しを抜け出す

		case count == 1:
			fmt.Println("不正解!")
			fmt.Println("もう一度チャレンジ!")
		default:
			fmt.Println("不正解!")
			fmt.Println("答えは2です")
		}
	}
}
