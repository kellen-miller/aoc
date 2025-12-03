package day11

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

const (
	blinkMemoIterations = 75
	decimalBase         = 10
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
		sstr := strings.FieldsSeq(sc.Text())
		for st := range sstr {
			s, err := strconv.Atoi(st)
			if err != nil {
				return "", err
			}

			stones += blinkMemo(s, blinkMemoIterations, memo)
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
		if digits%stoneSplitFactor != 0 {
			result = blinkMemo(stone*stoneMultiplier, iteration-1, memo)
			break
		}

		divisor := int(math.Pow10(digits / stoneSplitFactor))
		result = blinkMemo(stone/divisor, iteration-1, memo) + blinkMemo(stone%divisor, iteration-1, memo)
	}

	memo[key] = result
	return result
}
