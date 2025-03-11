package main

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/willemsiers/aoc2024/aoctools"
)

func Test_Day3_Part1_Example(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "3", "input_example_part1.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(161))
}

func Test_Day3_Part1(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "3", "input.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(167650499))
}

func Test_Day3_Part2_Example(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "3", "input_example_part2.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(48))
}

func Test_Day3_Part2(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "3", "input.txt"))
	t.Logf("result: %d", result)
	g.Expect(result).To(Equal(95846796))
}
