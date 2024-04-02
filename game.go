package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

type Game struct {
	lastGenerationTime time.Time
	state              string
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.state != "playing" {
		x, y := ebiten.CursorPosition()
		boardX, boardY := x/(screenWidth/boardSize), y/(screenHeight/boardSize)
		board[boardX][boardY] = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) && g.state != "playing" {
		g.state = "playing"
	}

	if g.state == "playing" && time.Since(g.lastGenerationTime) >= generationInterval {
		nextGeneration()
		g.lastGenerationTime = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] {
				screen.Set(i*(screenWidth/boardSize), j*(screenHeight/boardSize), color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
