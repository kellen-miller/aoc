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

	"github.com/kellen-miller/advent-of-code/go/internal"
	"github.com/kellen-miller/advent-of-code/go/internal/year2023"
	"github.com/kellen-miller/advent-of-code/go/internal/year2024"
	"github.com/kellen-miller/advent-of-code/go/pkg/io"
	"github.com/lmittmann/tint"
)

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

	flag.StringVar(&yearVal, "year", "all", "year to run")
	flag.StringVar(&dayVal, "day", "all", "day to run")
	flag.StringVar(&partVal, "part", "all", "part to run")

	flag.Parse()

	adventYears := []internal.AdventYear{
		new(year2023.Year),
		new(year2024.Year),
	}

	for _, adventYear := range adventYears {
		if yearVal == "all" {
			if err := runYear(adventYear, dayVal, partVal); err != nil {
				panic(err)
			}

			continue
		}

		yearInt, err := strconv.Atoi(yearVal)
		if err != nil {
			panic(err)
		}

		if adventYear.Year() == yearInt {
			if err := runYear(adventYear, dayVal, partVal); err != nil {
				panic(err)
			}
		}
	}
}

func runYear(adventYear internal.AdventYear, dayVal string, partVal string) error {
	for _, adventDay := range adventYear.AdventDays() {
		if dayVal == "all" {
			if err := runDay(adventDay, adventYear.Year(), adventDay.Day(), partVal); err != nil {
				return err
			}

			continue
		}

		dayInt, err := strconv.Atoi(dayVal)
		if err != nil {
			return err
		}

		if adventDay.Day() == dayInt {
			return runDay(adventDay, adventYear.Year(), adventDay.Day(), partVal)
		}
	}

	return nil
}

func runDay(adventDay internal.AdventDay, year int, day int, part string) error {
	fp := filepath.Join("internal", "year"+strconv.Itoa(year), "day"+strconv.Itoa(day),
		fmt.Sprintf("input%s.txt", part))
	sc, closeFile := io.GetScanner(fp)
	defer closeFile()

	var (
		p1Result string
		p2Result string
		err      error
	)
	switch part {
	case "all":
		p1Result, err = adventDay.Part1(sc)
		if err != nil {
			break
		}

		sc2, closeFile2 := io.GetScanner(fp)
		defer closeFile2()

		p2Result, err = adventDay.Part2(sc2)
	case "1":
		p1Result, err = adventDay.Part1(sc)
	case "2":
		p2Result, err = adventDay.Part2(sc)
	default:
		return errors.New("invalid part: " + part)
	}
	if err != nil {
		return err
	}

	if p1Result != "" {
		slog.Info(p1Result,
			slog.Int("year", year),
			slog.Int("day", day),
			slog.String("part", "1"),
		)
	}

	if p2Result != "" {
		slog.Info(p2Result,
			slog.Int("year", year),
			slog.Int("day", day),
			slog.String("part", "2"),
		)
	}

	return nil
}
