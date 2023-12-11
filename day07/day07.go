package day07

import (
	"slices"
	"strings"

	hf "github.com/FMinister/2023-advent-of-code/helperfunctions"
)

type CamelCards struct {
	cards []Card
	bid   int
}

type Card struct {
	name  string
	worth int
}

var winningTypes = map[string][]CamelCards{}

var winningSequence = []string{
	"High card",
	"One pair",
	"Two pair",
	"Three of a kind",
	"Full house",
	"Four of a kind",
	"Five of a kind",
}

func Day07_1() int {
	fileText := hf.OpenAndReadFile("./day07/example.input.txt")

	camelCards := []CamelCards{}
	winners := winningTypes

	for _, line := range fileText {
		items := strings.Split(line, " ")
		camelCards = append(camelCards, CamelCards{cardToCardAndWorth(items[0]), hf.StringToInt(items[1])})
		winningType := sortToWinningType(camelCards[len(camelCards)-1].cards)
		winners[winningType] = append(winners[winningType], CamelCards{cardToCardAndWorth(items[0]), hf.StringToInt(items[1])})
	}

	for _, winner := range winners {
		slices.SortFunc(winner, func(i, j CamelCards) int {
			for k := 0; k < len(i.cards); k++ {
				if i.cards[k].worth > j.cards[k].worth {
					return 1
				}
				if i.cards[k].worth < j.cards[k].worth {
					return -1
				}
			}
			return 0
		})
	}

	sum := 0
	count := 1

	for _, winningType := range winningSequence {
		for _, winner := range winners[winningType] {
			sum = sum + winner.bid*count
			count++
		}
	}

	return sum
}

func cardToCardAndWorth(card string) []Card {
	cards := []Card{}
	for _, c := range card {
		switch c {
		case 'A':
			cards = append(cards, Card{"A", 14})
		case 'K':
			cards = append(cards, Card{"K", 13})
		case 'Q':
			cards = append(cards, Card{"Q", 12})
		case 'J':
			cards = append(cards, Card{"J", 11})
		case 'T':
			cards = append(cards, Card{"T", 10})
		default:
			cards = append(cards, Card{string(c), hf.StringToInt(string(c))})
		}
	}

	return cards
}

func sortToWinningType(cards []Card) string {
	cardCount := make([]int, 15)

	for _, card := range cards {
		cardCount[card.worth]++
	}

	fullHouseOption := false
	twoPairsOption := false

	for _, card := range cardCount {
		if card == 5 {
			return "Five of a kind"
		}
		if card == 4 {
			return "Four of a kind"
		}
		if card == 2 && fullHouseOption {
			return "Full house"
		}
		if card == 3 {
			fullHouseOption = true
		}
		if card == 3 && twoPairsOption {
			return "Full house"
		}
		if card == 2 && twoPairsOption {
			return "Two pair"
		}
		if card == 2 {
			twoPairsOption = true
		}
	}

	if fullHouseOption {
		return "Three of a kind"
	}
	if twoPairsOption {
		return "One pair"
	}

	return "High card"
}
