package parts

import (
	"strings"

	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/pkg/io"
)

type file struct {
	name string
	size int
}

type folder struct {
	name     string
	parents  []folder
	children []folder
	files    []file
	size     int
}

func FindDirectoriesOfSizeTotal(input string) {
	if input == "" {
		input = internal.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var (
		fs      = make([]folder, 1)
		dirPath = make([]string, 1)
	)

	fs[0] = NewFolder("/") // root
	currentDir := fs[0]
	dirPath[0] = "/"

	for sc.Scan() {
		line := sc.Text()

		if strings.HasPrefix(line, "$") {
			cmdParts := strings.Split(line, " ")

			switch cmdParts[1] {
			case "cd":
				if cmdParts[2] == ".." {
					currentDir = currentDir.parents[len(currentDir.parents)-1]
				} else {
					for _, child := range currentDir.children {
						if child.name == cmdParts[2] {
							currentDir = child
							break
						}
					}
				}

			case "ls":

			}
		}
	}
}

func NewFolder(name string) folder {
	return folder{
		name:     name,
		parents:  make([]folder, 0),
		children: make([]folder, 0),
		files:    make([]file, 0),
	}
}
