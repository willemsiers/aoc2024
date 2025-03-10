package main

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/willemsiers/aoc2024/aoctools"
)

func Test_Part1_Example(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "1", "input_example.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(11))
}

func Test_Part1_Actual(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "1", "input.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(2000468))
}

func Test_Part2_Example(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "1", "input_example.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(31))
}

func Test_Part2_Actual(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "1", "input.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(18567089))
}
