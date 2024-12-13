package day1

import (
	"bufio"
	"errors"
	"strconv"
)

type location struct {
	id    int
	idStr string
}

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
	var (
		leftLocations []location
		frequencies   = make(map[string]int)
	)
	for sc.Scan() {
		line := sc.Text()

		matches := locationIDsInputRegex.FindStringSubmatch(line)
		if matches == nil || len(matches) != 3 {
			return "", errors.New("invalid input: expected two integers separated by whitespace")
		}

		leftLocationID, err := strconv.Atoi(matches[1])
		if err != nil {
			return "", err
		}

		leftLocations = append(leftLocations, location{
			id:    leftLocationID,
			idStr: matches[1],
		})
		frequencies[matches[2]]++
	}

	var similarityScore int
	for _, loc := range leftLocations {
		similarityScore += loc.id * frequencies[loc.idStr]
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(similarityScore), nil
}
