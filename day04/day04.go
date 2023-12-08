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
