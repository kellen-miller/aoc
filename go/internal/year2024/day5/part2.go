package day5

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ugurcsen/gods-generic/queues/arrayqueue"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
	var (
		parseUpdate    bool
		pageUpdates    [][]string
		pageOrderRules = make(map[string][]string)
	)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			parseUpdate = true
			continue
		}

		if parseUpdate {
			pageUpdates = append(pageUpdates, strings.Split(line, ","))
			continue
		}

		x, y, ok := strings.Cut(line, "|")
		if !ok {
			return "", fmt.Errorf("could not parse page order rule: %s", line)
		}

		pageOrderRules[x] = append(pageOrderRules[x], y)
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	var middlePageSum int
	for _, pageUpdate := range pageUpdates {
		if valid := isOrdered(pageUpdate, pageOrderRules); valid {
			continue
		}

		correctedUpdate, err := topologicalSort(pageUpdate, pageOrderRules)
		if err != nil {
			return "", fmt.Errorf("failed to reorder update %v: %w", pageUpdate, err)
		}

		middleIdx := len(correctedUpdate) / 2 //nolint:mnd // Middle index is calculated based on the length of the slice.
		middlePage, err := strconv.Atoi(correctedUpdate[middleIdx])
		if err != nil {
			return "", err
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
