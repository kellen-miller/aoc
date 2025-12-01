package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/lmittmann/tint"
)

const allKeyword = "all"

type runnerConfig struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
	WorkDir string   `json:"workdir"`
}

type target struct {
	Year     string
	Language string
}

func main() {
	var (
		yearFilter string
		dayFilter  string
		partFilter string
		langFilter string
	)

	flag.StringVar(&yearFilter, "year", "", "year to run (defaults to all)")
	flag.StringVar(&dayFilter, "day", "", "day to run (defaults to all)")
	flag.StringVar(&partFilter, "part", "", "part to run (defaults to all)")
	flag.StringVar(&langFilter, "lang", "", "language to run (defaults to all)")
	flag.Parse()

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))

	cwd, err := os.Getwd()
	if err != nil {
		slog.Error("unable to determine working directory", slog.String("error", err.Error()))
		os.Exit(1)
	}

	normalizedYear := normalizeFilter(yearFilter)
	normalizedDay := normalizeFilter(dayFilter)
	normalizedPart := normalizeFilter(partFilter)
	normalizedLang := normalizeFilter(langFilter)

	yearLanguages, err := discoverYearLanguages(cwd)
	if err != nil {
		slog.Error("failed to discover year/language layout", slog.String("error", err.Error()))
		os.Exit(1)
	}

	targets, err := selectTargets(yearLanguages, normalizedYear, normalizedLang)
	if err != nil {
		slog.Error("failed to select targets", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if len(targets) == 0 {
		slog.Error("no matching solutions found", slog.String("year", normalizedYear), slog.String("lang", normalizedLang))
		os.Exit(1)
	}

	configs := make(map[string]runnerConfig)
	for _, target := range targets {
		cfg, ok := configs[target.Language]
		if !ok {
			cfg, err = loadRunnerConfig(cwd, target.Language)
			if err != nil {
				slog.Error("failed to load runner config", slog.String("language", target.Language), slog.String("error", err.Error()))
				os.Exit(1)
			}
			configs[target.Language] = cfg
		}

		slog.Info("running", slog.String("language", target.Language), slog.String("year", target.Year), slog.String("day", fallbackValue(normalizedDay, allKeyword)), slog.String("part", fallbackValue(normalizedPart, allKeyword)))

		if err := executeLanguage(cfg, target.Language, target.Year, normalizedDay, normalizedPart, cwd); err != nil {
			slog.Error("run failed", slog.String("language", target.Language), slog.String("year", target.Year), slog.String("error", err.Error()))
			os.Exit(1)
		}
	}
}

func normalizeFilter(value string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return ""
	}
	if strings.EqualFold(trimmed, allKeyword) {
		return ""
	}
	return trimmed
}

func fallbackValue(value string, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

func discoverYearLanguages(root string) (map[string][]string, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	yearLanguages := make(map[string][]string)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !isYearDirectory(name) {
			continue
		}

		yearPath := filepath.Join(root, name)
		langEntries, err := os.ReadDir(yearPath)
		if err != nil {
			return nil, err
		}

		var langs []string
		for _, langEntry := range langEntries {
			if langEntry.IsDir() {
				langs = append(langs, langEntry.Name())
			}
		}

		if len(langs) == 0 {
			continue
		}

		slices.Sort(langs)
		yearLanguages[name] = langs
	}

	if len(yearLanguages) == 0 {
		return nil, errors.New("no year directories found")
	}

	return yearLanguages, nil
}

func isYearDirectory(name string) bool {
	if len(name) != 4 {
		return false
	}
	for _, r := range name {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func selectTargets(yearLanguages map[string][]string, yearFilter string, langFilter string) ([]target, error) {
	years := make([]string, 0, len(yearLanguages))
	for year := range yearLanguages {
		years = append(years, year)
	}
	slices.Sort(years)

	if yearFilter != "" {
		if _, ok := yearLanguages[yearFilter]; !ok {
			return nil, fmt.Errorf("year %s not found", yearFilter)
		}
		years = []string{yearFilter}
	}

	var targets []target
	var langFound bool

	for _, year := range years {
		langs := yearLanguages[year]
		for _, lang := range langs {
			if langFilter != "" && !strings.EqualFold(langFilter, lang) {
				continue
			}
			if langFilter != "" && strings.EqualFold(langFilter, lang) {
				langFound = true
			}
			targets = append(targets, target{Year: year, Language: lang})
		}
	}

	if langFilter != "" && !langFound {
		return nil, fmt.Errorf("language %s not found", langFilter)
	}

	return targets, nil
}

func loadRunnerConfig(root string, language string) (runnerConfig, error) {
	cfgPath := filepath.Join(root, "languages", language, "runner.json")
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return runnerConfig{}, err
	}

	var cfg runnerConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return runnerConfig{}, err
	}

	if cfg.Command == "" {
		return runnerConfig{}, errors.New("runner command is required")
	}

	if cfg.WorkDir == "" {
		cfg.WorkDir = "."
	}

	return cfg, nil
}

func executeLanguage(cfg runnerConfig, language string, year string, day string, part string, root string) error {
	args := make([]string, 0, len(cfg.Args)+8)
	args = append(args, cfg.Args...)
	args = append(args, "--year", year)
	if day != "" {
		args = append(args, "--day", day)
	}
	if part != "" {
		args = append(args, "--part", part)
	}
	args = append(args, "--lang", language)

	cmd := exec.Command(cfg.Command, args...)
	cmd.Dir = filepath.Join(root, cfg.WorkDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}
