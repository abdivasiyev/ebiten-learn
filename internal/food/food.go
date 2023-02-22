package food

import (
	"github.com/abdivasiyev/game/internal/point"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
)

type Food struct {
	position point.Point
	color    color.Color
}

func New(x, y float64) *Food {
	return &Food{
		position: point.Point{X: x, Y: y},
		color:    color.RGBA{R: 200, G: 200, B: 50, A: 150},
	}
}

func Generate(width, height int) *Food {
	foodX, foodY := rand.Int63n(int64(width-20)), rand.Int63n(int64(height-20))

	return New(float64(foodX), float64(foodY))
}

func (f *Food) Position() point.Point {
	return f.position
}

func (f *Food) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, f.position.X, f.position.Y, 20, 20, f.color)
}
