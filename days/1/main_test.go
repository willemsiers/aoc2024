package main

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_Part1_Example(t *testing.T) {
	g := NewWithT(t)

	result := part1(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "aoc2024", "days", "1", "input_example.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(11))
}

func Test_Part1_Actual(t *testing.T) {
	g := NewWithT(t)

	result := part1(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "aoc2024", "days", "1", "input.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(2000468))
}

func Test_Part2_Example(t *testing.T) {
	g := NewWithT(t)

	result := part2(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "aoc2024", "days", "1", "input_example.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(31))
}

func Test_Part2_Actual(t *testing.T) {
	g := NewWithT(t)

	result := part2(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "aoc2024", "days", "1", "input.txt"))
	t.Logf("result: %v", result)
	g.Expect(result).To(Equal(18567089))
}
