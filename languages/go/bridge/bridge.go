package bridge

import (
	"bufio"
	"fmt"
	"os"
)

// MaterializeScanner copies the contents of a scanner into a temporary file so legacy
// solutions that expect to read from disk can be reused without refactoring.
func MaterializeScanner(sc *bufio.Scanner, prefix string) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", fmt.Sprintf("aoc-%s-", prefix))
	if err != nil {
		return "", nil, err
	}

	cleanup := func() {
		name := tmpFile.Name()
		if err := os.Remove(name); err != nil {
			// best-effort cleanup; ignore error
		}
	}

	writer := bufio.NewWriter(tmpFile)
	for sc.Scan() {
		if _, err := writer.WriteString(sc.Text()); err != nil {
			_ = writer.Flush()
			_ = tmpFile.Close()
			cleanup()
			return "", nil, err
		}
		if err := writer.WriteByte('\n'); err != nil {
			_ = writer.Flush()
			_ = tmpFile.Close()
			cleanup()
			return "", nil, err
		}
	}

	if err := sc.Err(); err != nil {
		_ = writer.Flush()
		_ = tmpFile.Close()
		cleanup()
		return "", nil, err
	}

	if err := writer.Flush(); err != nil {
		_ = tmpFile.Close()
		cleanup()
		return "", nil, err
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, err
	}

	return tmpFile.Name(), cleanup, nil
}
