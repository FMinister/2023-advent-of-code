package day03

import (
	"log"
	"regexp"
	"strconv"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

func Day03_1() int {
	fileText := hf.OpenAndReadFile("./day03/input.txt")

	reSpecial := regexp.MustCompile(`[*/@\&=+#%$-]`)
	reDigit := regexp.MustCompile(`\d{1,}`)

	specials := make([][]int, len(fileText))

	sum := 0

	for i, line := range fileText {
		matches := reSpecial.FindAllStringIndex(line, -1)

		for _, match := range matches {
			specials[i] = append(specials[i], match[0])
		}
	}

	log.Printf("Specials: %v", specials)

	for i, line := range fileText {
		matches := reDigit.FindAllStringSubmatchIndex(line, -1)

		for _, match := range matches {
			if i > 0 {
				for _, special := range specials[i-1] {
					if match[0]-1 == special || match[0] == special || match[0]+1 == special || match[1] == special || match[1]-1 == special {
						number, _ := strconv.Atoi(line[match[0]:match[1]])
						sum = sum + number
					}
				}
			}
			if i < len(fileText)-1 {
				for _, special := range specials[i+1] {
					if match[0]-1 == special || match[0] == special || match[0]+1 == special || match[1] == special || match[1]-1 == special {
						number, _ := strconv.Atoi(line[match[0]:match[1]])
						sum = sum + number
					}
				}
			}
			for _, special := range specials[i] {
				if match[0]-1 == special || match[1] == special {
					number, _ := strconv.Atoi(line[match[0]:match[1]])
					sum = sum + number
				}
			}
		}
	}

	return sum
}
