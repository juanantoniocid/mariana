package play

import (
	"juanantoniocid/mariana/internal/geometry"
	"juanantoniocid/mariana/internal/play/characters"
)

// Play is the main game struct
type Play struct {
	boardWidth  int
	boardHeight int

	starship *characters.Starship

	timer    int
	moveTime int
}

// NewPlay creates a new game
func NewPlay(width, height int) *Play {
	g := &Play{
		boardWidth:  width,
		boardHeight: height,
		moveTime:    1,
	}
	g.Reset()

	return g
}

// Reset resets the game
func (p *Play) Reset() {
	p.initStarship()
}

func (p *Play) initStarship() {
	startingPoint := geometry.Point{
		X: p.boardWidth / 2,
		Y: p.boardHeight / 2,
	}
	endOfUniverse := geometry.Point{
		X: p.boardWidth - 1,
		Y: p.boardHeight - 1,
	}

	p.starship = characters.NewStarship(startingPoint, endOfUniverse)
}

func (p *Play) GetStarshipShape() geometry.Shape {
	return p.starship.GetShape()
}
