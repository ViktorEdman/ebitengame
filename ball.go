package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ball struct {
	x, y float64
	radius float64

}

func (b *ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(b.x), float32(b.y), float32(b.radius), color.White, false)
}

func (b *ball) Update() error {
	return nil
}
