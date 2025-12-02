package day2

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
			input: "example2.txt",
			want:  "4174379265",
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

			got, err := Part2(sc)

			if name == "input" && tc.want == "" {
				t.Logf("got %s", got)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
