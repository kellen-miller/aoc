package day4

import (
	"testing"

	"github.com/kellen-miller/advent-of-code/go/pkg/io"
	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example2.txt",
			want:  "9",
		},
		"input": {
			input: "input2.txt",
			want:  "1985",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			d := new(Day)
			got, err := d.Part2(sc)
			require.NoError(t, err)

			if got != tc.want && tc.want != "" {
				t.Errorf("got %s; want %s", got, tc.want)
			} else {
				t.Logf("got %s", got)
			}
		})
	}
}
