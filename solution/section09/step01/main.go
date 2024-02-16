package main

import (
	"fmt"
	"os"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	g := newGame()
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle("Poker")
	if err := ebiten.RunGame(g); err != nil {
		return err
	}

	return nil
}
