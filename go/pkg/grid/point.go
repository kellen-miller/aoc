package grid

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) Delta() (int, int) {
	switch d {
	case Up:
		return 0, -1
	case Right:
		return 1, 0
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	default:
		return 0, 0
	}
}

type Point struct {
	X int
	Y int
}

func (p Point) MoveDirection(d Direction) Point {
	switch d {
	case Up:
		return p.Up()
	case Down:
		return p.Down()
	case Left:
		return p.Left()
	case Right:
		return p.Right()
	}

	return p
}

func (p Point) Up() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (p Point) Down() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (p Point) Left() Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func (p Point) Right() Point {
	return Point{X: p.X + 1, Y: p.Y}
}
