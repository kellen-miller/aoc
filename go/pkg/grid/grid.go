package grid

import (
	"bufio"

	"github.com/ugurcsen/gods-generic/queues/arrayqueue"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

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

func NewGridFromFile[T any](sc *bufio.Scanner, lineParseFn func(line string, row int) []T) (*Grid[T], error) {
	//nolint: prealloc // We don't know the size of the grid yet
	var (
		rows [][]T
		i    int
	)
	for sc.Scan() {
		rows = append(rows, lineParseFn(sc.Text(), i))
		i++
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	cols := len(rows[0])
	grid := NewGrid[T](cols, len(rows))

	for i, row := range rows {
		for j, value := range row {
			grid.SetCell(Point{X: i, Y: j}, value)
		}
	}

	return grid, nil
}

func (g *Grid[T]) Do(f func(p Point, value T)) {
	for x := range g.cols {
		for y := range g.rows {
			f(Point{X: x, Y: y}, g.cells[x*g.cols+y])
		}
	}
}

func (g *Grid[T]) Cell(p Point) T {
	if !g.IsValidPoint(p) {
		return *new(T)
	}

	return g.cells[p.X*g.cols+p.Y]
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

func (g *Grid[T]) Up(p Point) Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func (g *Grid[T]) Down(p Point) Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func (g *Grid[T]) Left(p Point) Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (g *Grid[T]) Right(p Point) Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (g *Grid[T]) UpLeft(p Point) Point {
	return Point{X: p.X - 1, Y: p.Y - 1}
}

func (g *Grid[T]) UpRight(p Point) Point {
	return Point{X: p.X - 1, Y: p.Y + 1}
}

func (g *Grid[T]) DownLeft(p Point) Point {
	return Point{X: p.X + 1, Y: p.Y - 1}
}

func (g *Grid[T]) DownRight(p Point) Point {
	return Point{X: p.X + 1, Y: p.Y + 1}
}

func (g *Grid[T]) IsFirstColumn(p Point) bool {
	return p.Y == 0
}

func (g *Grid[T]) IsLastColumn(p Point) bool {
	return p.Y == g.Columns()-1
}

func (g *Grid[T]) IsFirstRow(p Point) bool {
	return p.X == 0
}

func (g *Grid[T]) IsLastRow(p Point) bool {
	return p.X == g.Rows()-1
}

func (g *Grid[T]) IsEdge(p Point) bool {
	return g.IsFirstRow(p) || g.IsFirstColumn(p) || g.IsLastRow(p) || g.IsLastColumn(p)
}

func (g *Grid[T]) DFS(p Point, f func(p Point, value T) bool) {
	g.dfs(p, hashset.New[Point](), f)
}

func (g *Grid[T]) dfs(p Point, visited *hashset.Set[Point], f func(p Point, value T) bool) {
	if visited.Contains(p) {
		return
	}

	visited.Add(p)
	f(p, g.Cell(p))

	for _, dirFn := range g.DirFns() {
		next := dirFn(p)
		if g.IsValidPoint(next) {
			g.dfs(next, visited, f)
		}
	}
}

func (g *Grid[T]) BFS(pt Point, f func(p Point, value T)) {
	visited := make(map[Point]bool)
	queue := arrayqueue.New[Point]()
	queue.Enqueue(pt)

	for queue.Size() > 0 {
		p, ok := queue.Dequeue()
		if !ok {
			break
		}

		if visited[p] {
			continue
		}

		visited[p] = true
		f(p, g.Cell(p))

		for _, dirFn := range g.DirFns() {
			next := dirFn(p)
			if g.IsValidPoint(next) {
				queue.Enqueue(next)
			}
		}
	}
}

func (g *Grid[T]) DirFns() []func(Point) Point {
	return []func(Point) Point{
		g.Up,
		g.Down,
		g.Left,
		g.Right,
		g.UpLeft,
		g.UpRight,
		g.DownLeft,
		g.DownRight,
	}
}
