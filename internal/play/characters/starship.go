package characters

import "juanantoniocid/mariana/internal/geometry"

// Starship represents a starship
type Starship struct {
	gravityCenter geometry.Point
	body          geometry.Shape
	speed         int
}

// Move moves the starship
func (s *Starship) Move(dir geometry.Direction) {
	incX, incY := s.getJump(dir)

	s.gravityCenter.X += incX
	s.gravityCenter.Y += incY

	bitmap := getBitmap()
	s.body = geometry.GetShapeFromBitmap(bitmap, s.gravityCenter, dir)
}

func getBitmap() [][]bool {
	o := false
	x := true

	return [][]bool{
		{x, x, x, x, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o},
		{x, x, x, x, x, x, x, x, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o},
		{x, x, x, x, x, x, x, x, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o},
		{x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, o, o, o, o, o, o, o},
		{x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x},
		{x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x},
		{x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, o, o, o, o, o, o, o, o, o, o, o},
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

// GetGravityCenter returns the gravity center of the starship
func (s *Starship) GetGravityCenter() geometry.Point {
	return s.gravityCenter
}

// NewStarship creates a new starship
func NewStarship(posX, posY int) *Starship {
	return &Starship{
		body: []geometry.Point{
			{X: posX, Y: posY},
			{X: posX - 1, Y: posY},
			{X: posX + 1, Y: posY},
			{X: posX, Y: posY - 1},
			{X: posX, Y: posY + 1},
		},
		speed: 1,
	}
}
