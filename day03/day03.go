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

func Day03_2() int {
	fileText := hf.OpenAndReadFile("./day03/input.txt")

	reSpecial := regexp.MustCompile(`[*]`)
	reDigit := regexp.MustCompile(`\d{1,}`)

	specials := make([][]int, len(fileText))

	sum := 0

	for i, line := range fileText {
		matches := reSpecial.FindAllStringIndex(line, -1)

		for _, match := range matches {
			specials[i] = append(specials[i], match[0])
		}
	}

	for i, special := range specials {
		if len(special) == 0 {
			continue
		}

		for _, spec := range special {
			matchesAbove := reDigit.FindAllStringSubmatchIndex(fileText[i-1][spec-1:spec+2], -1)
			matchesBesides := reDigit.FindAllStringSubmatchIndex(fileText[i][spec-1:spec+2], -1)
			matchesBelow := reDigit.FindAllStringSubmatchIndex(fileText[i+1][spec-1:spec+2], -1)

			if len(matchesAbove) == 2 {
				matchAbove := reDigit.FindAllStringSubmatchIndex(fileText[i-1][spec-3:spec+4], -1)

				match1, _ := strconv.Atoi(fileText[i-1][spec-3+matchAbove[0][0] : spec-3+matchAbove[0][1]])
				match2, _ := strconv.Atoi(fileText[i-1][spec-3+matchAbove[1][0] : spec-3+matchAbove[1][1]])

				sum = sum + match1*match2
			}
			if len(matchesBesides) == 2 {
				matchBeside := reDigit.FindAllStringSubmatchIndex(fileText[i][spec-3:spec+4], -1)

				match1, _ := strconv.Atoi(fileText[i][spec-3+matchBeside[0][0] : spec-3+matchBeside[0][1]])
				match2, _ := strconv.Atoi(fileText[i][spec-3+matchBeside[1][0] : spec-3+matchBeside[1][1]])

				sum = sum + match1*match2
			}
			if len(matchesBelow) == 2 {
				matchBelow := reDigit.FindAllStringSubmatchIndex(fileText[i+1][spec-3:spec+4], -1)

				match1, _ := strconv.Atoi(fileText[i+1][spec-3+matchBelow[0][0] : spec-3+matchBelow[0][1]])
				match2, _ := strconv.Atoi(fileText[i+1][spec-3+matchBelow[1][0] : spec-3+matchBelow[1][1]])

				sum = sum + match1*match2
			}

			if len(matchesAbove) == 1 && len(matchesBesides) == 1 {
				matchAbove := reDigit.FindAllStringSubmatchIndex(fileText[i-1][spec-3:spec+4], -1)
				matchBeside := reDigit.FindAllStringSubmatchIndex(fileText[i][spec-3:spec+4], -1)

				if len(matchAbove) > 1 {
					if matchAbove[0][1] >= 3 {
						matchAbove = matchAbove[:1]
					} else {
						matchAbove = matchAbove[1:]
					}
				}

				match1, _ := strconv.Atoi(fileText[i-1][spec-3+matchAbove[0][0] : spec-3+matchAbove[0][1]])
				match2, _ := strconv.Atoi(fileText[i][spec-3+matchBeside[0][0] : spec-3+matchBeside[0][1]])

				sum = sum + match1*match2
			} else if len(matchesAbove) == 1 && len(matchesBelow) == 1 {
				matchAbove := reDigit.FindAllStringSubmatchIndex(fileText[i-1][spec-3:spec+4], -1)
				matchBelow := reDigit.FindAllStringSubmatchIndex(fileText[i+1][spec-3:spec+4], -1)

				if len(matchAbove) > 1 {
					if matchAbove[0][1] >= 3 {
						matchAbove = matchAbove[:1]
					} else {
						matchAbove = matchAbove[1:]
					}
				}
				if len(matchBelow) > 1 {
					if matchBelow[0][1] >= 3 {
						matchBelow = matchBelow[:1]
					} else {
						matchBelow = matchBelow[1:]
					}
				}

				match1, _ := strconv.Atoi(fileText[i-1][spec-3+matchAbove[0][0] : spec-3+matchAbove[0][1]])
				match2, _ := strconv.Atoi(fileText[i+1][spec-3+matchBelow[0][0] : spec-3+matchBelow[0][1]])

				sum = sum + match1*match2
			} else if len(matchesBesides) == 1 && len(matchesBelow) == 1 {
				matchBelow := reDigit.FindAllStringSubmatchIndex(fileText[i+1][spec-3:spec+4], -1)
				matchBeside := reDigit.FindAllStringSubmatchIndex(fileText[i][spec-3:spec+4], -1)

				if len(matchBelow) > 1 {
					if matchBelow[0][1] >= 3 {
						matchBelow = matchBelow[:1]
					} else {
						matchBelow = matchBelow[1:]
					}
				}

				match1, _ := strconv.Atoi(fileText[i+1][spec-3+matchBelow[0][0] : spec-3+matchBelow[0][1]])
				match2, _ := strconv.Atoi(fileText[i][spec-3+matchBeside[0][0] : spec-3+matchBeside[0][1]])

				sum = sum + match1*match2
			}
		}
	}

	return sum
}
