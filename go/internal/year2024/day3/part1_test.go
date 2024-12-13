package day3

import (
	"testing"

	"github.com/kellen-miller/advent-of-code/go/pkg/io"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example1.txt",
			want:  "161",
		},
		"input": {
			input: "input1.txt",
			want:  "174960292",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			d := new(Day)
			got, err := d.Part1(sc)
			require.NoError(t, err)

			if got != tc.want && tc.want != "" {
				t.Errorf("got %s; want %s", got, tc.want)
			} else {
				t.Logf("got %s", got)
			}
		})
	}
}
