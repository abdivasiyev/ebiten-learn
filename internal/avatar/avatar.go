package avatar

import (
	"bytes"
	"github.com/abdivasiyev/game/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

type Avatar struct {
	horizontalFrames int
	verticalFrames   int
	frameCount       int
	startX           int
	startY           int
	width            int
	height           int

	img *ebiten.Image
}

func New(startX, startY, horizontalFrames, verticalFrames int, width, height int, img []byte) (*Avatar, error) {
	a := &Avatar{
		horizontalFrames: horizontalFrames,
		verticalFrames:   verticalFrames,
		frameCount:       0,
		startX:           startX,
		startY:           startY,
		width:            width,
		height:           height,
	}

	im, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		return a, err
	}

	a.img = ebiten.NewImageFromImage(im)

	return a, nil
}

func (a *Avatar) Update() {
	a.frameCount++
	if a.frameCount > config.MaxCount {
		a.frameCount = 0
	}
}

func (a *Avatar) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(a.width)/2, -float64(a.height)/2)
	op.GeoM.Translate(config.ScreenWidth/2, config.ScreenHeight/2)

	totalFrames := a.horizontalFrames * a.verticalFrames

	i := (a.frameCount / totalFrames) % a.horizontalFrames
	j := (a.frameCount / totalFrames) % a.verticalFrames

	sx, sy := a.startX+i*a.width, a.startY+j*a.height

	screen.DrawImage(
		a.img.SubImage(
			image.Rect(sx, sy, sx+a.width, sy+a.height),
		).(*ebiten.Image),
		op,
	)
}
