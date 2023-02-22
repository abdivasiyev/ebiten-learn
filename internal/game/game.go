package game

import (
	"fmt"
	"github.com/abdivasiyev/game/internal/food"
	"github.com/abdivasiyev/game/internal/snake"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math/rand"
	"time"
)

type Params struct {
	Width, Height   int
	BackgroundColor color.Color
	Snake           *snake.Snake
}

type myGame struct {
	width           int
	height          int
	backgroundColor color.Color
	snake           *snake.Snake
	food            *food.Food
	score           int64
}

func New(p Params) (ebiten.Game, error) {
	rand.Seed(time.Now().UnixNano())
	return &myGame{
		width:           p.Width,
		height:          p.Height,
		backgroundColor: p.BackgroundColor,
		snake:           p.Snake,
	}, nil
}

func (g *myGame) newFood() {
	if g.food == nil {
		g.food = food.Generate(g.width, g.height)
		for !g.snake.IsFoodCorrectPosition(g.food) {
			g.food = food.Generate(g.width, g.height)
		}

		return
	}

	if g.snake.IsFeed(g.food) {
		log.Println("eaten")
		g.snake.AddHead()
		g.food = food.Generate(g.width, g.height)
		for !g.snake.IsFoodCorrectPosition(g.food) {
			g.food = food.Generate(g.width, g.height)
		}
		g.score++

		switch {
		case g.score > 5:
			g.snake.SetInterval(200 * time.Millisecond)
		case g.score > 10:
			g.snake.SetInterval(100 * time.Millisecond)
		case g.score > 20:
			g.snake.SetInterval(50 * time.Millisecond)
		}
	}
}

func (g *myGame) Update() error {
	_ = g.snake.Update()

	g.newFood()

	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(g.backgroundColor)

	// draw snake
	if g.snake.IsDead() {
		ebitenutil.DebugPrint(screen, "Game Over")
		return
	}

	// draw food
	g.food.Draw(screen)
	g.snake.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Your score: %d", g.score))
}

func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
