package internal

import "bufio"

type AdventYear interface {
	Year() int
	AdventDays() []AdventDay
}

type AdventDay interface {
	Day() int
	Part1(sc *bufio.Scanner) (string, error)
	Part2(sc *bufio.Scanner) (string, error)
}