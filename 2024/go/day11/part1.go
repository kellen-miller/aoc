package day11

import (
	"bufio"
	"strconv"
	"strings"
)

func Part1(sc *bufio.Scanner) (string, error) {
	var stones []int
	for sc.Scan() {
		sstr := strings.Fields(sc.Text())
		for _, st := range sstr {
			s, err := strconv.Atoi(st)
			if err != nil {
				return "", err
			}

			stones = append(stones, s)
		}
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	for range 25 {
		var err error
		stones, err = blink(stones)
		if err != nil {
			return "", err
		}
	}

	return strconv.Itoa(len(stones)), nil
}

func blink(stones []int) ([]int, error) {
	var newStones []int
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		stoneStr := strconv.Itoa(stone)
		if len(stoneStr)%2 != 0 {
			newStones = append(newStones, stone*2024)
			continue
		}

		half := len(stoneStr) / 2

		left, err := strconv.Atoi(stoneStr[:half])
		if err != nil {
			return nil, err
		}

		right, err := strconv.Atoi(stoneStr[half:])
		if err != nil {
			return nil, err
		}

		newStones = append(newStones, left, right)
	}

	return newStones, nil
}
