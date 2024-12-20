package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	currentFrame int
	paused       bool
	debugEnabled bool
	turboMode    bool
	entities     []entity
	nextBallSize int
}

type entity interface {
	Draw(screen *ebiten.Image)
	Update() error
	ResolveCollisions()
}

const SCREEN_WIDTH = 1200
const SCREEN_HEIGHT = 1200

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
	if g.debugEnabled {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Frame %6d, %d entities", g.currentFrame, len(g.entities)))
		ebitenutil.DebugPrintAt(screen, getMemUsage(), 0, 10)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.2f FPS", ebiten.ActualFPS()), 0, 20)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.2f TPS", ebiten.ActualTPS()), 0, 30)
	}
	if g.turboMode {
		ebitenutil.DebugPrintAt(screen, "Turbo mode activated", SCREEN_WIDTH-(SCREEN_WIDTH/20), 0)
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Current ball size %d", g.nextBallSize), SCREEN_WIDTH-(SCREEN_WIDTH/5), 20)
	for i := range g.entities {
		g.entities[i].Draw(screen)
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
	if err := ebiten.RunGame(&Game{nextBallSize: 20}); err != nil {
		log.Fatal(err)
	}
}
