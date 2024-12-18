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

	var foundYear bool
	for _, adventYear := range internal.Years() {
		if yearVal == all {
			if err := runYear(adventYear, dayVal, partVal); err != nil {
				panic(err)
			}

			foundYear = true
			continue
		}

		yearInt, err := strconv.Atoi(yearVal)
		if err != nil {
			panic(err)
		}

		if adventYear.AdventYear() == yearInt {
			if err := runYear(adventYear, dayVal, partVal); err != nil {
				panic(err)
			}

			foundYear = true
		}
	}

	if !foundYear {
		panic("year not found or not registered")
	}
}

func runYear(adventYear advent.Year, dayVal string, partVal string) error {
	var foundDay bool
	for _, adventDay := range adventYear.AdventDays() {
		if dayVal == all {
			if err := runDay(adventDay, adventYear.AdventYear(), adventDay.AdventDay(), partVal); err != nil {
				return err
			}

			foundDay = true
			continue
		}

		dayInt, err := strconv.Atoi(dayVal)
		if err != nil {
			return err
		}

		if adventDay.AdventDay() == dayInt {
			return runDay(adventDay, adventYear.AdventYear(), adventDay.AdventDay(), partVal)
		}
	}

	if !foundDay {
		return errors.New("day not found or not registered")
	}

	return nil
}

func runDay(adventDay advent.Day, year int, day int, part string) error {
	parts := []string{part}
	if part == all {
		parts = []string{"1", "2"}
	}

	for _, p := range parts {
		fp := filepath.Join("internal", "year"+strconv.Itoa(year), "day"+strconv.Itoa(day),
			fmt.Sprintf("input%s.txt", p))
		sc, closeFile := io.GetScanner(fp)
		var (
			result string
			err    error
			stop   time.Time
			start  = time.Now()
		)
		switch p {
		case "1":
			result, err = adventDay.Part1(sc)
		case "2":
			result, err = adventDay.Part2(sc)
		default:
			return errors.New("invalid part: " + p)
		}
		if err != nil {
			closeFile()
			return err
		}

		stop = time.Now()
		slog.Info(result,
			slog.Int("year", year),
			slog.Int("day", day),
			slog.String("part", p),
			slog.Duration("duration", stop.Sub(start)),
		)
		closeFile()
	}

	return nil
}
