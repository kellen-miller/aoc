package day8

import (
	"testing"

	"github.com/kellen-miller/aoc/go/pkg/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay_Part1(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example1.txt",
			want:  "14",
		},
		"input": {
			input: "input1.txt",
			want:  "42283209483350",
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
