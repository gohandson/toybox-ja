package main

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"slices"
)

type suit int

const (
	suitHeart suit = iota
	suitClub
	suitDiamond
	suitSpade
)

func (s suit) String() string {
	return map[suit]string{
		suitHeart:   "heart",
		suitClub:    "club",
		suitDiamond: "diamond",
		suitSpade:   "spade",
	}[s]
}

type card struct {
	suit   suit
	number int
}

func (c *card) String() string {
	s := map[suit]string{
		suitHeart:   "♥",
		suitClub:    "♣",
		suitDiamond: "◆",
		suitSpade:   "♠",
	}[c.suit]
	return fmt.Sprintf("%s%d", s, c.number)
}

// 山札を作ります
func newAllCards() []*card {
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
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	return all
}

// 山札からカードを引く
func draw(all []*card, count int) ([]*card, []*card) {
	cards := all[:count]
	all = all[count:]
	slices.SortFunc(cards, func(a, b *card) int {
		return cmp.Compare(a.number, b.number)
	})
	return all, cards[0:count:count]
}

// 手札からカードを捨てます
func discard(cards []*card, idx []int) []*card {
	count := len(cards) - len(idx)
	if count < 0 {
		return nil
	}

	remains := make([]*card, 0, count)
	for i := range cards {
		if !slices.Contains(idx, i) {
			remains = append(remains, cards[i])
		}
	}

	return remains
}

func change(all, cards []*card, idx []int) ([]*card, []*card) {
	cards = discard(cards, idx)
	all, newCards := draw(all, len(idx))
	cards = append(cards, newCards...)
	slices.SortFunc(cards, func(a, b *card) int {
		return cmp.Compare(a.number, b.number)
	})
	return all, cards
}
