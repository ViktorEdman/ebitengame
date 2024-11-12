package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handleInputs() {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		g.paused = !g.paused
	}
	if g.paused {
		return
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		g.entities = append(g.entities, &ball{float64(x), float64(y), 2})
	}
} 
