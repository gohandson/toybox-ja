package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	useCoin int = 10
)

type status int

const (
	statusInit status = iota
	statusSelect
	statusResult
)

type game struct {
	status  status
	coin    int
	all     []*card
	cards   []*card
	uicards []*uicard
	msg     string
	width   int
	height  int
}

func newGame() *game {
	var g game
	g.reset()
	return &g
}

func (g *game) reset() {
	g.status = statusInit
	g.coin = 100
	g.all = newAllCards()
	g.cards = nil
	g.uicards = make([]*uicard, 5)
	for i := range g.uicards {
		g.uicards[i] = &uicard{
			idx: i,
		}
	}
	g.width = 740
	g.height = 500
}

func (g *game) Update() error {

	switch g.status {
	case statusInit:
		if len(g.all) < len(g.uicards)*2 || g.coin < useCoin {
			g.reset()
		}

		var keys []ebiten.Key
		keys = inpututil.AppendJustReleasedKeys(keys)
		if len(keys) != 1 {
			return nil
		}

		if keys[0] == ebiten.KeySpace {
			g.all, g.cards = draw(g.all, len(g.uicards))
			for i := range g.uicards {
				g.uicards[i].card = g.cards[i]
			}
			fmt.Println(g.cards)
			g.msg = "Select discard cards"
			g.status = statusSelect
		}
	case statusSelect:

		var keys []ebiten.Key
		keys = inpututil.AppendJustReleasedKeys(keys)
		if len(keys) != 1 {
			return nil
		}

		if keys[0] == ebiten.KeySpace {
			var idx []int
			for i := range g.uicards {
				if g.uicards[i].selected {
					idx = append(idx, i)
				}
				g.uicards[i].selected = false
			}
			g.all, g.cards = change(g.all, g.cards, idx)
			for i := range g.uicards {
				g.uicards[i].card = g.cards[i]
			}
			fmt.Println(g.cards)

			hand := judge(g.cards)
			g.msg = hand.String()
			g.coin += -useCoin + hand.Ratio()*useCoin

			g.status = statusResult
			return nil
		}

		kn, err := strconv.ParseInt(strings.TrimPrefix(keys[0].String(), "Digit"), 10, 64)
		if err != nil || kn <= 0 || int(kn) > len(g.uicards) {
			return nil
		}

		idx := int(kn) - 1
		g.uicards[idx].selected = !g.uicards[idx].selected
	case statusResult:
		var keys []ebiten.Key
		keys = inpututil.AppendJustReleasedKeys(keys)
		if len(keys) != 1 {
			return nil
		}

		if keys[0] == ebiten.KeySpace {
			g.status = statusInit
		}
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	msg := fmt.Sprintf("coin %5d", g.coin)
	if g.msg != "" {
		msg += "\n" + g.msg
	}
	ebitenutil.DebugPrint(screen, msg)
	for _, card := range g.uicards {
		card.draw(screen)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
