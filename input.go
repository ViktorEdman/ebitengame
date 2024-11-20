package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handleInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.debugEnabled = !g.debugEnabled
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}
	if g.paused {
		return
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.addBall(float64(x), float64(y))
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		g.clearEntities()
	}
}

func (g *Game) clearEntities() {
	g.entities = []entity{}
}
