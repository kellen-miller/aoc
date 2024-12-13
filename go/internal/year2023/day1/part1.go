package day1

import (
	"bufio"
	"log"
	"strconv"
	"unicode"
)

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
	var sum int

	for sc.Scan() {
		var (
			line  = sc.Text()
			left  uint8
			right uint8
		)

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				left = line[i]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				right = line[i]
				break
			}
		}

		val, err := strconv.Atoi(string(left) + string(right))
		if err != nil {
			log.Panic(err)
		}

		sum += val
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(sum), nil
}
