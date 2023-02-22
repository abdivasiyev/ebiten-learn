package main

import (
	"github.com/abdivasiyev/game/internal/config"
	"github.com/abdivasiyev/game/internal/game"
	"github.com/abdivasiyev/game/internal/snake"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	s := snake.New()

	g, err := game.New(game.Params{
		Width:           config.ScreenWidth,
		Height:          config.ScreenHeight,
		BackgroundColor: config.BackgroundColor,
		Snake:           s,
	})
	if err != nil {
		panic(err)
	}
	ebiten.SetInitFocused(true)
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
