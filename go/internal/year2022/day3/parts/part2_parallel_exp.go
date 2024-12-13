package parts

import (
	"sync"

	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/pkg/io"
)

func BadgePriorityTotalParallelGroups(input string) int {
	if input == "" {
		input = internal.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		total int
		wg    sync.WaitGroup
	)

	for {
		rucksacks := make([]string, 0, groupSize)
		for i := 0; i < groupSize; i++ {
			if !sc.Scan() {
				break
			}

			rucksacks = append(rucksacks, sc.Text())
		}

		if len(rucksacks) == 0 {
			break
		}

		wg.Add(1)
		go func(group []string) {
			total += findGroupBadgePriority(group)
			wg.Done()
		}(rucksacks)
	}

	wg.Wait()
	return total
}

func BadgePriorityTotalChannels(input string) int {
	if input == "" {
		input = internal.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		total   int
		wg      sync.WaitGroup
		groups  = make(chan []string)
		results = make(chan int)
	)

	wg.Add(1)
	go func() {
		for {
			group := make([]string, 0, groupSize)
			for i := 0; i < groupSize; i++ {
				if !sc.Scan() {
					break
				}

				group = append(group, sc.Text())
			}

			if len(group) == 0 {
				break
			}

			groups <- group
		}

		close(groups)
		wg.Done()
	}()

	for group := range groups {
		wg.Add(1)
		go func(group []string) {
			results <- findGroupBadgePriority(group)
			wg.Done()
		}(group)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		total += result
	}

	return total
}
