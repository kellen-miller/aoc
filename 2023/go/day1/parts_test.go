package day1

import (
	"bufio"
	"testing"

	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	"github.com/stretchr/testify/require"
)

type calibrationExample struct {
	name  string
	input string
	want  string
}

func TestPart1(t *testing.T) {
	checkCalibration(t, Part1, []calibrationExample{
		{name: "example", input: "example1.txt", want: "142"},
		{name: "input", input: "input.txt", want: "54630"},
	})
}

func TestPart2(t *testing.T) {
	checkCalibration(t, Part2, []calibrationExample{
		{name: "example", input: "example2.txt", want: "281"},
		{name: "input", input: "input.txt", want: "54770"},
	})
}

func checkCalibration(t *testing.T, solve func(*bufio.Scanner) (string, error), examples []calibrationExample) {
	t.Helper()

	for _, example := range examples {
		t.Run(example.name, func(t *testing.T) {
			scanner, closeFile := io.GetScanner(example.input)
			defer closeFile()

			got, err := solve(scanner)
			require.NoError(t, err)
			require.Equal(t, example.want, got)
		})
	}
}
