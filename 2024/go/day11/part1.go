package day11

import (
	"bufio"
	"strconv"
	"strings"
)

const (
	blinkIterations  = 25
	stoneSplitFactor = 2
	stoneMultiplier  = 2024
	defaultStoneCap  = 64
)

func Part1(sc *bufio.Scanner) (string, error) {
	stones := make([]int, 0, defaultStoneCap)
	for sc.Scan() {
		sstr := strings.FieldsSeq(sc.Text())
		for st := range sstr {
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

	for range blinkIterations {
		var err error
		stones, err = blink(stones)
		if err != nil {
			return "", err
		}
	}

	return strconv.Itoa(len(stones)), nil
}

func blink(stones []int) ([]int, error) {
	newStones := make([]int, 0, len(stones)*stoneSplitFactor)
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		stoneStr := strconv.Itoa(stone)
		if len(stoneStr)%stoneSplitFactor != 0 {
			newStones = append(newStones, stone*stoneMultiplier)
			continue
		}

		half := len(stoneStr) / stoneSplitFactor

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
