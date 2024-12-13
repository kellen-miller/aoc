package day5

import (
	"testing"

	"github.com/kellen-miller/aoc/go/pkg/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay_Part2(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example2.txt",
			want:  "",
		},
		"input": {
			input: "input2.txt",
			want:  "",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			got, err := new(Day).Part1(sc)
			require.NoError(t, err)

			assert.Equal(t, tc.want, got)
		})
	}
}
