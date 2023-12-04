package day01

import (
	"regexp"
	"unicode"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func Day01_1() int {
	fileText := hf.OpenAndReadFile("./day01/input.txt")

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

func Day01_2() int {
	fileText := hf.OpenAndReadFile("./day01/input.txt")

	sum := 0
	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

	for _, line := range fileText {
		var numbers []string

		for i := range line {
			found := re.FindString(line[i:])
			if found != "" {
				numbers = append(numbers, found)
			}
		}

		if first, ok := digits[numbers[0]]; ok {
			sum = sum + first*10
		}

		if last, ok := digits[numbers[len(numbers)-1]]; ok {
			sum = sum + last
		}
	}

	return sum
}
