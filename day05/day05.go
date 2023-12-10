package day05

import (
	"regexp"
	"strings"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

type Map struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

func Day05_1() int {
	fileText := hf.OpenAndReadFile("./day05/input.txt")

	re := regexp.MustCompile(`(\d+)`)

	seedStrings := re.FindAllString(strings.Split(fileText[0], ":")[1], -1)
	seeds := []int{}
	for _, seedString := range seedStrings {
		seeds = append(seeds, hf.StringToInt(seedString))
	}

	maps := [][]Map{}
	m := []Map{}

	for i := 3; i < len(fileText); i++ {
		line := fileText[i]

		xToYMap := re.FindAllString(line, -1)

		if len(xToYMap) == 0 {
			if len(m) > 0 {
				maps = append(maps, m)
				m = []Map{}
			}
			continue
		} else {
			destStart := hf.StringToInt(xToYMap[0])
			sourceStart := hf.StringToInt(xToYMap[1])
			rangeLen := hf.StringToInt(xToYMap[2])

			m = append(m, Map{destStart, sourceStart, rangeLen})
		}
	}

	maps = append(maps, m)

	locations := []int{}

	for _, seed := range seeds {
		for _, m := range maps {
			for _, n := range m {
				if n.sourceStart <= seed && seed < n.sourceStart+n.rangeLen {
					seed = n.destStart + seed - n.sourceStart
					break
				}
			}
		}

		locations = append(locations, seed)
	}

	min := locations[0]

	for _, location := range locations {
		if location < min {
			min = location
		}
	}

	return min
}

func Day05_2() int {
	fileText := hf.OpenAndReadFile("./day05/input.txt")

	re := regexp.MustCompile(`(\d+)`)

	seedStrings := re.FindAllString(strings.Split(fileText[0], ":")[1], -1)
	seeds := []Map{}
	for i := 0; i < len(seedStrings); i = i + 2 {
		destStart := hf.StringToInt(seedStrings[i])
		rangeLen := hf.StringToInt(seedStrings[i+1])
		seeds = append(seeds, Map{destStart, 0, rangeLen})
	}

	maps := [][]Map{}
	m := []Map{}

	for i := 3; i < len(fileText); i++ {
		line := fileText[i]

		xToYMap := re.FindAllString(line, -1)

		if len(xToYMap) == 0 {
			if len(m) > 0 {
				maps = append(maps, m)
				m = []Map{}
			}
			continue
		} else {
			destStart := hf.StringToInt(xToYMap[0])
			sourceStart := hf.StringToInt(xToYMap[1])
			rangeLen := hf.StringToInt(xToYMap[2])

			m = append(m, Map{destStart, sourceStart, rangeLen})
		}
	}

	maps = append(maps, m)

	locations := -1

	for _, seed := range seeds {
		for i := seed.destStart; i < seed.destStart+seed.rangeLen; i++ {
			currentSeed := i
			for _, m := range maps {
				for _, n := range m {
					if n.sourceStart <= currentSeed && currentSeed < n.sourceStart+n.rangeLen {
						currentSeed = n.destStart + currentSeed - n.sourceStart
						break
					}
				}
			}

			if locations == -1 {
				locations = currentSeed
			} else if locations > currentSeed {
				locations = currentSeed
			}
		}
	}

	return locations
}
