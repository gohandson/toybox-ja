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

	var answer int
	fmt.Print("回答>")
	fmt.Scanln(&answer)
	
	if /* TODO: answerが2の場合 */ {
		fmt.Println("正解!")
	} /* それ以外の場合 */ {
		fmt.Println("不正解!")
		fmt.Println("答えは2です")
	}
}
