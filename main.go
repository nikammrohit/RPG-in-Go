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

type Player struct {
	*Sprite
	Health uint
}

type Enemy struct {
	*Sprite
	followsPlayer bool
}

type Game struct {
	// Values which are passed into Game
	player  *Player  // pointer to sprite struct
	enemies []*Enemy // slice containing list of all sprites in game (except for player)
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

	// Algorithm to make enemies follow player
	ENEMY_SPEED := 0.5
	for _, sprite := range g.enemies {
		if sprite.followsPlayer { // only sprites with followsPlayer := true property will follow player
			if sprite.X < g.player.X {
				sprite.X += float64(ENEMY_SPEED)
			} else if sprite.X > g.player.X {
				sprite.X -= float64(ENEMY_SPEED)
			}
			if sprite.Y < g.player.Y {
				sprite.Y += float64(ENEMY_SPEED)
			} else if sprite.Y > g.player.Y {
				sprite.Y -= float64(ENEMY_SPEED)
			}
		}
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
	opts.GeoM.Reset()

	// Iterate through sprites[] slice and draw them
	for _, sprite := range g.enemies {
		// essentially copying draw code for player but for all sprites
		opts.GeoM.Translate(sprite.X, sprite.Y)
		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(42, 24, 55, 40), //! Change coords for enemies
			).(*ebiten.Image),
			&opts,
		)
		opts.GeoM.Reset() // Reset translation back to (0, 0) for next entity
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize() // Let surface be same size as window
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled) // Lets us resize our window

	// Load player spritesheet from file
	playerImg, _, err := ebitenutil.NewImageFromFile("./assets/Characters/_Source/Human/TRIMMED_idle.png")
	//! Insert enemyImg
	if err != nil {
		// Print err then exit game
		log.Fatal(err)
	}

	// Load enemy spritesheet from file
	skeletonImg, _, err := ebitenutil.NewImageFromFile("./assets/Characters/Skeleton/PNG/skeleton_idle_strip6.png")
	//! Insert enemyImg
	if err != nil {
		// Print err then exit game
		log.Fatal(err)
	}

	game := Game{
		player: &Player{
			Sprite: &Sprite{
				Img: playerImg,
				X:   50.0,
				Y:   50.0,
			},
			Health: 5,
		},
		// additional sprite slices of type Enemy struct
		enemies: []*Enemy{
			{
				&Sprite{
					Img: skeletonImg, //! Change source img
					X:   70.0,
					Y:   70.0,
				},
				false,
			},
			{
				&Sprite{
					Img: skeletonImg, //! Change source img
					X:   90.0,
					Y:   90.0,
				},
				true,
			},
			{
				&Sprite{
					Img: skeletonImg, //! Change source img
					X:   110.0,
					Y:   110.0,
				},
				true,
			},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
