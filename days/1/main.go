package main

import (
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("day 1")
	inputFilePath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "aoc2024", "days", "1", "input.txt")

	logrus.Infof("result part 1: %v (total distance)", part1(inputFilePath))
	logrus.Infof("result part 2: %v (similarity score)", part2(inputFilePath))
}

func part1(filePath string) int {
	column1, column2 := readColumns(filePath)

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
	column1, column2 := readColumns(filePath)
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

func readColumns(filePath string) ([]int, []int) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		logrus.Fatalf("failed to stat: %v", err)
	}

	if fileInfo.IsDir() {
		logrus.Fatalf("input is dir")
	}

	fileContents := make([]byte, fileInfo.Size())
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	n, err := file.Read(fileContents)
	if err != nil {
		logrus.Fatalf("failed to read file: %v", err)
	}

	if int64(n) != fileInfo.Size() {
		logrus.Fatalf("read %d out of %d bytes", n, fileInfo.Size())
	}

	fileContentsString := string(fileContents)

	const separator = "\r\n"
	lines := strings.Split(fileContentsString, separator)

	var column1, column2 []int

	for _, l := range lines {
		columns := strings.Split(l, "   ")
		if len(columns) != 2 {
			logrus.Fatalf("got %d columns", len(columns))
		}
		i1, err := strconv.Atoi(columns[0])
		if err != nil {
			logrus.Fatalf("invalid int")
		}
		i2, err := strconv.Atoi(columns[1])
		if err != nil {
			logrus.Fatalf("invalid int")
		}
		column1 = append(column1, i1)
		column2 = append(column2, i2)
	}

	slices.Sort(column1)
	slices.Sort(column2)

	if len(column1) != len(column2) {
		logrus.Fatalf("column length mismatch")
	}

	return column1, column2
}
