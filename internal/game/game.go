package game

import (
	"github.com/abdivasiyev/game/internal/avatar"
	"github.com/abdivasiyev/game/internal/config"
	avImg "github.com/abdivasiyev/game/res/avatar"
	"github.com/hajimehoshi/ebiten/v2"
)

type myGame struct {
	width  int
	height int

	avatar *avatar.Avatar
}

func New() (ebiten.Game, error) {
	g := &myGame{
		width:  config.ScreenWidth,
		height: config.ScreenHeight,
	}

	av, err := avatar.New(280, 100, 4, 2, 610, 900, avImg.Girl)
	if err != nil {
		return nil, err
	}
	g.avatar = av
	return g, nil
}

func (g *myGame) Update() error {
	g.avatar.Update()

	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	screen.Clear()

	g.avatar.Draw(screen)
}

func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
