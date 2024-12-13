package day3

import "testing"

func TestPart1(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "example.txt",
			want:  4361,
		},
		"input": {
			input: "input.txt",
			want:  531561,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got := FindValidParts(tc.input)
			if got != tc.want && tc.want != 0 {
				t.Errorf("FindValidParts = %d; want %d", got, tc.want)
			} else {
				t.Logf("FindValidParts = %d", got)
			}
		})
	}
}
