package main

// 役
type hand int

const (
	handNothing hand = iota
	handOnePair
	handTwoPair
	handThreeCards
	handStraight
	handFlash
	handFullHouse
	handFourCards
	handStraightFlash
	handRoyalFlash
)

func (h hand) String() string {
	switch h {
	case handRoyalFlash:
		return "Royal Flash"
	case handStraightFlash:
		return "Straight Flash"
	case handFourCards:
		return "Four Cards"
	case handFullHouse:
		return "Full House"
	case handFlash:
		return "Flash"
	case handStraight:
		return "Straight"
	case handThreeCards:
		return "Three Cards"
	case handTwoPair:
		return "Two Pair"
	case handOnePair:
		return "One Pair"
	default:
		return ""
	}

}

func (h hand) Ratio() int {
	switch h {
	case handRoyalFlash:
		return 100
	case handStraightFlash:
		return 50
	case handFourCards:
		return 20
	case handFullHouse:
		return 7
	case handFlash:
		return 5
	case handStraight:
		return 4
	case handThreeCards:
		return 3
	case handTwoPair:
		return 2
	case handOnePair:
		return 1
	default:
		return 0
	}
}

func judge(cards []*card) hand {
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

	switch {
	case isStraight && isFlash && cards[0].number == 10:
		return handRoyalFlash
	case isStraight && isFlash:
		return handStraightFlash
	case maxSame == 4:
		return handFourCards
	case len(numCount) == 2:
		return handFullHouse
	case isFlash:
		return handFlash
	case isStraight:
		return handStraight
	case maxSame == 3:
		return handThreeCards
	case len(numCount) == 3:
		return handTwoPair
	case len(numCount) == 4:
		return handOnePair
	default:
		return handNothing
	}
}
