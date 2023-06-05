package game

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"juanantoniocid/mariana/internal/geometry"
	"juanantoniocid/mariana/internal/play"
)

const (
	ScreenWidth            = 320
	ScreenHeight           = 200
	gridSize       float32 = 1
	SizeMultiplier         = 3
)

type Game struct {
	play *play.Play
}

func (g *Game) Update() error {
	g.iterate()
	return nil
}

func (g *Game) iterate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.play.Reset()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	dir := getStarshipDirection()
	g.play.MoveStarship(dir)
}

func getStarshipDirection() geometry.Direction {
	dir := geometry.DirNone

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			dir = geometry.DirLeftUp
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			dir = geometry.DirRightUp
		} else {
			dir = geometry.DirUp
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			dir = geometry.DirLeftDown
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			dir = geometry.DirRightDown
		} else {
			dir = geometry.DirDown
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dir = geometry.DirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dir = geometry.DirRight
	}

	return dir
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawStarship(screen)
}

func (g *Game) drawStarship(screen *ebiten.Image) {
	starship := g.play.GetStarshipShape()
	for _, p := range starship {
		vector.DrawFilledRect(
			screen, float32(p.X)*gridSize, float32(p.Y)*gridSize, gridSize, gridSize,
			color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff}, false)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{
		play: play.NewPlay(ScreenWidth/int(gridSize), ScreenHeight/int(gridSize)),
	}

	return g
}
