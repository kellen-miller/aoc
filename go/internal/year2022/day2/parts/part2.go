//nolint:mnd // package is outdated and not maintained
package parts

import (
	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/pkg/io"
)

const (
	choices = 3
)

func SetRoundResult(input string) int {
	if input == "" {
		input = internal.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		total    int
		scoreMap = map[uint8]int{
			'A': rock,
			'B': paper,
			'C': scissors,
			'X': loss,
			'Y': draw,
			'Z': win,
		}
	)

	for sc.Scan() {
		var (
			line         = sc.Text()
			oppChoice    = scoreMap[line[0]]
			resultWanted = scoreMap[line[2]]
		)

		total += resultWanted + 1
		switch resultWanted {
		case win:
			total += (oppChoice + 1) % choices
		case draw:
			total += oppChoice
		case loss:
			total += (oppChoice + 2) % choices
		}
	}

	return total
}
