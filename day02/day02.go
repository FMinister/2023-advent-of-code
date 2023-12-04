package day02

import (
	"log"
	"regexp"
	"strconv"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

func Day02_1() int {
	fileText := hf.OpenAndReadFile("./day02/input.txt")

	bag := []int{12, 13, 14}
	re := regexp.MustCompile(`(\d+) (\w+)`)

	sum := 0

	for i, line := range fileText {
		colorCounts := make(map[string]int)
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			count := match[1]
			color := match[2]

			countInt, _ := strconv.Atoi(count)

			existingCount, ok := colorCounts[color]
			if !ok {
				colorCounts[color] = countInt
			} else if countInt > existingCount {
				colorCounts[color] = countInt
			}
		}

		colorArray := make([]int, 3)

		colorArray[0] = colorCounts["red"]
		colorArray[1] = colorCounts["green"]
		colorArray[2] = colorCounts["blue"]

		shouldNotAdd := false

		for j, colorCount := range colorArray {
			if colorCount > bag[j] {
				shouldNotAdd = true
			}
		}

		log.Printf("Line %d: %v; %v; !shouldNotAdd: %t\n", i+1, colorArray, colorCounts, !shouldNotAdd)

		if !shouldNotAdd {
			sum = sum + i + 1
		}
	}

	return sum
}
