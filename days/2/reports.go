package main

import (
	"strings"

	"github.com/willemsiers/aoc2024/aoctools"
)

func part1(filePath string) int {
	input := aoctools.ReadInput(filePath)
	reports := strings.Split(input, aoctools.Separator)

	safeCount := 0

	for _, r := range reports {
		levelsStr := strings.Split(r, " ")
		levels := aoctools.MapAtoi(levelsStr)

		if len(r) == 1 {
			safeCount++
			continue
		}

		increasing := (levels[1] - levels[0]) > 0 // true if seq is increasing

		safe := true

		for i := 1; i < len(levels); i++ {

			diff := levels[i] - levels[i-1]

			diffAbs := max(diff, -diff)

			if diffAbs > 3 || diffAbs < 1 { // step must be at least 1, at most 3.
				safe = false
				break
			}

			if (increasing && diff < 0) || (!increasing && diff > 0) { // it must be monotonic
				safe = false
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	return safeCount
}

func part2(filePath string) int {
	input := aoctools.ReadInput(filePath)
	reports := strings.Split(input, aoctools.Separator)

	safeCount := 0

	for _, r := range reports {

		levelsStr := strings.Split(r, " ")
		levels := aoctools.MapAtoi(levelsStr)

		// Either safe without any levels removed.
		if isSafe(levels) {
			safeCount++
			continue
		}

		// Or safe with 1 level removed.
		for i := 0; i < len(levels); i++ {
			reportWithOneLevelRemoved := aoctools.RemoveIndex(levels, i)
			if isSafe(reportWithOneLevelRemoved) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func isSafe(levels []int) bool {
	if len(levels) == 1 {
		return true
	}

	increasing := (levels[1] - levels[0]) > 0 // true if seq is increasing

	safe := true
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		diffAbs := max(diff, -diff)

		if diffAbs > 3 || diffAbs < 1 { // step must be at least 1, at most 3.
			safe = false
			break
		}

		if (increasing && diff < 0) || (!increasing && diff > 0) { // it must be monotonic
			safe = false
			break
		}
	}
	return safe

}
