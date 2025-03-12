package main

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/willemsiers/aoc2024/aoctools"
)

func Test_Day4_Part1_Example(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "4", "input_example.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(18))
}

func Test_Day4_Part1(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "4", "input.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(2593))
}

func Test_Day4_Part2_Example(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "4", "input_example.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(9))
}

func Test_Day4_Part2(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "4", "input.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(1950))
}
