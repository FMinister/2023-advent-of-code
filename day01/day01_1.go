package day01

import (
	"unicode"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

func Day01_1() int {
	fileText := hf.OpenAndReadFile("./day01/input1.txt")

	sum := 0

	for _, line := range fileText {
		var numbers []int
		for _, char := range line {
			if unicode.IsDigit(char) {
				numbers = append(numbers, int(char-'0'))
			}
		}

		sum = sum + numbers[0]*10 + numbers[len(numbers)-1]
	}

	return sum
}
