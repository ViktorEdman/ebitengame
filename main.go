package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	currentFrame int
	paused bool
	entities []entity
}

type entity interface {
	Draw(screen *ebiten.Image)
	Update() error
}

const baseRes = 128

func (g *Game) Update() error {
	g.handleInputs()
	if g.paused {
		return nil
	}
	g.currentFrame++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World! Frame %d", g.currentFrame))
	for _, e := range g.entities {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 4*baseRes, 3*baseRes
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
