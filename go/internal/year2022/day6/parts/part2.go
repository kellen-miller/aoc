package parts

import (
	"github.com/kellen-miller/aoc/go/internal/year2022"
	"github.com/kellen-miller/aoc/go/pkg/io"
)

const (
	messageStartSize = 14
)

func StartOfMessage(input string) []int {
	if input == "" {
		input = year2022.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var ms []int //nolint:prealloc // we don't know how many messages there are
	for sc.Scan() {
		ms = append(ms, findUniqueSetOfSize(sc.Text(), messageStartSize))
	}

	return ms
}
