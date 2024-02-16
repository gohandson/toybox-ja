package main

import (
	"fmt"
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type game struct {
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
	var keys []ebiten.Key
	// TODO: ちょうどキーから離したものを取得

	if len(keys) != 1 {
		return nil
	}
	
	if /* TODO: スペースキーだった場合 */ {
		g.all, g.cards = draw(g.all, len(g.uicards))
		for i := range g.uicards {
			g.uicards[i].card = g.cards[i]
		}
		fmt.Println(g.cards)
		g.msg = "Cards are drawn"
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
