package bridge

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

// MaterializeScanner copies the contents of a scanner into a temporary file so legacy
// solutions that expect to read from disk can be reused without refactoring.
func MaterializeScanner(sc *bufio.Scanner, prefix string) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", fmt.Sprintf("aoc-%s-", prefix))
	if err != nil {
		return "", nil, fmt.Errorf("create temp file: %w", err)
	}

	cleanup := makeCleanup(tmpFile.Name())

	if err := copyScannerToFile(sc, tmpFile); err != nil {
		return failTempFile(tmpFile, cleanup, err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

func copyScannerToFile(sc *bufio.Scanner, file *os.File) error {
	writer := bufio.NewWriter(file)
	for sc.Scan() {
		if _, err := writer.WriteString(sc.Text()); err != nil {
			return fmt.Errorf("write scanner line: %w", err)
		}
		if err := writer.WriteByte('\n'); err != nil {
			return fmt.Errorf("write newline: %w", err)
		}
	}

	if err := sc.Err(); err != nil {
		return fmt.Errorf("scan input: %w", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush writer: %w", err)
	}

	return nil
}

func failTempFile(file *os.File, cleanup func(), failure error) (string, func(), error) {
	if closeErr := file.Close(); closeErr != nil {
		slog.Warn(
			"close temp file after failure",
			slog.String("path", file.Name()),
			slog.String("error", closeErr.Error()),
		)
	}
	cleanup()
	return "", nil, failure
}

func makeCleanup(path string) func() {
	return func() {
		if err := os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
			slog.Debug("remove temp file", slog.String("path", path), slog.String("error", err.Error()))
		}
	}
}
