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

	for /* TODO: countを1から2まで1ずつ増やしながら繰り返す */ {
		var answer int
		fmt.Print("回答>")
		fmt.Scanln(&answer)

		switch {
		// TODO: answerが2の時（正解の場合）
			fmt.Println("正解!")
			// 1度目で正解しても2回答える必要がある
		// TODO: 1度目のチャレンジで不正解の場合
			fmt.Println("不正解!")
			fmt.Println("もう一度チャレンジ!")
		// TODO: それ以外（2度目のチャレンジで不正解の場合）
			fmt.Println("不正解!")
			fmt.Println("答えは2です")
		}
	}
}
