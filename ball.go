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
	x, y   float64
	vx, vy float64
	radius float64
	g      *Game
}

func (b *ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(b.x), float32(b.y), float32(b.radius), color.White, true)
}

func (b *ball) Update() error {
	b.y += b.vy
	b.x += b.vx
	b.vy += gravity
	b.ResolveCollisions()
	b.Bounce(SCREEN_WIDTH, SCREEN_HEIGHT)
	return nil
}

func (g *Game) addBall(x, y float64) {
	radius := 20

	g.entities = append(g.entities, &ball{x, y, (rand.Float64() - 0.5) * 30, gravity, float64(radius), g})
}

func (b *ball) ResolveCollisions() {
	for i := range b.g.entities {
		ball, ok := b.g.entities[i].(*ball)
		if !ok {
			continue
		}
		if ball == b {
			continue
		}
		resolveCollision(b, ball)
	}
}

func resolveCollision(b1, b2 *ball) {
	dx := b1.x - b2.x
	dy := b1.y - b2.y
	distance := math.Sqrt(dx*dx + dy*dy)
	if distance < b1.radius+b2.radius {
		nx, ny := dx/distance, dy/distance
		relativeVelocity := (b2.vx-b1.vx)*nx + (b2.vy-b1.vy)*ny
		if relativeVelocity > 0 {
			impulse := relativeVelocity * 0.9
			b1.vx += impulse * nx
			b1.vy += impulse * ny
			b2.vx -= impulse * nx
			b2.vy -= impulse * ny
		}
	}
}

func (b *ball) Bounce(width, height int) {
	if math.Abs(float64(b.vy)) < 1.5 && b.y+b.radius*2 > SCREEN_HEIGHT {
		b.vy = 0
		b.y = SCREEN_HEIGHT - b.radius

	}
	if math.Abs(float64(b.vx)) < 10 && (b.x > SCREEN_WIDTH || b.x < 0) {
		b.vx = 0

		if b.x-b.radius > SCREEN_WIDTH-b.radius {
			b.x = SCREEN_WIDTH - b.radius
		}
		if b.x+b.radius < 0+b.radius {
			b.x = 0 + b.radius
		}

	}
	if (b.x-b.radius) < 0 || (b.x+b.radius) > float64(width) {
		b.vx *= -0.85
	}
	if b.y-b.radius < 0 || b.y+b.radius > float64(height) {
		b.vy *= -0.85
	}
}
