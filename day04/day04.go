package day04

import (
	"regexp"
	"strings"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

func Day04_1() int {
	fileText := hf.OpenAndReadFile("./day04/input.txt")

	re := regexp.MustCompile(`(\d+)`)

	sum := 0

	for _, line := range fileText {
		splitLine := strings.Split(line, ":")[1:]
		leftLine := strings.Split(splitLine[0], "|")[0]
		rightLine := strings.Split(splitLine[0], "|")[1]

		left := re.FindAllString(leftLine, -1)
		right := re.FindAllString(rightLine, -1)

		subSum := 0

		for _, leftNumber := range left {
			for _, rightNumber := range right {
				if leftNumber == rightNumber {
					if subSum == 0 {
						subSum = 1
					} else {
						subSum = subSum * 2
					}
				}
			}
		}

		sum = sum + subSum
	}

	return sum
}

func Day04_2() int {
	fileText := hf.OpenAndReadFile("./day04/input.txt")

	re := regexp.MustCompile(`(\d+)`)

	sum := 0

	var rules []string
	appearances := []int{}

	for _, line := range fileText {
		splitLine := strings.Split(line, ":")[1:]
		rules = append(rules, splitLine[0])
		appearances = append(appearances, 1)
	}

	for i := 0; i < len(appearances); i++ {
		leftLine := strings.Split(rules[i], "|")[0]
		rightLine := strings.Split(rules[i], "|")[1]

		left := re.FindAllString(leftLine, -1)
		right := re.FindAllString(rightLine, -1)

		subSum := 0

		for _, leftNumber := range left {
			for _, rightNumber := range right {
				if leftNumber == rightNumber {
					subSum = subSum + 1
				}
			}
		}

		for j := 0; j < appearances[i]; j++ {
			for k := 0; k < subSum; k++ {
				appearances[i+1+k] = appearances[i+1+k] + 1
			}
		}
	}

	for _, appearance := range appearances {
		sum = sum + appearance
	}

	return sum
}
