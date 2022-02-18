package main

import "fmt"

func start(converters []converter) {
	for {
		// 単位変換器を選ぶ
		no, ok := inputConverterNo(converters)
		if !ok {
			return
		}

		c := converters[no]
		from := inputValue(c)
		to := c.calc(from)
		fmt.Printf("%.2f[%s] -> %.2f[%s]\n\n", from, c.fromUnit, to, c.toUnit)
	}
}

// 単位変換器を選ぶ関数
// 引数: converters（単位変換器）
// 戻り値：選んだ変換器の添字と処理を継続するかどうか
func inputConverterNo(converters []converter) (int, bool) {
	fmt.Println("以下の単位変換器から選んでください。")
	for i, c := range converters {
		fmt.Printf("%d: %s\n", i+1, c.name)
	}
	fmt.Printf("%d: 終了\n", len(converters)+1)

	var no int
	for no <= 0 || no > len(converters)+1 {
		fmt.Printf("1〜%dの数字>", len(converters)+1)
		fmt.Scanln(&no)
	}

	// 終了する
	if no == len(converters)+1 {
		// 終了する場合は0とfalseを返す
		return 0, false
	}

	return no - 1, true
}

func inputValue(c converter) float64 {
	var v float64
	fmt.Println(c.name)
	fmt.Printf("変換する値[%s]>", c.fromUnit)
	fmt.Scanln(&v)
	return v
}
