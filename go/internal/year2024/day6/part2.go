package day6

import (
	"bufio"
	"strconv"
	"sync"

	"github.com/kellen-miller/aoc/go/pkg/grid"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

const maxGoroutines = 1000

type JumpValue struct {
	NextPos      grid.Point
	NewDir       grid.Direction
	LoopDetected bool
}

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
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

	width, height := 0, row
	for pos := range lab {
		if pos.X+1 > width {
			width = pos.X + 1
		}
		if pos.Y+1 > height {
			height = pos.Y + 1
		}
	}

	return strconv.Itoa(obstaclePositionCausesLoop(lab, GuardState{
		Position:  guardPosition,
		Direction: grid.Up,
	})), nil
}

func obstaclePositionCausesLoop(lab map[grid.Point]rune, state GuardState) int {
	var (
		originalRoutePoints = predictRoute(lab, state)
		obstacleCandidates  = make([]grid.Point, 0, originalRoutePoints.Size())
	)
	for pos := range lab {
		if originalRoutePoints.Contains(pos) && pos != state.Position {
			obstacleCandidates = append(obstacleCandidates, pos)
		}
	}

	var (
		semaphore = make(chan struct{}, maxGoroutines)
		results   = make(chan struct{}, len(obstacleCandidates))
		wg        sync.WaitGroup
	)

	wg.Add(len(obstacleCandidates))
	for _, obstaclePos := range obstacleCandidates {
		semaphore <- struct{}{}
		go func(obstaclePos grid.Point) {
			defer wg.Done()
			defer func() { <-semaphore }()

			labCopy := make(map[grid.Point]rune, len(lab))
			for pos, char := range lab {
				labCopy[pos] = char
			}

			labCopy[obstaclePos] = obstacle

			if simulateRoute(labCopy, state) {
				results <- struct{}{}
			}
		}(obstaclePos)
	}

	wg.Wait()
	close(results)

	return len(results)
}

func simulateRoute(lab map[grid.Point]rune, guardState GuardState) bool {
	var (
		position        = guardState.Position       // initial position of the guard
		direction       = guardState.Direction      // initial direction is always up
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
			return false
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
			return true
		}

		guardStatesSeen.Add(currState)
	}
}
