package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: string型をベースにしてsuit型を定義する

const (
	suitHeart   suit = "♥"
	suitClub    suit = "♣"
	suitDiamond suit = "◆"
	// TODO: ♠を表す定数suitSpadeを定義する
)

type card struct {
	suit   suit
	// TODO: int型のnumberフィールドを定義する
}

func main() {

	suits := []suit{
		suitHeart,
		suitClub,
		suitDiamond,
		suitSpade,
	}

	// 山札を作る
	all := make([]card, 0, 13*4)
	for _, s := range suits {
		for n := 2; n <= 14; n++ {
			all = append(all, card{
				// TODO: マークをセットする
				number: n,
			})
		}
	}

	// 乱数の種をセットする
	t := time.Now().UnixNano()
	rand.Seed(t)

	// 山札をシャッフルさせる
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	// 手札
	cards := all[:5]
	all = all[5:]

	// 手札を表示させる
	// cardsの要素を1つずつ取り出し変数cに入れる
	for _, c := range cards {
		// "♥ "のように出力する
		fmt.Print(c.suit, " ")
		switch c.number {
		case 11:
			fmt.Println("J")
		case 12:
			fmt.Println("Q")
		case 13:
			fmt.Println("K")
		case 14:
			fmt.Println("A")
		default:
			// TODO: 番号を改行ありで出力する
		}
	}
}
