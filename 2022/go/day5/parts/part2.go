package parts

import (
	"regexp"
	"strings"

	"github.com/kellen-miller/aoc/2022/go/shared"
	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	sll "github.com/ugurcsen/gods-generic/lists/singlylinkedlist"
	lls "github.com/ugurcsen/gods-generic/stacks/linkedliststack"
)

func RearrangeCratesMulti(input string) string {
	if input == "" {
		input = shared.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var (
		stacks []*sll.List[string]
		re     = regexp.MustCompile(`\d+`)
	)

	for sc.Scan() {
		line := sc.Text()
		switch {
		case strings.ContainsRune(line, '['):
			stacks = appendCrateLayer(stacks, line)
		case strings.HasPrefix(line, "move"):
			stacks = moveCratesBatch(stacks, line, re)
		}
	}

	return topOfStacks(stacks)
}

func moveCratesBatch(stacks []*sll.List[string], line string, re *regexp.Regexp) []*sll.List[string] {
	nums := re.FindAllString(line, -1)
	if len(nums) < moveInstructionParts {
		return stacks
	}

	moving := mustAtoi(nums[0])
	from := mustAtoi(nums[1]) - 1
	to := mustAtoi(nums[2]) - 1
	stacks = ensureStackCapacity(stacks, from)
	stacks = ensureStackCapacity(stacks, to)

	moveStack := lls.New[string]()
	for range moving {
		top := stacks[from].Size() - 1
		if top < 0 {
			break
		}

		cr, ok := stacks[from].Get(top)
		if !ok {
			break
		}

		stacks[from].Remove(top)
		moveStack.Push(cr)
	}

	for moveStack.Size() > 0 {
		cr, ok := moveStack.Pop()
		if !ok {
			break
		}
		stacks[to].Append(cr)
	}

	return stacks
}
