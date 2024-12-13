package parts

import (
	"github.com/kellen-miller/advent-of-code/go/internal"
	"github.com/kellen-miller/advent-of-code/go/pkg/io"
)

const (
	win      = 6
	draw     = 3
	loss     = 0
	rock     = 0
	paper    = 1
	scissors = 2
)

func TotalScore(input string) int {
	if input == "" {
		input = internal.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		total    int
		scoreMap = map[uint8]int{
			'A': rock,
			'X': rock,
			'B': paper,
			'Y': paper,
			'C': scissors,
			'Z': scissors,
		}
	)

	for sc.Scan() {
		var (
			line      = sc.Text()
			oppChoice = scoreMap[line[0]]
			myChoice  = scoreMap[line[2]]
		)

		total += myChoice + 1
		if oppChoice == myChoice {
			total += draw
		} else if oppChoice+1 == myChoice || oppChoice-2 == myChoice {
			total += win
		}
	}

	return total
}
