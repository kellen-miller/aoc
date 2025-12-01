package day11

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

type Key struct {
	Stone     int
	Iteration int
}

func Part2(sc *bufio.Scanner) (string, error) {
	var (
		stones int
		memo   = make(map[Key]int)
	)
	for sc.Scan() {
		sstr := strings.Fields(sc.Text())
		for _, st := range sstr {
			s, err := strconv.Atoi(st)
			if err != nil {
				return "", err
			}

			stones += blinkMemo(s, 75, memo)
		}
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(stones), nil
}

func blinkMemo(stone int, iteration int, memo map[Key]int) int {
	key := Key{
		Stone:     stone,
		Iteration: iteration,
	}

	if val, ok := memo[key]; ok {
		return val
	}

	var result int
	switch {
	case iteration == 0:
		result = 1
	case stone == 0:
		result = blinkMemo(1, iteration-1, memo)
	default:
		digits := int(math.Floor(math.Log10(float64(stone)))) + 1
		if digits%2 != 0 {
			result = blinkMemo(stone*2024, iteration-1, memo)
			break
		}

		divisor := int(math.Pow(10, float64(digits/2)))
		result = blinkMemo(stone/divisor, iteration-1, memo) + blinkMemo(stone%divisor, iteration-1, memo)
	}

	memo[key] = result
	return result
}
