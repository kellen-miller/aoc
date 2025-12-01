package parts

import (
	"github.com/kellen-miller/aoc/2022/go/shared"
	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	hs "github.com/ugurcsen/gods-generic/sets/hashset"
)

const (
	alphabetLen     = 26
	packetStartSize = 4
)

func StartOfPacket(input string) []int {
	if input == "" {
		input = shared.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var ps []int //nolint:prealloc // we don't know how many packets there are
	for sc.Scan() {
		ps = append(ps, findUniqueSetOfSize(sc.Text(), packetStartSize))
	}

	return ps
}

func findUniqueSetOfSize(line string, size int) int {
	var (
		chars  = hs.New[uint8]()
		freq   = make([]int, alphabetLen)
		winBeg = 0
		winEnd = 0
	)

	for winEnd < len(line) && winEnd-winBeg < size {
		var (
			endChar    = line[winEnd]
			endCharIdx = endChar - 'a'
		)

		if freq[endCharIdx] != 0 {
			for line[winBeg] != endChar && winBeg < winEnd {
				chars.Remove(line[winBeg])
				freq[line[winBeg]-'a']--
				winBeg++
			}

			winBeg++
		} else {
			chars.Add(endChar)
			freq[endCharIdx]++
		}

		winEnd++
	}

	return winEnd
}
