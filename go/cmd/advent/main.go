package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/internal/advent"
	"github.com/kellen-miller/aoc/go/pkg/io"
	"github.com/lmittmann/tint"
)

const all = "all"

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
	)

	flag.StringVar(&yearVal, "year", all, "year to run")
	flag.StringVar(&dayVal, "day", all, "day to run")
	flag.StringVar(&partVal, "part", all, "part to run")

	flag.Parse()

	if err := runAdvent(yearVal, dayVal, partVal); err != nil {
		slog.Error("failure during advent run", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func runAdvent(year string, day string, part string) error {
	var (
		allYears   = internal.Years()
		yearsToRun = make([]advent.Year, 0, len(allYears))
	)
	for _, y := range allYears {
		if year == all {
			yearsToRun = append(yearsToRun, y)
			continue
		}

		yearInt, err := strconv.Atoi(year)
		if err != nil {
			return err
		}

		if y.AdventYear() == yearInt {
			yearsToRun = append(yearsToRun, y)
		}
	}

	if len(yearsToRun) == 0 {
		return errors.New("year not found or not registered")
	}

	for _, adventYear := range yearsToRun {
		if err := runYear(adventYear, day, part); err != nil {
			return err
		}
	}

	return nil
}

func runYear(year advent.Year, day string, part string) error {
	var (
		allDays   = year.AdventDays()
		daysToRun = make([]advent.Day, 0, len(allDays))
	)
	for _, d := range allDays {
		if day == all {
			daysToRun = append(daysToRun, d)
			continue
		}

		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return err
		}

		if d.AdventDay() == dayInt {
			daysToRun = append(daysToRun, d)
		}
	}

	if len(daysToRun) == 0 {
		return errors.New("day not found or not registered")
	}

	for _, adventDay := range daysToRun {
		if err := runDay(year.AdventYear(), adventDay, part); err != nil {
			return err
		}
	}

	return nil
}

func runDay(year int, day advent.Day, part string) error {
	parts := []string{part}
	if part == all {
		parts = []string{"1", "2"}
	}

	for _, p := range parts {
		start := time.Now()

		result, err := execDay(year, day, p)
		if err != nil {
			return err
		}

		slog.Info(result,
			slog.Int("year", year),
			slog.Int("day", day.AdventDay()),
			slog.String("part", p),
			slog.Duration("duration", time.Now().Sub(start)),
		)
	}

	return nil
}

func execDay(year int, day advent.Day, part string) (string, error) {
	sc, closeFile := io.GetScanner(filepath.Join(
		"internal",
		"year"+strconv.Itoa(year),
		"day"+strconv.Itoa(day.AdventDay()),
		fmt.Sprintf("input%s.txt", part)))
	defer closeFile()

	switch part {
	case "1":
		return day.Part1(sc)
	case "2":
		return day.Part2(sc)
	default:
		return "", errors.New("invalid part: " + part)
	}
}
