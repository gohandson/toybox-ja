package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type suit int

const (
	suitHeart suit = iota
	suitClub
	suitDiamond
	suitSpade
)

type card struct {
	suit   suit
	number int
}

func main() {
	marks := map[suit]string{
		suitHeart:   "♥",
		suitClub:    "♣",
		suitDiamond: "◆",
		suitSpade:   "♠",
	}

	// 山札を作ります
	all := make([]*card, 0, 13*4)
	for s := suitHeart; s <= suitSpade; s++ {
		for n := 2; n <= 14; n++ {
			all = append(all, &card{
				suit:   s,
				number: n,
			})
		}
	}
	// 山札をシャッフルさせます
	now := time.Now()
	rndsrc := rand.NewSource(now.UnixNano())
	rnd := rand.New(rndsrc)
	rnd.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	coin := 100
	for coin > 0 && len(all) > 5 {
		// 使用するコインの枚数
		var useCoin int
		for {
			fmt.Printf("コインを何枚かけますか？（最大%d枚）\n", coin)
			fmt.Printf(">")
			fmt.Scanln(&useCoin)
			if useCoin > 0 && useCoin <= coin {
				break
			}
			fmt.Println("正しいコイン枚数を入れてください")
		}

		// 手札
		cards := all[:5]
		all = all[5:]
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].number < cards[j].number
		})

		// 手札の表示
		fmt.Println("手札")
		for _, c := range cards {
			fmt.Print(marks[c.suit], " ")
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
				fmt.Println(c.number)
			}
		}

		// 残す枚数
		var remains int
		for {
			fmt.Println("何枚残しますか？（最大5枚）")
			fmt.Printf(">")
			fmt.Scanln(&remains)
			if remains >= 0 && remains <= 5 {
				break
			}
			fmt.Println("0以上5以下です")
		}

		cards = append(cards[:remains], all[:5-remains]...)
		all = all[5-remains:]
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].number < cards[j].number
		})

		// 手札の表示
		fmt.Println("手札")
		for _, c := range cards {
			fmt.Print(marks[c.suit], " ")
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
				fmt.Println(c.number)
			}
		}

		numCount := make(map[int]int)
		var maxSame int
		isStraight := true
		isFlash := true
		for i := 0; i < len(cards); i++ {
			numCount[cards[i].number]++
			if maxSame < numCount[cards[i].number] {
				maxSame = numCount[cards[i].number]
			}

			if i > 0 {
				isStraight = isStraight && cards[i].number-cards[i-1].number == 1
				isFlash = isFlash && cards[i].suit == cards[i-1].suit
			}
		}

		var ratio int
		switch {
		case isStraight && isFlash && cards[0].number == 10:
			fmt.Println("ロイヤルフラッシュ")
			ratio = 100
		case isStraight && isFlash:
			fmt.Println("ストレートフラッシュ")
			ratio = 50
		case maxSame == 4:
			fmt.Println("フォーカード")
			ratio = 20
		case len(numCount) == 2:
			fmt.Println("フルハウス")
			ratio = 7
		case isFlash:
			fmt.Println("フラッシュ")
			ratio = 5
		case isStraight:
			fmt.Println("ストレート")
			ratio = 4
		case maxSame == 3:
			fmt.Println("スリーカード")
		case len(numCount) == 3:
			fmt.Println("ツーペア")
			ratio = 2
		case len(numCount) == 4:
			fmt.Println("ワンペア")
			ratio = 1
		default:
			fmt.Println("役無し")
		}

		increase := useCoin * ratio
		afterCoin := coin - useCoin + increase
		fmt.Printf("%d * %d = %d\n", useCoin, ratio, increase)
		fmt.Printf("手持ちコイン: %d -> %d\n", coin, afterCoin)
		coin = afterCoin
	}
}
