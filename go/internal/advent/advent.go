package advent

import (
	"bufio"
)

type Year interface {
	AdventYear() int
	AdventDays() []Day
}

type Day interface {
	AdventDay() int
	Part1(sc *bufio.Scanner) (string, error)
	Part2(sc *bufio.Scanner) (string, error)
}
