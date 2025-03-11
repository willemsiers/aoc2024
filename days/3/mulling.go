package main

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/willemsiers/aoc2024/aoctools"
)

func part1(filePath string) int {
	input := aoctools.ReadInput(filePath)

	// example: mul(44,46)

	sumOfMuls := 0

	for i := 0; i < len(input); {

		cur := input[i]

		if cur == 'm' {
			cur, i = advance(input, i)
			if cur == 'u' {
				cur, i = advance(input, i)
				if cur == 'l' {
					cur, i = advance(input, i)
					if cur == '(' {
						cur, i = advance(input, i)

						number1Start := i              // incl
						number1End := number1Start     // excl
						for cur >= '0' && cur <= '9' { // is digit
							number1End++
							cur = input[number1End]
						}

						number1Str := input[number1Start:number1End]
						if len(number1Str) == 0 {
							continue
						}
						i += len(number1Str) // advance
						number1, err := strconv.Atoi(number1Str)
						if err != nil {
							logrus.Fatalf("invalid atoi at index %d on '%v'", i, number1Str)
						}

						if cur == ',' {
							cur, i = advance(input, i)

							number2Start := i              // incl
							number2End := number2Start     // excl
							for cur >= '0' && cur <= '9' { // is digit
								number2End++
								cur = input[number2End]
							}

							number2Str := input[number2Start:number2End]
							if len(number2Str) == 0 {
								continue
							}
							i += len(number2Str) // advance
							number2, err := strconv.Atoi(number2Str)
							if err != nil {
								logrus.Fatalf("invalid atoi at index %d on '%v'", i, number2Str)
							}

							if cur == ')' {
								i++ // advance
								mul := number1 * number2
								sumOfMuls += mul
								logrus.Debugf("found %d x %d = %d (total=%d)", number1, number2, mul, sumOfMuls)
								continue
							}
						}
					}
				}
			}
		}
		i++
	}
	return sumOfMuls
}

func part2(filePath string) int {
	input := aoctools.ReadInput(filePath)

	// example: mul(44,46)
	// example: do()
	// example: don't()

	isMulEnabled := true
	sumOfMuls := 0

	for i := 0; i < len(input); {

		cur := input[i]
		// lex mul(<num>,<num>)
		if cur == 'm' {
			cur, i = advance(input, i)
			if cur == 'u' {
				cur, i = advance(input, i)
				if cur == 'l' {
					cur, i = advance(input, i)
					if cur == '(' {
						cur, i = advance(input, i)

						number1Start := i              // incl
						number1End := number1Start     // excl
						for cur >= '0' && cur <= '9' { // is digit
							number1End++
							cur = input[number1End]
						}

						number1Str := input[number1Start:number1End]
						if len(number1Str) == 0 {
							continue
						}
						i += len(number1Str) // advance
						number1, err := strconv.Atoi(number1Str)
						if err != nil {
							logrus.Fatalf("invalid atoi at index %d on '%v'", i, number1Str)
						}

						if cur == ',' {
							cur, i = advance(input, i)

							number2Start := i              // incl
							number2End := number2Start     // excl
							for cur >= '0' && cur <= '9' { // is digit
								number2End++
								cur = input[number2End]
							}

							number2Str := input[number2Start:number2End]
							if len(number2Str) == 0 {
								continue
							}
							i += len(number2Str) // advance
							number2, err := strconv.Atoi(number2Str)
							if err != nil {
								logrus.Fatalf("invalid atoi at index %d on '%v'", i, number2Str)
							}

							if cur == ')' {
								i++ // advance
								mul := number1 * number2
								if isMulEnabled {
									logrus.Debugf("found %d x %d = %d (total=%d)", number1, number2, mul, sumOfMuls)
									sumOfMuls += mul
								} else {
									logrus.Debugf("found ignored mul")
								}

								continue
							}
							continue
						}
						continue
					}
					continue
				}
				continue
			}
			continue
		}

		// lex do()/don't()
		if cur == 'd' {
			cur, i = advance(input, i)
			if cur == 'o' {
				cur, i = advance(input, i)

				if cur == '(' {
					cur, i = advance(input, i)
					if cur == ')' {
						cur, i = advance(input, i)
						logrus.Debugf("found do()")
						isMulEnabled = true
						continue
					}
					continue
				}

				if cur == 'n' {
					cur, i = advance(input, i)
					if cur == '\'' {
						cur, i = advance(input, i)
						if cur == 't' {
							cur, i = advance(input, i)
							if cur == '(' {
								cur, i = advance(input, i)
								if cur == ')' {
									cur, i = advance(input, i)
									logrus.Debugf("found don't()") // 𓀀
									isMulEnabled = false
									continue
								}
								continue
							}
							continue
						}
						continue
					}
					continue
				}
				continue
			}
			continue
		}
		i++
	}
	return sumOfMuls
}

func advance(input string, i int) (c uint8, newIndex int) {
	return input[i+1], i + 1
}
