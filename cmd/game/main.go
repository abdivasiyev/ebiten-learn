package main

import (
	"github.com/abdivasiyev/game/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.New()
	if err != nil {
		panic(err)
	}
	ebiten.SetInitFocused(true)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
