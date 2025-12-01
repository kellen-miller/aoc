package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"time"

	"github.com/lmittmann/tint"
)

const all = "all"

//go:generate go run ./generate.go

type partRunner func(*bufio.Scanner) (string, error)

type dayRunners struct {
	Part1 partRunner
	Part2 partRunner
}

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	var (
		yearVal string
		dayVal  string
		partVal string
		langVal string
	)

	flag.StringVar(&yearVal, "year", all, "year to run")
	flag.StringVar(&dayVal, "day", all, "day to run")
	flag.StringVar(&partVal, "part", all, "part to run")
	flag.StringVar(&langVal, "lang", "go", "language identifier (noop)")

	flag.Parse()

	slog.Debug("executing go solutions",
		slog.String("lang", langVal),
		slog.String("year", yearVal),
		slog.String("day", dayVal),
		slog.String("part", partVal),
	)

	if err := run(yearVal, dayVal, partVal); err != nil {
		slog.Error("failure during advent run", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func run(yearArg, dayArg, partArg string) error {
	years, err := selectYears(yearArg)
	if err != nil {
		return err
	}

	parts, err := selectParts(partArg)
	if err != nil {
		return err
	}

	for _, year := range years {
		days, err := selectDays(year, dayArg)
		if err != nil {
			return err
		}

		for _, day := range days {
			if err := runDay(year, day, parts); err != nil {
				return err
			}
		}
	}

	return nil
}

func selectYears(arg string) ([]int, error) {
	if arg == all {
		return sortedYearKeys(solutionIndex), nil
	}

	year, err := strconv.Atoi(arg)
	if err != nil {
		return nil, fmt.Errorf("parse year %q: %w", arg, err)
	}

	if _, ok := solutionIndex[year]; !ok {
		return nil, fmt.Errorf("year %d not found", year)
	}

	return []int{year}, nil
}

func selectDays(year int, arg string) ([]int, error) {
	daysForYear, ok := solutionIndex[year]
	if !ok {
		return nil, fmt.Errorf("year %d not found", year)
	}

	if arg == all {
		return sortedDayKeys(daysForYear), nil
	}

	day, err := strconv.Atoi(arg)
	if err != nil {
		return nil, fmt.Errorf("parse day %q: %w", arg, err)
	}

	if _, ok := daysForYear[day]; !ok {
		return nil, fmt.Errorf("day %d not found for year %d", day, year)
	}

	return []int{day}, nil
}

func selectParts(arg string) ([]string, error) {
	if arg == all {
		return []string{"1", "2"}, nil
	}

	if arg != "1" && arg != "2" {
		return nil, fmt.Errorf("invalid part %s", arg)
	}

	return []string{arg}, nil
}

func runDay(year, day int, parts []string) error {
	runners := solutionIndex[year][day]

	for _, part := range parts {
		var fn partRunner
		switch part {
		case "1":
			fn = runners.Part1
		case "2":
			fn = runners.Part2
		default:
			return fmt.Errorf("unknown part %s", part)
		}

		if fn == nil {
			return fmt.Errorf("part %s not implemented for %d day %d", part, year, day)
		}

		start := time.Now()

		result, skipped, err := execPart(year, day, part, fn)
		if skipped {
			slog.Warn(result,
				slog.Int("year", year),
				slog.Int("day", day),
				slog.String("part", part),
			)
			continue
		}
		if err != nil {
			return err
		}

		slog.Info(result,
			slog.Int("year", year),
			slog.Int("day", day),
			slog.String("part", part),
			slog.Duration("duration", time.Since(start)),
		)
	}

	return nil
}

func execPart(year, day int, part string, fn partRunner) (string, bool, error) {
	inputPath := filepath.Join(
		strconv.Itoa(year),
		"go",
		fmt.Sprintf("day%d", day),
		fmt.Sprintf("input%s.txt", part),
	)

	sc, closeFile, err := openInputScanner(inputPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "missing input file: " + inputPath, true, nil
		}
		return "", false, err
	}
	defer closeFile()

	result, err := fn(sc)
	return result, false, err
}

func openInputScanner(path string) (*bufio.Scanner, func(), error) {
	if !filepath.IsAbs(path) {
		path = filepath.Clean(path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("open input %s: %w", path, err)
	}

	closeFunc := func() {
		if cerr := file.Close(); cerr != nil {
			slog.Warn("close input file", slog.String("path", path), slog.String("error", cerr.Error()))
		}
	}

	return bufio.NewScanner(file), closeFunc, nil
}

func sortedYearKeys(m map[int]map[int]dayRunners) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func sortedDayKeys(m map[int]dayRunners) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
