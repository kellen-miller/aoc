//nolint:intrange,errcheck // package is outdated and not maintained
package parts

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kellen-miller/aoc/go/internal/year2022"
	"github.com/kellen-miller/aoc/go/pkg/io"
	sll "github.com/ugurcsen/gods-generic/lists/singlylinkedlist"
)

const (
	crateDescLen = 3
)

type crate struct {
	val   string
	stack int
}

func RearrangeCrates(input string) string {
	if input == "" {
		input = year2022.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var (
		stacks []*sll.List[string]
		re     = regexp.MustCompile(`\d+`)
	)

	for sc.Scan() {
		line := sc.Text()
		if strings.ContainsRune(line, '[') {
			crates := getCrates(line)

			for len(stacks) <= crates[len(crates)-1].stack {
				stacks = append(stacks, sll.New[string]())
			}

			for _, crate := range crates {
				stacks[crate.stack].Prepend(crate.val)
			}
		} else if strings.HasPrefix(line, "move") {
			var (
				nums = re.FindAllString(line, -1)

				moving, _ = strconv.Atoi(nums[0])
				from, _   = strconv.Atoi(nums[1])
				to, _     = strconv.Atoi(nums[2])
			)

			// decrement for indexing
			from--
			to--

			for i := 0; i < moving; i++ {
				top := stacks[from].Size() - 1

				cr, ok := stacks[from].Get(top)
				if !ok {
					break
				}

				for len(stacks) <= to {
					stacks = append(stacks, sll.New[string]())
				}

				stacks[from].Remove(top)
				stacks[to].Append(cr)
			}
		}
	}

	return topOfStacks(stacks)
}

func getCrates(line string) []crate {
	var crates []crate
	for pos, char := range line {
		if char == '[' {
			crates = append(crates, crate{
				val:   line[pos : pos+crateDescLen],
				stack: (pos + crateDescLen) / (crateDescLen + 1),
			})
		}
	}

	return crates
}

func topOfStacks(stacks []*sll.List[string]) string {
	var sb strings.Builder
	for _, stack := range stacks {
		cr, ok := stack.Get(stack.Size() - 1)
		if ok {
			sb.WriteRune(rune(cr[1]))
		}
	}

	return sb.String()
}
