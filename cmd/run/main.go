package main

import (
	"context"
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

const (
	allKeyword       = "all"
	yearDigits       = 4
	runnerArgPadding = 8
	runnerTimeout    = 2 * time.Minute
)

type runnerConfig struct {
	Command string   `json:"command"`
	WorkDir string   `json:"work_dir"`
	Args    []string `json:"args"`
}

type target struct {
	Year     string
	Language string
}

type filterOptions struct {
	year string
	day  string
	part string
	lang string
}

func main() {
	opts := parseFilterFlags()
	configureLogger()

	if err := runApp(opts); err != nil {
		slog.Error("run failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func parseFilterFlags() filterOptions {
	var opts filterOptions
	flag.StringVar(&opts.year, "year", "", "year to run (defaults to all)")
	flag.StringVar(&opts.day, "day", "", "day to run (defaults to all)")
	flag.StringVar(&opts.part, "part", "", "part to run (defaults to all)")
	flag.StringVar(&opts.lang, "lang", "", "language to run (defaults to all)")
	flag.Parse()
	return opts
}

func configureLogger() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))
}

func runApp(opts filterOptions) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("determine working directory: %w", err)
	}

	normalizedYear := normalizeFilter(opts.year)
	normalizedDay := normalizeFilter(opts.day)
	normalizedPart := normalizeFilter(opts.part)
	normalizedLang := normalizeFilter(opts.lang)

	yearLanguages, err := discoverYearLanguages(cwd)
	if err != nil {
		return fmt.Errorf("discover year/language layout: %w", err)
	}

	targets, err := selectTargets(yearLanguages, normalizedYear, normalizedLang)
	if err != nil {
		return fmt.Errorf("select targets: %w", err)
	}
	if len(targets) == 0 {
		return fmt.Errorf("no matching solutions found for year %q and language %q", normalizedYear, normalizedLang)
	}

	return runTargets(cwd, targets, normalizedDay, normalizedPart)
}

func runTargets(root string, targets []target, day, part string) error {
	configs := make(map[string]runnerConfig)
	for _, target := range targets {
		cfg, err := ensureRunnerConfig(configs, root, target.Language)
		if err != nil {
			return err
		}

		slog.Info(
			"running",
			slog.String("language", target.Language),
			slog.String("year", target.Year),
			slog.String("day", fallbackValue(day, allKeyword)),
			slog.String("part", fallbackValue(part, allKeyword)),
		)

		if err := executeLanguage(cfg, target.Language, target.Year, day, part, root); err != nil {
			return fmt.Errorf("execute %s/%s: %w", target.Language, target.Year, err)
		}
	}

	return nil
}

func ensureRunnerConfig(cache map[string]runnerConfig, root, language string) (runnerConfig, error) {
	if cfg, ok := cache[language]; ok {
		return cfg, nil
	}

	cfg, err := loadRunnerConfig(root, language)
	if err != nil {
		return runnerConfig{}, fmt.Errorf("load runner config for %s: %w", language, err)
	}
	cache[language] = cfg
	return cfg, nil
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

func validateCommand(command string) error {
	if strings.ContainsRune(command, os.PathSeparator) {
		return fmt.Errorf("command %q must not include path separators", command)
	}

	allowed := map[string]struct{}{
		"bash":    {},
		"bun":     {},
		"cargo":   {},
		"deno":    {},
		"dotnet":  {},
		"go":      {},
		"java":    {},
		"kotlinc": {},
		"make":    {},
		"node":    {},
		"npm":     {},
		"pnpm":    {},
		"python":  {},
		"python3": {},
		"ruby":    {},
		"sh":      {},
		"swift":   {},
		"yarn":    {},
		"zig":     {},
	}

	if _, ok := allowed[command]; !ok {
		return fmt.Errorf("command %q is not in the allowed runner list", command)
	}

	return nil
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
	if len(name) != yearDigits {
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
	if err := validateCommand(cfg.Command); err != nil {
		return err
	}

	args := make([]string, 0, len(cfg.Args)+runnerArgPadding)
	args = append(args, cfg.Args...)
	args = append(args, "--year", year)
	if day != "" {
		args = append(args, "--day", day)
	}
	if part != "" {
		args = append(args, "--part", part)
	}
	args = append(args, "--lang", language)

	ctx, cancel := context.WithTimeout(context.Background(), runnerTimeout)
	defer cancel()

	// #nosec G204 - cfg.Command validated via validateCommand
	cmd := exec.CommandContext(ctx, cfg.Command, args...)
	cmd.Dir = filepath.Join(root, cfg.WorkDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}
