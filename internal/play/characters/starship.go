package characters

import "juanantoniocid/mariana/internal/geometry"

// Starship represents a starship
type Starship struct {
	body  geometry.Shape
	speed int
}

// Move moves the starship
func (s *Starship) Move(dir geometry.Direction) {
	incX, incY := s.getJump(dir)

	for i := int64(len(s.body)) - 1; i >= 0; i-- {
		s.body[i].X += incX
		s.body[i].Y += incY
	}
}

func (s *Starship) getJump(dir geometry.Direction) (int, int) {
	var incX, incY int

	if dir == geometry.DirLeft || dir == geometry.DirLeftDown || dir == geometry.DirLeftUp {
		incX = -s.speed
	} else if dir == geometry.DirRight || dir == geometry.DirRightDown || dir == geometry.DirRightUp {
		incX = s.speed
	}

	if dir == geometry.DirDown || dir == geometry.DirLeftDown || dir == geometry.DirRightDown {
		incY = s.speed
	} else if dir == geometry.DirUp || dir == geometry.DirLeftUp || dir == geometry.DirRightUp {
		incY = -s.speed
	}

	return incX, incY
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
		speed: 1,
	}
}
