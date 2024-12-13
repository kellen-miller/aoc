//nolint:intrange,errcheck // package is outdated and not maintained
package parts

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/pkg/io"
	sll "github.com/ugurcsen/gods-generic/lists/singlylinkedlist"
	lls "github.com/ugurcsen/gods-generic/stacks/linkedliststack"
)

func RearrangeCratesMulti(input string) string {
	if input == "" {
		input = internal.Input
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

			moveStack := lls.New[string]()
			for i := 0; i < moving; i++ {
				top := stacks[from].Size() - 1

				cr, ok := stacks[from].Get(top)
				if !ok {
					break
				}

				stacks[from].Remove(top)
				moveStack.Push(cr)
			}

			for len(stacks) <= to {
				stacks = append(stacks, sll.New[string]())
			}

			for i := 0; i < moveStack.Size(); i++ {
				cr, _ := moveStack.Pop()
				stacks[to].Append(cr)
			}
		}
	}

	return topOfStacks(stacks)
}
