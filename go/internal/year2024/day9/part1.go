package day9

import (
	"bufio"
	"strconv"
)

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
	var layout []int
	for sc.Scan() {
		layout = append(layout, parseLayout(sc.Text())...)
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(compact(layout)), nil
}

func parseLayout(line string) []int {
	var (
		id     int
		layout []int
	)
	for i, c := range line {
		space := int(c - '0')
		if i%2 == 0 {
			for range space {
				layout = append(layout, id)
			}

			id++
		} else {
			for range space {
				layout = append(layout, -1)
			}
		}
	}

	return layout
}

func compact(layout []int) int {
	var (
		start    = 0
		end      = len(layout) - 1
		checksum int
	)

	for start < end {
		if layout[start] != -1 {
			checksum += start * layout[start]
			start++
			continue
		}

		if layout[end] == -1 {
			end--
			continue
		}

		layout[start], layout[end] = layout[end], layout[start]
		checksum += start * layout[start]
		start++
		end--
	}

	if start == end {
		checksum += start * layout[start]
	}

	return checksum
}
