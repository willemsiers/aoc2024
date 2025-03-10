package aoctools

import (
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const Separator = "\r\n"

func ReadInput(filePath string) string {
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

	return string(fileContents)
}

func PathRelativeToAOCRoot(parts ...string) string {
	p := append([]string{os.Getenv("GOPATH"), "src", "github.com", "aoc2024"}, parts...)
	return filepath.Join(p...)
}

// SplitColumns: specific to day 1. Maybe move it back.
func SplitColumns(text string) ([]int, []int) {
	lines := strings.Split(text, Separator)

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

func MapAtoi(ss []string) []int {
	result := make([]int, 0, len(ss))
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			logrus.Fatalf("failed to mapAtoi: %v", err)
		}
		result = append(result, i)
	}
	return result
}

func RemoveIndex(s []int, i int) []int {
	result := make([]int, 0, len(s)-1)
	result = append(result, s[:i]...)
	return append(result, s[i+1:]...)
}
