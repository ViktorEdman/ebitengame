package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const gravity = 0.98

type ball struct {
	x, y float32
	vy, xy float32
	radius float32

}

func (b *ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.x, b.y, b.radius, color.White, false)
}

func (b *ball) Update() error {
	b.y += gravity
	return nil
}
