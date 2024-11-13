package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const gravity = 1

type ball struct {
	x, y float32
	vx, vy float32
	radius float32

}

func (b *ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.x, b.y, b.radius, color.White, true)
}

func (b *ball) Update() error {
	b.y += b.vy
	b.x += b.vx
	b.vy += gravity
	b.Bounce(SCREEN_WIDTH, SCREEN_HEIGHT)
	return nil
}

func (g *Game) addBalls(number, x, y int) {
	radius := 20
	for i := 0; i < number; i++ {

	g.entities = append(g.entities, &ball{float32(x), float32(y), (rand.Float32()-0.5)*30, float32(gravity), float32(radius)})
	}
}

func (b *ball) Bounce(width, height int) {
	if math.Abs(float64(b.vy)) < 1.5 && b.y + b.radius*2 > SCREEN_HEIGHT {
		b.vy = 0
		b.y = SCREEN_HEIGHT - b.radius
		
	}
	if math.Abs(float64(b.vx)) < 10 && (b.x  > SCREEN_WIDTH || b.x  < 0) {
		b.vx = 0
		
		if b.x - b.radius> SCREEN_WIDTH - b.radius{
			b.x = SCREEN_WIDTH - b.radius
		}
		if b.x + b.radius< 0  + b.radius{
			b.x = 0 + b.radius
		}
		
	}
	if (b.x - b.radius )< 0 || (b.x + b.radius )> float32(width)  {
		b.vx *= -0.85
	}
	if b.y-b.radius < 0 || b.y+b.radius > float32(height) {
		b.vy *= -0.85
	}
}
