package grid

type Grid[T any] struct {
	cells []T
	cols  int
	rows  int
}

func NewGrid[T any](cols int, rows int) *Grid[T] {
	return &Grid[T]{
		cells: make([]T, cols*rows),
		cols:  cols,
		rows:  rows,
	}
}

func (g *Grid[T]) Do(f func(p Point, value T)) {
	for x := 0; x < g.cols; x++ {
		for y := 0; y < g.rows; y++ {
			f(Point{X: x, Y: y}, g.cells[x*g.cols+y])
		}
	}
}

func (g *Grid[T]) Cell(p Point) *T {
	if !g.IsValidPoint(p) {
		return nil
	}

	return &g.cells[p.X*g.cols+p.Y]
}

func (g *Grid[T]) SetCell(p Point, v T) {
	if g.IsValidPoint(p) {
		g.cells[p.X*g.cols+p.Y] = v
	}
}

func (g *Grid[T]) Rows() int {
	return g.rows
}

func (g *Grid[T]) Columns() int {
	return g.cols
}

func (g *Grid[T]) Len() int {
	return g.rows * g.cols
}

func (g *Grid[T]) IsValidPoint(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < g.cols && p.Y < g.rows
}

func (g *Grid[T]) Copy() *Grid[T] {
	c := make([]T, g.cols*g.rows)
	copy(c, g.cells)

	return &Grid[T]{
		cells: c,
		cols:  g.cols,
		rows:  g.rows,
	}
}

func (g *Grid[T]) Up(p Point) T {
	return g.cells[(p.X-1)*g.cols+p.Y]
}

func (g *Grid[T]) Down(p Point) T {
	return g.cells[(p.X+1)*g.cols+p.Y]
}

func (g *Grid[T]) Left(p Point) T {
	return g.cells[p.X*g.cols+p.Y-1]
}

func (g *Grid[T]) Right(p Point) T {
	return g.cells[p.X*g.cols+p.Y+1]
}

func (g *Grid[T]) UpLeft(p Point) T {
	return g.cells[(p.X-1)*g.cols+p.Y-1]
}

func (g *Grid[T]) UpRight(p Point) T {
	return g.cells[(p.X-1)*g.cols+p.Y+1]
}

func (g *Grid[T]) DownLeft(p Point) T {
	return g.cells[(p.X+1)*g.cols+p.Y-1]
}

func (g *Grid[T]) DownRight(p Point) T {
	return g.cells[(p.X+1)*g.cols+p.Y+1]
}
