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
	ScreenWidth          = 640
	ScreenHeight         = 480
	gridSize     float32 = 5
)

type Game struct {
	play *play.Play
}

func (g *Game) Update() error {
	g.iterate()
	return nil
}

func (g *Game) iterate() {
	dir := geometry.DirNone

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.play.Reset()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dir = geometry.DirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dir = geometry.DirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dir = geometry.DirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dir = geometry.DirRight
	}
	g.play.MoveStarship(dir)
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
