package day2

import (
	"testing"

	"github.com/kellen-miller/aoc/go/pkg/io"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example.txt",
			want:  "8",
		},
		"input": {
			input: "input.txt",
			want:  "2593",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			got, err := new(Day).Part1(sc)
			require.NoError(t, err)

			if got != tc.want && tc.want != "" {
				t.Errorf("GamesPossibleSum = %s; want %s", got, tc.want)
			} else {
				t.Logf("GamesPossibleSum = %s", got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example.txt",
			want:  "2286",
		},
		"input": {
			input: "input.txt",
			want:  "54699",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			got, err := new(Day).Part2(sc)
			require.NoError(t, err)

			if got != tc.want && tc.want != "" {
				t.Errorf("GamesPossiblePowerSum = %s; want %s", got, tc.want)
			} else {
				t.Logf("GamesPossiblePowerSum = %s", got)
			}
		})
	}
}
