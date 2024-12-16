package day6

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/go/pkg/grid"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

const (
	guard    = '^'
	obstacle = '#'
	dirs     = 4
)

type GuardState struct {
	Position  grid.Point
	Direction grid.Direction
}

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
	var (
		lab           = make(map[grid.Point]rune)
		guardPosition grid.Point
		row           int
	)
	for sc.Scan() {
		for x, char := range sc.Text() {
			pos := grid.Point{X: x, Y: row}
			lab[pos] = char
			if char == guard {
				guardPosition = pos
			}
		}

		row++
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(predictRoute(lab, GuardState{
		Position:  guardPosition,
		Direction: grid.Up,
	}).Size()), nil
}

func predictRoute(lab map[grid.Point]rune, state GuardState) *hashset.Set[grid.Point] {
	var (
		position        = state.Position
		direction       = state.Direction           // initial direction is always up
		pointsVisited   = hashset.New[grid.Point]() // unique positions visited by the guard
		guardStatesSeen = hashset.New[GuardState]() // unique states of the guard (position & direction)
	)

	pointsVisited.Add(position)
	guardStatesSeen.Add(GuardState{
		Position:  position,
		Direction: direction,
	})

	for {
		nextPos := position.MoveDirection(direction)
		if _, ok := lab[nextPos]; !ok { // guard has moved out of the lab
			break
		}

		if lab[nextPos] == obstacle {
			direction = (direction + 1) % dirs // turn right
			continue
		}

		position = nextPos          // move guard
		pointsVisited.Add(position) // mark position as visited

		currState := GuardState{
			Position:  position,
			Direction: direction,
		}

		if guardStatesSeen.Contains(currState) { // we've seen this state before, so we're in a loop
			break
		}

		guardStatesSeen.Add(currState)
	}

	return pointsVisited
}
