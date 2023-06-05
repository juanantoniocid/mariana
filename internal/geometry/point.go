package geometry

type Point struct {
	X int
	Y int
}

type Shape []Point

func GetShapeFromBitmap(bitmap [][]bool, gravityCenter Point, direction Direction) Shape {
	if direction == DirLeft || direction == DirLeftDown || direction == DirLeftUp {
		bitmap = MirrorBitmap(bitmap)
	}

	shape := make(Shape, 0)

	width := len(bitmap[0])
	height := len(bitmap)
	left := gravityCenter.X - width/2
	top := gravityCenter.Y - height/2

	for y, row := range bitmap {
		for x, value := range row {
			if value {
				shape = append(shape, Point{
					X: x + left,
					Y: y + top,
				})
			}
		}
	}

	return shape
}

func MirrorBitmap(bitmap [][]bool) [][]bool {
	width := len(bitmap[0])
	height := len(bitmap)

	mirrored := make([][]bool, height)

	for y, row := range bitmap {
		mirrored[y] = make([]bool, width)

		for x, value := range row {
			mirrored[y][width-x-1] = value
		}
	}

	return mirrored
}
