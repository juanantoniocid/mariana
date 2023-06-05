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
	p.starship = characters.NewStarship(p.boardWidth/2, p.boardHeight/2)
}

func (p *Play) GetStarshipShape() geometry.Shape {
	return p.starship.GetShape()
}

func (p *Play) GetStarshipGravityCenter() geometry.Point {
	return p.starship.GetGravityCenter()
}
