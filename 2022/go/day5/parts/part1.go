package parts

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kellen-miller/aoc/2022/go/shared"
	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	sll "github.com/ugurcsen/gods-generic/lists/singlylinkedlist"
)

const (
	crateDescLen         = 3
	moveInstructionParts = 3
)

type crate struct {
	val   string
	stack int
}

func RearrangeCrates(input string) string {
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
			stacks = moveCratesSingle(stacks, line, re)
		}
	}

	return topOfStacks(stacks)
}

func appendCrateLayer(stacks []*sll.List[string], line string) []*sll.List[string] {
	crates := getCrates(line)
	if len(crates) == 0 {
		return stacks
	}

	maxIndex := crates[len(crates)-1].stack
	stacks = ensureStackCapacity(stacks, maxIndex)

	for _, cr := range crates {
		stacks[cr.stack].Prepend(cr.val)
	}

	return stacks
}

func moveCratesSingle(stacks []*sll.List[string], line string, re *regexp.Regexp) []*sll.List[string] {
	nums := re.FindAllString(line, -1)
	if len(nums) < moveInstructionParts {
		return stacks
	}

	moving := mustAtoi(nums[0])
	from := mustAtoi(nums[1]) - 1
	to := mustAtoi(nums[2]) - 1
	stacks = ensureStackCapacity(stacks, from)
	stacks = ensureStackCapacity(stacks, to)

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
		stacks[to].Append(cr)
	}

	return stacks
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

func ensureStackCapacity(stacks []*sll.List[string], index int) []*sll.List[string] {
	for len(stacks) <= index {
		stacks = append(stacks, sll.New[string]())
	}
	return stacks
}

func mustAtoi(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("invalid integer %q: %v", value, err))
	}
	return num
}
