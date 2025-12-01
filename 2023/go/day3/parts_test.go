package day3

import (
	"testing"

	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example.txt",
			want:  "4361",
		},
		"input": {
			input: "input.txt",
			want:  "531561",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			got, err := Part2(sc)
			require.NoError(t, err)

			if got != tc.want && tc.want != "" {
				t.Errorf("FindValidParts = %s; want %s", got, tc.want)
			} else {
				t.Logf("FindValidParts = %s", got)
			}
		})
	}
}
