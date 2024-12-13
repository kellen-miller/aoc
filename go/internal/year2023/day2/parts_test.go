package day2

import "testing"

func TestPart1(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "example.txt",
			want:  8,
		},
		"input": {
			input: "input.txt",
			want:  2593,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got := GamesPossibleSum(tc.input)
			if got != tc.want && tc.want != 0 {
				t.Errorf("GamesPossibleSum = %d; want %d", got, tc.want)
			} else {
				t.Logf("GamesPossibleSum = %d", got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "example.txt",
			want:  2286,
		},
		"input": {
			input: "input.txt",
			want:  54699,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got := GamesPossiblePowerSum(tc.input)
			if got != tc.want && tc.want != 0 {
				t.Errorf("GamesPossiblePowerSum = %d; want %d", got, tc.want)
			} else {
				t.Logf("GamesPossiblePowerSum = %d", got)
			}
		})
	}
}
