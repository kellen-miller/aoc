package day5

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"

	"github.com/ugurcsen/gods-generic/queues/arrayqueue"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

func Part2(sc *bufio.Scanner) (string, error) {
	manual, err := readPageManual(sc)
	if err != nil {
		return "", err
	}

	var middlePageSum int
	for _, pageUpdate := range manual.updates {
		if valid := isOrdered(pageUpdate, manual.rules); valid {
			continue
		}

		correctedUpdate, err := topologicalSort(pageUpdate, manual.rules)
		if err != nil {
			return "", fmt.Errorf("failed to reorder update %v: %w", pageUpdate, err)
		}

		middleIdx := len(
			correctedUpdate,
		) / 2 //nolint:mnd // Middle index is calculated based on the length of the slice.
		middlePage, err := strconv.Atoi(correctedUpdate[middleIdx])
		if err != nil {
			return "", fmt.Errorf("parse middle page %q: %w", correctedUpdate[middleIdx], err)
		}

		middlePageSum += middlePage
	}

	return strconv.Itoa(middlePageSum), nil
}

// topologicalSort performs a topological sort on the given pages based on the ruleMap (Kahn's algorithm).
func topologicalSort(pageUpdate []string, pageOrderRules map[string][]string) ([]string, error) {
	var (
		pageSet  = hashset.New[string]()
		inDegree = make(map[string]int)
	)
	for _, page := range pageUpdate {
		pageSet.Add(page)
		inDegree[page] = 0
	}

	adj := make(map[string][]string)
	for _, page := range pageUpdate {
		for _, after := range pageOrderRules[page] {
			if pageSet.Contains(after) {
				adj[page] = append(adj[page], after)
				inDegree[after]++
			}
		}
	}

	queue := arrayqueue.New[string]()
	for _, page := range pageUpdate {
		if inDegree[page] == 0 {
			queue.Enqueue(page)
		}
	}

	sorted := make([]string, 0, queue.Size())
	for !queue.Empty() {
		current, ok := queue.Dequeue()
		if !ok {
			return nil, errors.New("queue is empty")
		}

		sorted = append(sorted, current)
		for _, neighbor := range adj[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue.Enqueue(neighbor)
			}
		}
	}

	if len(sorted) != len(pageUpdate) {
		return nil, errors.New("cycle detected or missing rules")
	}

	return sorted, nil
}
