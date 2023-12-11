package day06

import (
	"regexp"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

func Day06_1() int {
	fileText := hf.OpenAndReadFile("./day06/input.txt")

	time := []int{}
	distance := []int{}
	re := regexp.MustCompile(`(\d+)`)

	for i := 0; i < len(fileText); i++ {
		line := fileText[i]

		timeOrDist := re.FindAllString(line, -1)

		if len(time) == 0 {
			time = hf.StringsToInts(timeOrDist)
		} else {
			distance = hf.StringsToInts(timeOrDist)
		}
	}

	wins := 1

	for i := 0; i < len(time); i++ {
		possibleWins := 0
		timeAvailable := time[i]
		distanceToBeat := distance[i]

		for j := 0; j < timeAvailable; j++ {
			speed := j
			timeCharging := j
			timeSprinting := timeAvailable - timeCharging
			distanceTravelled := timeSprinting * speed
			if distanceTravelled > distanceToBeat {
				possibleWins++
			}
		}

		wins = wins * possibleWins
	}

	return wins
}
