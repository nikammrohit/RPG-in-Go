package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	// Values which are passed into Game
	PlayerImage *ebiten.Image //uppercase P is a public field
	X, Y        float64
}

func (g *Game) Update() error {
	// React to keypressed
	SPEED := 2
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.X += float64(SPEED)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.X -= float64(SPEED)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Y -= float64(SPEED)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Y += float64(SPEED)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 250, 255})
	//ebitenutil.DebugPrint(screen, "Hello, World!")

	// Draw player
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.X, g.Y) // Allows us to translate player for movement

	screen.DrawImage(
		g.PlayerImage.SubImage( // grab vector values for a single frame from spritesheet
			image.Rect(0, 0, 13, 18), // start(inclusive), end(exclusive)
		).(*ebiten.Image), // convs image.Image to *ebiten.Image which we want
		&opts, // img options
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize() // Let surface be same size as window
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled) // Lets us resize our window

	playerImg, _, err := ebitenutil.NewImageFromFile("./assets/Characters/_Source/Human/TRIMMED_walk.png")
	if err != nil {
		// Print err then exit game
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{PlayerImage: playerImg, X: 100, Y: 100}); err != nil {
		log.Fatal(err)
	}
}
