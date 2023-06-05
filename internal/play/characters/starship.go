package characters

import (
	"juanantoniocid/mariana/internal/geometry"
)

// Starship represents a starship
type Starship struct {
	gravityCenter geometry.Point
	endUniverse   geometry.Point
	speed         int
	orientation   geometry.Direction

	body    geometry.Shape
	marginX int
	marginY int
}

// NewStarship creates a new starship
func NewStarship(pos, endUniverse geometry.Point) *Starship {
	starship := &Starship{
		gravityCenter: pos,
		endUniverse:   endUniverse,
		speed:         1,
		orientation:   geometry.DirRight,
	}
	starship.initializeShape()

	return starship
}

func (s *Starship) initializeShape() {
	bitmap := getStarshipBitmap()
	s.marginX = len(bitmap[0]) / 2
	s.marginY = len(bitmap) / 2

	s.setShape()
}

// GetShape returns the shape of the starship
func (s *Starship) GetShape() geometry.Shape {
	return s.body
}

// Move moves the starship
func (s *Starship) Move(dir geometry.Direction) {
	s.setOrientation(dir)
	s.shift(dir)
	s.setShape()
}

func (s *Starship) setOrientation(dir geometry.Direction) {
	if dir == geometry.DirLeft || dir == geometry.DirLeftDown || dir == geometry.DirLeftUp {
		s.orientation = geometry.DirLeft
	} else if dir == geometry.DirRight || dir == geometry.DirRightDown || dir == geometry.DirRightUp {
		s.orientation = geometry.DirRight
	}
}

func (s *Starship) shift(dir geometry.Direction) {
	incX, incY := s.getShift(dir)
	s.gravityCenter.X += incX
	s.gravityCenter.Y += incY
	s.keepWithinUniverse()
}

func (s *Starship) getShift(dir geometry.Direction) (int, int) {
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

func (s *Starship) keepWithinUniverse() {
	if s.gravityCenter.X < s.marginX {
		s.gravityCenter.X = s.marginX
	} else if s.gravityCenter.X > s.endUniverse.X-s.marginX {
		s.gravityCenter.X = s.endUniverse.X - s.marginX
	}

	if s.gravityCenter.Y < s.marginY {
		s.gravityCenter.Y = s.marginY
	} else if s.gravityCenter.Y > s.endUniverse.Y-s.marginY {
		s.gravityCenter.Y = s.endUniverse.Y - s.marginY
	}
}

func (s *Starship) setShape() {
	bitmap := getStarshipBitmap()
	s.body = geometry.GetShapeFromBitmap(bitmap, s.gravityCenter, s.orientation)
}

func getStarshipBitmap() [][]bool {
	const o = false
	const x = true

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
