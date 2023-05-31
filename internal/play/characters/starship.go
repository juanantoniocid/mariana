package characters

import "juanantoniocid/mariana/internal/geometry"

// Starship represents a starship
type Starship struct {
	body geometry.Shape
}

// Move moves the starship
func (s *Starship) Move(dir geometry.Direction) {
	var incX, incY int

	switch dir {
	case geometry.DirLeft:
		incX = -1
	case geometry.DirRight:
		incX = 1
	case geometry.DirDown:
		incY = 1
	case geometry.DirUp:
		incY = -1
	}

	for i := int64(len(s.body)) - 1; i >= 0; i-- {
		s.body[i].X += incX
		s.body[i].Y += incY
	}
}

// GetShape returns the shape of the starship
func (s *Starship) GetShape() geometry.Shape {
	return s.body
}

// NewStarship creates a new starship
func NewStarship(posX, posY int) *Starship {
	return &Starship{
		body: []geometry.Position{
			{X: posX, Y: posY},
			{X: posX - 1, Y: posY},
			{X: posX + 1, Y: posY},
			{X: posX, Y: posY - 1},
			{X: posX, Y: posY + 1},
		},
	}
}
