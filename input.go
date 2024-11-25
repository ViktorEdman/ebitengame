package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handleInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.turboMode = !g.turboMode
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.debugEnabled = !g.debugEnabled
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}
	if g.paused {
		return
	}
	switch g.turboMode {
	case true:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			g.addBall(float64(x), float64(y))
		}
	case false:
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			g.addBall(float64(x), float64(y))
		}
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		g.clearEntities()
	}
	_, dy := ebiten.Wheel()
	if g.nextBallSize >= 10 {
		g.nextBallSize += int(dy / 10)
	}
	if g.nextBallSize < 10 {
		g.nextBallSize = 10
	}
}

func (g *Game) clearEntities() {
	g.entities = []entity{}
}
