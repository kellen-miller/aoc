package parts

import (
	"github.com/kellen-miller/aoc/2022/go/shared"
	"github.com/kellen-miller/aoc/languages/go/pkg/io"
)

const (
	messageStartSize = 14
)

func StartOfMessage(input string) []int {
	if input == "" {
		input = shared.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var ms []int //nolint:prealloc // we don't know how many messages there are
	for sc.Scan() {
		ms = append(ms, findUniqueSetOfSize(sc.Text(), messageStartSize))
	}

	return ms
}
