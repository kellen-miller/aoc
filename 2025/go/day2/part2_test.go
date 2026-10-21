package day2

import (
	"strconv"
	"testing"

	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type part2TestCase struct {
	input string
	want  string
}

func TestPart2(t *testing.T) {
	tcs := map[string]part2TestCase{
		"example": {
			input: "example2.txt",
			want:  "4174379265",
		},
		"input": {
			input: "input2.txt",
			want:  "66500947346",
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

func TestGenerateInvalidValuesMatchesBruteForce(t *testing.T) {
	const maxValue = 200_000

	got, _ := generateInvalidValues(digitCount(maxValue), maxValue)
	want := make([]int64, 0)
	for value := int64(1); value <= maxValue; value++ {
		text := strconv.FormatInt(value, decimalBase)
		for blockDigits := 1; blockDigits*minRepeatCount <= len(text); blockDigits++ {
			if len(text)%blockDigits != 0 {
				continue
			}

			repeated := true
			for digit := blockDigits; digit < len(text); digit++ {
				if text[digit] != text[digit%blockDigits] {
					repeated = false
					break
				}
			}

			if repeated {
				want = append(want, value)
				break
			}
		}
	}

	require.Equal(t, want, got)
}
