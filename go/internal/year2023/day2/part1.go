package day2

import (
	"bufio"
	"strconv"
	"strings"
)

const (
	MaxRedCubes   = 12
	MaxGreenCubes = 13
	MaxBlueCubes  = 14
)

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
	var sum int
	for sc.Scan() {
		sum += parseGame(sc.Text())
	}

	return strconv.Itoa(sum), nil
}

func parseGame(line string) int {
	const gamePartsWant = 2

	gameParts := strings.Split(line, ":")
	if len(gameParts) != gamePartsWant {
		panic("invalid game")
	}

	if !parseRounds(gameParts[1]) {
		return 0
	}

	return parseID(gameParts[0])
}

func parseID(gamePart string) int {
	const idPartsWant = 2

	idParts := strings.Split(gamePart, " ")
	if len(idParts) != idPartsWant {
		panic("invalid game id")
	}

	id, err := strconv.Atoi(idParts[1])
	if err != nil {
		panic("invalid game id")
	}

	return id
}

func parseRounds(roundPart string) bool {
	for _, round := range strings.Split(roundPart, ";") {
		for _, cube := range strings.Split(round, ",") {
			if !isCubeValid(cube) {
				return false
			}
		}
	}

	return true
}

func isCubeValid(cube string) bool {
	const cubePartsWant = 2

	cubeParts := strings.Split(strings.TrimSpace(cube), " ")
	if len(cubeParts) != cubePartsWant {
		panic("invalid cube part")
	}

	cubes, err := strconv.Atoi(cubeParts[0])
	if err != nil {
		panic("invalid cube amount")
	}

	switch cubeParts[1] {
	case "red":
		if cubes > MaxRedCubes {
			return false
		}
	case "green":
		if cubes > MaxGreenCubes {
			return false
		}
	case "blue":
		if cubes > MaxBlueCubes {
			return false
		}
	default:
		panic("unknown cube color")
	}

	return true
}
