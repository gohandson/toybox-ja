package main

import (
	"embed"
	"fmt"
	"image"
	"os"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

//go:embed imgs/*.png
var imgs embed.FS

type uicard struct {
	idx      int
	card     *card
	selected bool
}

func (c *uicard) draw(screen *ebiten.Image) {

	cardimg, err := c.cardimg()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	w, _ := cardimg.Size()

	x := float64(10 + (5+w)*c.idx)
	y := float64(70)

	if c.selected {
		bg, err := c.selectedimg()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}

		var opbg ebiten.DrawImageOptions
		opbg.GeoM.Translate(x-1, y-1)
		screen.DrawImage(bg, &opbg)
	}

	var opcard ebiten.DrawImageOptions
	opcard.GeoM.Translate(x, y)
	screen.DrawImage(cardimg, &opcard)
}

func (c *uicard) cardimg() (*ebiten.Image, error) {
	fname := "imgs/card_back.png"
	if c.card != nil {
		fname = fmt.Sprintf("imgs/card_%s_%02d.png", c.card.suit, c.card.number)
	}

	file, err := imgs.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func (c *uicard) selectedimg() (*ebiten.Image, error) {
	fname := "imgs/card_selected.png"
	file, err := imgs.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
