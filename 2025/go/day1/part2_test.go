package day1

import (
	"testing"

	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example1.txt",
			want:  "6",
		},
		"input": {
			input: "input1.txt",
			want:  "6819",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			got, err := Part2(sc)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
