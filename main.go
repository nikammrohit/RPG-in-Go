package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	// Image and position variables for our player
	Img  *ebiten.Image //spritesheet
	X, Y float64
}
type Game struct {
	// Values which are passed into Game
	player *Sprite // pointer to sprite struct
}

func (g *Game) Update() error {
	// React to keypressed
	SPEED := 2
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.X += float64(SPEED)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.X -= float64(SPEED)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Y -= float64(SPEED)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.Y += float64(SPEED)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 250, 255})
	//ebitenutil.DebugPrint(screen, "Hello, World!")

	// Draw player
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y) // Allows us to translate player for movement

	screen.DrawImage(
		g.player.Img.SubImage( // grab vector values for a single frame from spritesheet
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

	game := Game{
		player: &Sprite{
			Img: playerImg,
			X:   50.0,
			Y:   50.0,
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
