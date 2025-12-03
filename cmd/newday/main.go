package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const (
	secureDirPerm  = 0o750
	secureFilePerm = 0o600
	commandTimeout = 30 * time.Second
)

type templateData struct {
	YearString  string
	DayString   string
	DayPadded   string
	PackageName string
	Language    string
	Year        int
	Day         int
}

func main() {
	var (
		yearFlag = flag.Int("year", 0, "Advent year")
		dayFlag  = flag.Int("day", 0, "Day number (1-25)")
		langFlag = flag.String("lang", "go", "Language key (e.g. go)")
	)
	flag.Parse()

	if err := run(*yearFlag, *dayFlag, strings.TrimSpace(*langFlag)); err != nil {
		slog.Error("failed to scaffold day", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func run(year, day int, lang string) error {
	if year <= 0 {
		return errors.New("year must be provided and > 0")
	}
	if day <= 0 {
		return errors.New("day must be provided and > 0")
	}
	if lang == "" {
		return errors.New("language must be provided")
	}

	templateDir := filepath.Join("templates", lang)
	if info, err := os.Stat(templateDir); err != nil || !info.IsDir() {
		return fmt.Errorf("language templates not found: %s", templateDir)
	}

	dayDirName := fmt.Sprintf("day%d", day)
	destDir := filepath.Join(strconv.Itoa(year), lang, dayDirName)

	if _, err := os.Stat(destDir); err == nil {
		return fmt.Errorf("destination already exists: %s", destDir)
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err := ensureSecureDir(destDir); err != nil {
		return fmt.Errorf("create destination: %w", err)
	}

	data := templateData{
		Year:        year,
		YearString:  strconv.Itoa(year),
		Day:         day,
		DayString:   strconv.Itoa(day),
		DayPadded:   fmt.Sprintf("%02d", day),
		PackageName: dayDirName,
		Language:    lang,
	}

	generatedFiles, err := hydrateTemplates(templateDir, destDir, &data)
	if err != nil {
		return err
	}

	if lang == "go" {
		if err := gofmtFiles(generatedFiles.goFiles); err != nil {
			return err
		}
		if err := runCommand("go", "generate", "./languages/go/cmd/advent"); err != nil {
			return err
		}
	}

	slog.Info("created day", slog.String("language", lang), slog.Int("year", year), slog.Int("day", day))
	return nil
}

type generatedSet struct {
	goFiles []string
}

func hydrateTemplates(templateDir, destDir string, data *templateData) (*generatedSet, error) {
	result := &generatedSet{}
	walker := func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}

		return renderTemplateFile(templateDir, destDir, path, data, result)
	}

	if err := filepath.WalkDir(templateDir, walker); err != nil {
		return nil, err
	}

	return result, nil
}

func gofmtFiles(files []string) error {
	if len(files) == 0 {
		return nil
	}

	args := append([]string{"-w"}, files...)
	return runCommand("gofmt", args...)
}

func runCommand(name string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func renderTemplateFile(templateDir, destDir, path string, data *templateData, result *generatedSet) error {
	rel, err := filepath.Rel(templateDir, path)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(rel, ".tmpl") {
		return fmt.Errorf("template files must use .tmpl extension: %s", rel)
	}

	destRel := strings.TrimSuffix(rel, ".tmpl")
	destPath := filepath.Join(destDir, destRel)
	if err := ensureSecureDir(filepath.Dir(destPath)); err != nil {
		return fmt.Errorf("prepare directory for %s: %w", destPath, err)
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tpl, err := template.New(rel).Parse(string(contents))
	if err != nil {
		return fmt.Errorf("parse template %s: %w", rel, err)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("render template %s: %w", rel, err)
	}

	if err := os.WriteFile(destPath, buf.Bytes(), secureFilePerm); err != nil {
		return fmt.Errorf("write file %s: %w", destPath, err)
	}

	if strings.HasSuffix(destPath, ".go") {
		result.goFiles = append(result.goFiles, destPath)
	}

	return nil
}

func ensureSecureDir(path string) error {
	return os.MkdirAll(path, secureDirPerm)
}
