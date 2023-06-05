package play

import "juanantoniocid/mariana/internal/geometry"

func (p *Play) MoveStarship(dir geometry.Direction) {
	p.moveStarship(dir)
	p.timer++
}

func (p *Play) moveStarship(dir geometry.Direction) {
	if p.needsToMoveStarship() {
		p.starship.Move(dir)
	}
}

func (p *Play) needsToMoveStarship() bool {
	return p.timer%p.moveTime == 0
}
