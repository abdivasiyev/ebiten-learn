package ball

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Ball struct {
	x  float64
	y  float64
	r  float64
	dx float64
	dy float64

	maxX float64
	maxY float64

	color color.Color
}

func New(x, y, r, dx, dy, maxX, maxY float64, color color.Color) *Ball {
	return &Ball{
		x:     x,
		y:     y,
		r:     r,
		dx:    dx,
		dy:    dy,
		maxX:  maxX,
		maxY:  maxY,
		color: color,
	}
}

func (b *Ball) Update() {
	if b.x+b.dx < b.r || b.x+b.dx > b.maxX-b.r {
		b.dx = -b.dx
	}

	if b.y+b.dy < b.r || b.y+b.dy > b.maxY-b.r {
		b.dy = -b.dy
	}

	b.x += b.dx
	b.y += b.dy
}

func (b *Ball) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, b.x, b.y, b.r, b.color)
}
