package day2

import (
	"bufio"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	decimalBase    = 10
	minRepeatCount = 2
)

// This solution stays “range friendly” by precomputing every number that can be
// expressed as a repeated digit block (e.g., 123123 or 77). We then sort that set
// once and use prefix sums to answer “sum within range” queries with binary
// search rather than iterating over every raw ID in the input ranges.

// Part2 sums every invalid ID in the provided ranges by generating the invalid
// candidates once, then slicing into that sorted list per range.
func Part2(sc *bufio.Scanner) (string, error) {
	ranges, maxValue, maxDigits, err := readRanges(sc)
	if err != nil {
		return "", err
	}
	if len(ranges) == 0 {
		return "0", nil
	}

	invalidValues, prefix := generateInvalidValues(maxDigits, maxValue)
	var total int64
	for _, rg := range ranges {
		// Range sums are now cheap lookups over the sorted invalid list.
		total += sumInvalidInRange(invalidValues, prefix, rg.lo, rg.hi)
	}

	return strconv.FormatInt(total, 10), nil
}

// idRange stores the normalized inclusive bounds used throughout the solver.
type idRange struct {
	lo int64
	hi int64
}

// readRanges parses the comma-separated `lo-hi` entries, normalizes the bounds,
// and records the largest ID / digit width so later stages know how many
// repeated patterns need to be generated.
func readRanges(sc *bufio.Scanner) ([]idRange, int64, int, error) {
	var (
		maxDigits = 1
		ranges    []idRange
		maxValue  int64
	)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		parsed, lineMaxValue, lineMaxDigits, err := parseRangeLine(line)
		if err != nil {
			return nil, 0, 0, err
		}

		ranges = append(ranges, parsed...)
		if lineMaxValue > maxValue {
			maxValue = lineMaxValue
		}
		if lineMaxDigits > maxDigits {
			maxDigits = lineMaxDigits
		}
	}

	if err := sc.Err(); err != nil {
		return nil, 0, 0, fmt.Errorf("read input: %w", err)
	}

	return ranges, maxValue, maxDigits, nil
}

func parseRangeLine(line string) ([]idRange, int64, int, error) {
	parts := strings.SplitSeq(line, ",")
	var (
		lineMaxValue  int64
		lineMaxDigits = 1
	)
	parsed := make([]idRange, 0, strings.Count(line, ",")+1)

	for raw := range parts {
		part := strings.TrimSpace(raw)
		if part == "" {
			continue
		}

		lo, hi, err := parseBounds(part)
		if err != nil {
			return nil, 0, 0, err
		}

		parsed = append(parsed, idRange{lo: lo, hi: hi})
		if hi > lineMaxValue {
			lineMaxValue = hi
		}
		if digits := digitCount(hi); digits > lineMaxDigits {
			lineMaxDigits = digits
		}
	}

	return parsed, lineMaxValue, lineMaxDigits, nil
}

func parseBounds(part string) (int64, int64, error) {
	loStr, hiStr, ok := strings.Cut(part, "-")
	if !ok {
		return 0, 0, fmt.Errorf("could not parse id range: %s", part)
	}

	lo, err := strconv.ParseInt(strings.TrimSpace(loStr), decimalBase, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse id %s: %w", loStr, err)
	}

	hi, err := strconv.ParseInt(strings.TrimSpace(hiStr), decimalBase, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse id %s: %w", hiStr, err)
	}

	if lo > hi {
		lo, hi = hi, lo
	}

	return lo, hi, nil
}

func generateInvalidValues(maxDigits int, maxValue int64) ([]int64, []int64) {
	if maxDigits < minRepeatCount || maxValue <= 0 {
		return nil, nil
	}

	pow10 := buildPow10(maxDigits)
	invalid := buildRepeatedValues(maxDigits, maxValue, pow10)
	if len(invalid) == 0 {
		return nil, nil
	}

	slices.Sort(invalid)
	prefix := buildPrefixSums(invalid)
	return invalid, prefix
}

func buildPow10(maxDigits int) []int64 {
	pow10 := make([]int64, maxDigits+1)
	pow10[0] = 1
	for i := 1; i <= maxDigits; i++ {
		pow10[i] = pow10[i-1] * decimalBase
	}
	return pow10
}

func buildRepeatedValues(maxDigits int, maxValue int64, pow10 []int64) []int64 {
	limitBaseDigits := maxDigits / evenSplitDivisor
	seen := make(map[int64]struct{})
	var values []int64
	for baseDigits := 1; baseDigits <= limitBaseDigits; baseDigits++ {
		values = append(values, collectValuesForBaseDigits(baseDigits, maxDigits, maxValue, pow10, seen)...)
	}
	return values
}

func collectValuesForBaseDigits(
	baseDigits, maxDigits int,
	maxValue int64,
	pow10 []int64,
	seen map[int64]struct{},
) []int64 {
	minBase := pow10[baseDigits-1]
	maxBase := pow10[baseDigits] - 1
	var results []int64
	for repeats := minRepeatCount; baseDigits*repeats <= maxDigits; repeats++ {
		multiplier := repeatMultiplier(baseDigits, repeats, pow10)
		baseUpper := maxBase
		if div := maxValue / multiplier; div < baseUpper {
			baseUpper = div
		}

		if baseUpper < minBase {
			continue
		}

		for base := minBase; base <= baseUpper; base++ {
			value := base * multiplier
			if _, ok := seen[value]; ok {
				continue
			}
			seen[value] = struct{}{}
			results = append(results, value)
		}
	}

	return results
}

func buildPrefixSums(values []int64) []int64 {
	prefix := make([]int64, len(values))
	var running int64
	for i, v := range values {
		running += v
		prefix[i] = running
	}
	return prefix
}

// repeatMultiplier builds the tiling factor used to stamp the base digits
// repeatedly (e.g., baseDigits=2, repeats=3 -> multiplier 10101).
func repeatMultiplier(baseDigits, repeats int, pow10 []int64) int64 {
	step := pow10[baseDigits]
	multiplier := int64(0)
	factor := int64(1)
	for range repeats {
		multiplier += factor
		factor *= step
	}

	return multiplier
}

// sumInvalidInRange uses two binary searches to find the indices in `values`
// that fall inside [lo, hi] and returns the corresponding sum via the prefix
// array.
func sumInvalidInRange(values, prefix []int64, lo, hi int64) int64 {
	if len(values) == 0 || hi < lo {
		return 0
	}

	// Standard lower/upper-bound searches isolate the slice that intersects the
	// target range.
	left := sort.Search(len(values), func(i int) bool { return values[i] >= lo })
	right := sort.Search(len(values), func(i int) bool { return values[i] > hi })
	if left >= right {
		return 0
	}

	// Convert the index window into a sum via prefix differences.
	total := prefix[right-1]
	if left > 0 {
		total -= prefix[left-1]
	}

	return total
}

// digitCount returns the number of decimal digits in n (treating 0 specially).
func digitCount(n int64) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n > 0 {
		count++
		n /= decimalBase
	}

	return count
}
