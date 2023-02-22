package snake

import (
	"github.com/abdivasiyev/game/internal/config"
	"github.com/abdivasiyev/game/internal/food"
	"github.com/abdivasiyev/game/internal/point"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"math"
	"time"
)

var opposites = map[ebiten.Key]ebiten.Key{
	ebiten.KeyUp:    ebiten.KeyDown,
	ebiten.KeyDown:  ebiten.KeyUp,
	ebiten.KeyLeft:  ebiten.KeyRight,
	ebiten.KeyRight: ebiten.KeyLeft,
}

type Snake struct {
	body           []point.Point
	width, height  float64
	timer          time.Time
	updateInterval time.Duration

	isDead bool

	direction ebiten.Key
}

func New() *Snake {
	return &Snake{
		body: []point.Point{
			{0, 0},
		},
		width:          20,
		height:         20,
		direction:      ebiten.KeyRight,
		timer:          time.Now(),
		updateInterval: 250 * time.Millisecond,
	}
}

func (s *Snake) Head() point.Point {
	return s.body[len(s.body)-1]
}

func (s *Snake) SetInterval(t time.Duration) {
	s.updateInterval = t
}

func (s *Snake) IsDirectionChanged() (ebiten.Key, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		return ebiten.KeyUp, true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		return ebiten.KeyDown, true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		return ebiten.KeyLeft, true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		return ebiten.KeyRight, true
	}

	return 0, false
}

func (s *Snake) AddHead() {
	h := s.Head()
	newHead := point.Point{
		X: h.X,
		Y: h.Y,
	}

	switch s.direction {
	case ebiten.KeyRight:
		newHead.X += s.width
	case ebiten.KeyLeft:
		newHead.X -= s.width
	case ebiten.KeyUp:
		newHead.Y -= s.height
	case ebiten.KeyDown:
		newHead.Y += s.height
	}

	s.body = append(s.body, newHead)
}

func (s *Snake) CheckCollision() {
	// check head is collided with borders
	head := s.Head()
	if head.X < 0 || head.Y < 0 {
		s.isDead = true
		return
	}

	if head.X+s.width > config.ScreenWidth || head.Y+s.height > config.ScreenHeight {
		s.isDead = true
		return
	}

	for i, b := range s.body {
		if i == len(s.body)-1 {
			continue
		}

		if b.X == head.X && b.Y == head.Y {
			s.isDead = true
			return
		}
	}
}

func (s *Snake) IsFeed(f *food.Food) bool {
	h := s.Head()

	x0, y0 := h.X, h.Y
	x1, y1 := h.X+s.width, h.Y+s.height

	fx0, fy0 := f.Position().X, f.Position().Y
	fx1, fy1 := f.Position().X+s.width, f.Position().Y+s.height

	byWidthCollides := math.Min(x1, fx1) > math.Max(x0, fx0)
	byHeightCollides := math.Min(y1, fy1) > math.Max(y0, fy0)

	return byWidthCollides && byHeightCollides
}

func (s *Snake) IsFoodCorrectPosition(f *food.Food) bool {
	for _, b := range s.body {
		if b.X == f.Position().X && b.Y == f.Position().Y {
			return false
		}
	}

	return true
}

func (s *Snake) IsDead() bool {
	return s.isDead
}

func (s *Snake) Update() error {
	s.CheckCollision()

	if key, ok := s.IsDirectionChanged(); ok && opposites[key] != s.direction {
		s.direction = key
	}

	interval := time.Since(s.timer)
	if interval < s.updateInterval {
		return nil
	}
	s.timer = time.Now()

	s.AddHead()

	s.body = s.body[1:]

	return nil
}

func (s *Snake) Draw(screen *ebiten.Image) {
	if s.isDead {
		return
	}

	for i, b := range s.body {
		if i == len(s.body)-1 {
			ebitenutil.DrawCircle(screen, b.X+s.width/2, b.Y+s.height/2, s.width/2, color.RGBA{R: 200, G: 50, B: 150, A: 150})
		} else {
			ebitenutil.DrawRect(screen, b.X, b.Y, s.width, s.height, color.RGBA{R: 200, G: 50, B: 150, A: 150})
		}
	}
}
