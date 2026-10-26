package day2

import (
	"bufio"
	"testing"

	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	"github.com/stretchr/testify/require"
)

type gameExample struct {
	name  string
	input string
	want  string
}

func TestPart1(t *testing.T) {
	checkGames(t, Part1, []gameExample{
		{name: "example", input: "example.txt", want: "8"},
		{name: "input", input: "input.txt", want: "2593"},
	})
}

func TestPart2(t *testing.T) {
	checkGames(t, Part2, []gameExample{
		{name: "example", input: "example.txt", want: "2286"},
		{name: "input", input: "input.txt", want: "54699"},
	})
}

func checkGames(t *testing.T, solve func(*bufio.Scanner) (string, error), examples []gameExample) {
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
