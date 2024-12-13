package day2

import (
	"strconv"
	"strings"

	"github.com/kellen-miller/aoc/go/pkg/io"
)

type CubeMaxes struct {
	Red   int
	Green int
	Blue  int
}

func GamesPossiblePowerSum(input string) int {
	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var sum int

	for sc.Scan() {
		maxes := parseRoundsPower(sc.Text())
		sum += maxes.Red * maxes.Green * maxes.Blue
	}

	return sum
}

func parseRoundsPower(line string) *CubeMaxes {
	const gamePartsWant = 2

	gameParts := strings.Split(line, ":")
	if len(gameParts) != gamePartsWant {
		panic("invalid game")
	}

	maxesSeen := new(CubeMaxes)
	for _, round := range strings.Split(gameParts[1], ";") {
		for _, cube := range strings.Split(round, ",") {
			parseCubePower(cube, maxesSeen)
		}
	}

	return maxesSeen
}

func parseCubePower(cube string, maxesSeen *CubeMaxes) {
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
		if cubes > maxesSeen.Red {
			maxesSeen.Red = cubes
		}
	case "green":
		if cubes > maxesSeen.Green {
			maxesSeen.Green = cubes
		}
	case "blue":
		if cubes > maxesSeen.Blue {
			maxesSeen.Blue = cubes
		}
	default:
		panic("unknown cube color")
	}
}
