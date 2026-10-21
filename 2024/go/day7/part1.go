package day7

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	Numbers []int
	Answer  int
}

func Part1(sc *bufio.Scanner) (string, error) {
	var equations []*Equation
	for sc.Scan() {
		line := sc.Text()
		eq, err := parseLine(line)
		if err != nil {
			return "", fmt.Errorf("parse equation %q: %w", line, err)
		}

		equations = append(equations, eq)
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("scan equations: %w", err)
	}

	var total int
	for _, eq := range equations {
		if dp1(eq.Answer, eq.Numbers, eq.Numbers[0], 1) {
			total += eq.Answer
		}
	}

	return strconv.Itoa(total), nil
}

func parseLine(line string) (*Equation, error) {
	a, ns, ok := strings.Cut(line, ": ")
	if !ok {
		return nil, errors.New("invalid input line: " + line)
	}

	numbersStr := strings.Split(ns, " ")

	numbers := make([]int, len(numbersStr))
	for i, n := range numbersStr {
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("parse equation number %q: %w", n, err)
		}

		numbers[i] = number
	}

	answer, err := strconv.Atoi(a)
	if err != nil {
		return nil, fmt.Errorf("parse equation answer %q: %w", a, err)
	}

	return &Equation{
		Answer:  answer,
		Numbers: numbers,
	}, nil
}

func dp1(target int, numbers []int, total int, idx int) bool {
	if idx >= len(numbers) {
		return total == target
	}

	return dp1(target, numbers, total+numbers[idx], idx+1) ||
		dp1(target, numbers, total*numbers[idx], idx+1)
}
