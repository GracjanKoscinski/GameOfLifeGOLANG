package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"time"
)

const (
	screenWidth        = 800
	screenHeight       = 800
	boardSize          = 100
	generationInterval = time.Second
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")

	game := &Game{
		lastGenerationTime: time.Now(),
		state:              "drawing",
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
