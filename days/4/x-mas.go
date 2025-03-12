package main

import (
	"github.com/willemsiers/aoc2024/aoctools"
	"strings"
)

func part1(filePath string) int {
	input := aoctools.ReadInput(filePath)
	lines := strings.Split(input, aoctools.Separator)
	get := getterFunc(lines)

	solutions := 0
	for y := range lines {
		for x := range lines[y] {
			east := get(x, y) == 'X' && get(x+1, y) == 'M' && get(x+2, y) == 'A' && get(x+3, y) == 'S'
			west := get(x, y) == 'X' && get(x-1, y) == 'M' && get(x-2, y) == 'A' && get(x-3, y) == 'S'
			south := get(x, y) == 'X' && get(x, y+1) == 'M' && get(x, y+2) == 'A' && get(x, y+3) == 'S'
			north := get(x, y) == 'X' && get(x, y-1) == 'M' && get(x, y-2) == 'A' && get(x, y-3) == 'S'
			southEast := get(x, y) == 'X' && get(x+1, y+1) == 'M' && get(x+2, y+2) == 'A' && get(x+3, y+3) == 'S'
			northEast := get(x, y) == 'X' && get(x+1, y-1) == 'M' && get(x+2, y-2) == 'A' && get(x+3, y-3) == 'S'
			northWest := get(x, y) == 'X' && get(x-1, y-1) == 'M' && get(x-2, y-2) == 'A' && get(x-3, y-3) == 'S'
			southWest := get(x, y) == 'X' && get(x-1, y+1) == 'M' && get(x-2, y+2) == 'A' && get(x-3, y+3) == 'S'

			solutions += countTrue(east, west, south, north, southEast, northEast, northWest, southWest)
		}
	}
	return solutions
}

func part2(filePath string) int {
	input := aoctools.ReadInput(filePath)
	lines := strings.Split(input, aoctools.Separator)
	get := getterFunc(lines)

	solutions := 0
	for y := range lines {
		for x := range lines[y] {

			/**
			  fig 1 fig2  fig3  fig4
			  M.S   S.S   M.M   S.M
			  .A.   .A.   .A.   .A.
			  M.S   M.M   S.S   S.M
			*/

			// We search centered around A
			if get(x, y) != 'A' {
				continue
			}
			fig1 := get(x-1, y-1) == 'M' && get(x+1, y-1) == 'S' && get(x-1, y+1) == 'M' && get(x+1, y+1) == 'S'
			fig2 := get(x-1, y-1) == 'S' && get(x+1, y-1) == 'S' && get(x-1, y+1) == 'M' && get(x+1, y+1) == 'M'
			fig3 := get(x-1, y-1) == 'M' && get(x+1, y-1) == 'M' && get(x-1, y+1) == 'S' && get(x+1, y+1) == 'S'
			fig4 := get(x-1, y-1) == 'S' && get(x+1, y-1) == 'M' && get(x-1, y+1) == 'S' && get(x+1, y+1) == 'M'

			solutions += countTrue(fig1, fig2, fig3, fig4)
		}
	}
	return solutions
}

func getterFunc(lines []string) func(x int, y int) uint8 {
	get := func(x, y int) uint8 {
		if x < 0 || y < 0 || len(lines)-1 < y || len(lines[y])-1 < x {
			return '0' // pretend error, because '0' doesn't occur in our input.
		}
		return lines[y][x]
	}
	return get
}

func countTrue(bools ...bool) int {
	count := 0
	for _, b := range bools {
		if b {
			count++
		}
	}
	return count
}
