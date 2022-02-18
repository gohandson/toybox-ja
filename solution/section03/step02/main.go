package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 山札を作ります
	all := make([]int, 0, 13)
	for n := 2; n <= 14; n++ {
		all = append(all, n)
	}

	// 乱数の種をセットする
	t := time.Now().UnixNano()
	rand.Seed(t)

	// 山札をシャッフルさせる
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	// 山札の前方5枚を手札としcardsに入れる
	cards := all[:5]
	// 6枚目以降を新しい山札とする
	all = all[5:]

	// 手札を表示させます
	for i, n := range cards {
		fmt.Printf("%d番目: ", i+1)
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
