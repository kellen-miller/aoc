package day5

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPageManual(t *testing.T) {
	t.Run("rules and updates", func(t *testing.T) {
		scanner := bufio.NewScanner(strings.NewReader("1|2\n1|3\n\n1,2,3\n\n3,2,1\n"))
		manual, err := readPageManual(scanner)
		require.NoError(t, err)
		require.Equal(t, map[string][]string{"1": {"2", "3"}}, manual.rules)
		require.Equal(t, [][]string{{"1", "2", "3"}, {"3", "2", "1"}}, manual.updates)
	})

	t.Run("invalid rule", func(t *testing.T) {
		_, err := readPageManual(bufio.NewScanner(strings.NewReader("1,2\n")))
		require.EqualError(t, err, "could not parse page order rule: 1,2")
	})

	t.Run("scanner error", func(t *testing.T) {
		scanner := bufio.NewScanner(strings.NewReader(strings.Repeat("x", bufio.MaxScanTokenSize)))
		_, err := readPageManual(scanner)
		require.ErrorIs(t, err, bufio.ErrTooLong)
	})
}
