package main

import (
	"github.com/willemsiers/aoc2024/aoctools"
)

func part1(filePath string) int {
	column1, column2 := aoctools.SplitColumns(aoctools.ReadInput(filePath))

	diffSum := 0
	for i := 0; i < len(column1); i++ {
		diff := column1[i] - column2[i]
		if diff < 0 {
			diffSum -= diff
		} else {
			diffSum += diff
		}
	}

	return diffSum
}

func part2(filePath string) int {
	column1, column2 := aoctools.SplitColumns(aoctools.ReadInput(filePath))

	occurrencesInColumn2 := map[int]int{}
	for _, id := range column2 {
		occurrencesInColumn2[id]++
	}
	similarity := 0
	for _, id := range column1 {
		similarity += id * occurrencesInColumn2[id]
	}
	return similarity
}
