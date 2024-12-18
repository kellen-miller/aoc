package day9

import (
	"bufio"
	"strconv"
)

type DiskSegment struct {
	ID    int
	Start int
	Space int
}

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
	var (
		layout    []int
		files     []*DiskSegment
		freeSpace []*DiskSegment
	)
	for sc.Scan() {
		l, f, fs := parseLayout2(sc.Text())
		layout = append(layout, l...)
		files = append(files, f...)
		freeSpace = append(freeSpace, fs...)
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(compact2(layout, files, freeSpace)), nil
}

func parseLayout2(line string) ([]int, []*DiskSegment, []*DiskSegment) {
	var (
		id        int
		layout    []int
		files     []*DiskSegment
		freeSpace []*DiskSegment
	)
	for i, c := range line {
		space := int(c - '0')
		if i%2 == 0 {
			files = append(files, &DiskSegment{
				ID:    id,
				Start: len(layout),
				Space: space,
			})

			for range space {
				layout = append(layout, id)
			}

			id++
		} else {
			freeSpace = append(freeSpace, &DiskSegment{
				ID:    -1,
				Start: len(layout),
				Space: space,
			})

			for range space {
				layout = append(layout, -1)
			}
		}
	}

	return layout, files, freeSpace
}

func compact2(layout []int, files []*DiskSegment, freeSpace []*DiskSegment) int {
	for i := len(files) - 1; i >= 0; i-- {
		for _, fs := range freeSpace {
			f := files[i]
			if f.Start < fs.Start {
				break
			}

			if f.Space > fs.Space {
				continue
			}

			for j := range f.Space {
				layout[fs.Start+j] = f.ID
				layout[f.Start+j] = -1
			}

			fs.Start += f.Space
			fs.Space -= f.Space
			break
		}
	}

	var checksum int
	for i := range layout {
		if layout[i] != -1 {
			checksum += i * layout[i]
		}
	}

	return checksum
}
