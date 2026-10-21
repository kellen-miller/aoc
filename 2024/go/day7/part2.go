package day7

import (
	"bufio"
	"fmt"
	"strconv"
)

func Part2(sc *bufio.Scanner) (string, error) {
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
		if dp2(eq.Answer, eq.Numbers, eq.Numbers[0], 1) {
			total += eq.Answer
		}
	}

	return strconv.Itoa(total), nil
}

func dp2(target int, numbers []int, total int, idx int) bool {
	if idx >= len(numbers) {
		return total == target
	}

	return dp2(target, numbers, concat(total, numbers[idx]), idx+1) || // concatenation
		dp2(target, numbers, total+numbers[idx], idx+1) || // addition
		dp2(target, numbers, total*numbers[idx], idx+1) // multiplication
}

func concat(a int, b int) int {
	str := strconv.Itoa(a) + strconv.Itoa(b)
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return num
}
