package main

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/willemsiers/aoc2024/aoctools"
	// . "github.com/onsi/gomega"
)

func Test_Day2_Part1_Example(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "2", "input_example.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(2))
}

func Test_Day2_Part1(t *testing.T) {
	g := NewWithT(t)

	result := part1(aoctools.PathRelativeToAOCRoot("days", "2", "input.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(379))
}

func Test_Day2_Part2_Example(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "2", "input_example.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(4))
}

func Test_Day2_Part2(t *testing.T) {
	g := NewWithT(t)

	result := part2(aoctools.PathRelativeToAOCRoot("days", "2", "input.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(430))
}
