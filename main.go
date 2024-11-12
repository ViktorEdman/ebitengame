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

const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 600


func (g *Game) Update() error {
	g.handleInputs()
	if g.paused {
		return nil
	}
	g.currentFrame++
	for i := range g.entities {
		g.entities[i].Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Frame %6d, %d entities", g.currentFrame, len(g.entities)))
	
	ebitenutil.DebugPrintAt(screen, getMemUsage(), 0, 20)
	for i := range g.entities {
		g.entities[i].Draw(screen)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Entity %d: %+v", i, g.entities[i]), 0, 30+(i*12))	
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth = SCREEN_WIDTH
	screenHeight = SCREEN_HEIGHT
	return
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
