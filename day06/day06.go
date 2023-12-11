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

func Day06_2() int {
	fileText := hf.OpenAndReadFile("./day06/input.txt")

	time := 0
	distance := 0
	re := regexp.MustCompile(`(\d+)`)

	for i := 0; i < len(fileText); i++ {
		line := fileText[i]

		timeOrDist := re.FindAllString(line, -1)

		timeOrDistString := ""

		for _, item := range timeOrDist {
			timeOrDistString += item
		}

		if time == 0 {
			time = hf.StringToInt(timeOrDistString)
		} else {
			distance = hf.StringToInt(timeOrDistString)
		}
	}

	minWin := 0
	maxWin := 0

	for i := 0; i < time; i++ {
		speed := i
		timeCharging := i
		timeSprinting := time - timeCharging
		distanceTravelled := timeSprinting * speed
		if distanceTravelled > distance {
			minWin = i
			break
		}
	}

	for i := time; i > 0; i-- {
		speed := i
		timeCharging := i
		timeSprinting := time - timeCharging
		distanceTravelled := timeSprinting * speed
		if distanceTravelled > distance {
			maxWin = i
			break
		}
	}

	wins := maxWin - minWin + 1

	return wins
}
