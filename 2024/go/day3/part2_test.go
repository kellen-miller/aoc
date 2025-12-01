package day3

import (
	"bufio"
	"testing"

	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  string
	}{
		"example": {
			input: "example2.txt",
			want:  "48",
		},
		"input": {
			input: "input2.txt",
			want:  "56275602",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			sc, closeFile := io.GetScanner(tc.input)
			defer closeFile()

			got, err := Part2(sc)
			require.NoError(t, err)

			if got != tc.want && tc.want != "" {
				t.Errorf("got %s; want %s", got, tc.want)
			} else {
				t.Logf("got %s", got)
			}
		})
	}
}

func BenchmarkDay_Part2(b *testing.B) {

	b.Run("input1.txt", func(b *testing.B) {
		f, closeFile := io.OpenFile("input1.txt")
		defer closeFile()

		b.Run("Part2 - all regex", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Part2(bufio.NewScanner(f))
			}
		})

		b.Run("Part2 - split/cut to regex", func(b *testing.B) {
			for range b.N {
				_, _ = part2SplitCutRegex(bufio.NewScanner(f))
			}
		})
	})

	b.Run("input2.txt", func(b *testing.B) {
		f, closeFile := io.OpenFile("input2.txt")
		defer closeFile()

		b.Run("Part2 - all regex", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Part2(bufio.NewScanner(f))
			}
		})

		b.Run("Part2 - split/cut to regex", func(b *testing.B) {
			for range b.N {
				_, _ = part2SplitCutRegex(bufio.NewScanner(f))
			}
		})
	})
}
