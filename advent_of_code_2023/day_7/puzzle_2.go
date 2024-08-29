package day_7

import (
	steveutil "advent_of_code_2023/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

var CardPower = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var CardPowerJoker = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

var HandPower = map[string]int{
	"Five_of_a_kind":  7,
	"Four_of_a_kind":  6,
	"Full_house":      5,
	"Three_of_a_kind": 4,
	"Two_pair":        3,
	"Pair":            2,
	"High_card":       1,
}

type Hand struct {
	cards []rune
	bid   int
	power string
}

func Puzzle_2() {
	f, err := os.Open(path.Join("day_7", "inputs", "puzzle.txt"))
	steveutil.Check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	hands := []Hand{}

	jokerInPlay := true

	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		hand := []rune(fields[0])
		bid, err := strconv.Atoi(fields[1])
		steveutil.Check(err)
		power := DetermineHandPower(hand, jokerInPlay)
		hands = append(hands, Hand{cards: hand, bid: bid, power: power})
	}

	sort.Slice(hands, func(i, j int) bool {
		if HandPower[hands[i].power] == HandPower[hands[j].power] {
			return DetermineHigherCard(hands[i].cards, hands[j].cards, jokerInPlay)
		} else {
			return HandPower[hands[i].power] > HandPower[hands[j].power]
		}
	})

	totalWinnings := 0
	rankMultipler := 1
	for i := len(hands) - 1; i >= 0; i-- {
		totalWinnings += (hands[i].bid * rankMultipler)
		rankMultipler++
		fmt.Printf("Hand: %s, Power: %s, Bid: %d\n", string(hands[i].cards), hands[i].power, hands[i].bid)
	}

	fmt.Println(totalWinnings)
}

func DetermineHigherCard(hand1, hand2 []rune, jokerInPlay bool) bool {
	hand1Higher := false

	cardPower := CardPower
	if jokerInPlay {
		cardPower = CardPowerJoker
	}

	for i := 0; i < len(hand1); i++ {
		if cardPower[hand1[i]] > cardPower[hand2[i]] {
			hand1Higher = true
			break
		} else if cardPower[hand1[i]] < cardPower[hand2[i]] {
			hand1Higher = false
			break
		} else {
			continue
		}
	}
	return hand1Higher
}

func DetermineHandPower(cards []rune, useJokers bool) string {
	cardCounter := map[rune]int{}
	cardPairs := map[rune]string{}
	pairs := []string{}
	handPower := "High_card"
	jokerCount := 0

	for _, card := range cards {

		if !useJokers || card != 'J' {
			cardCounter[card]++
			handPower := "High_card"

			if cardCounter[card] == 5 {
				handPower = "Five_of_a_kind"
			} else if cardCounter[card] == 4 {
				handPower = "Four_of_a_kind"
			} else if cardCounter[card] == 3 {
				handPower = "Three_of_a_kind"
			} else if cardCounter[card] == 2 {
				handPower = "Pair"
			}

			if handPower != "High_card" {
				cardPairs[card] = handPower
			}
		} else {
			jokerCount++
		}
	}

	for _, pair := range cardPairs {
		pairs = append(pairs, pair)
	}

	if len(pairs) == 1 {
		handPower = pairs[0]
	} else if len(pairs) > 1 {
		if pairs[0] == "Three_of_a_kind" || pairs[1] == "Three_of_a_kind" {
			handPower = "Full_house"
		} else if pairs[0] == "Pair" || pairs[1] == "Pair" {
			handPower = "Two_pair"
		} else {
			panic("Somehow calculated non pair")
		}
	}

	if useJokers && jokerCount > 0 {
		handPower = calculateJokers(handPower, jokerCount)
	}

	return handPower
}

func calculateJokers(handPower string, jokerCount int) string {

	if jokerCount == 0 {
		return handPower
	}

	if handPower == "Five_of_a_kind" || jokerCount == 5 {
		return "Five_of_a_kind"
	}

	switch handPower {
	case "Four_of_a_kind":
		return "Five_of_a_kind"
	case "Three_of_a_kind":
		if jokerCount == 1 {
			return "Four_of_a_kind"
		} else if jokerCount == 2 {
			return "Five_of_a_kind"
		}
	case "Two_pair":
		return "Full_house"
	case "Pair":
		if jokerCount == 1 {
			return "Three_of_a_kind"
		} else if jokerCount == 2 {
			return "Four_of_a_kind"
		} else if jokerCount == 3 {
			return "Five_of_a_kind"
		}
	default:
		if jokerCount == 1 {
			return "Pair"
		} else if jokerCount == 2 {
			return "Three_of_a_kind"
		} else if jokerCount == 3 {
			return "Four_of_a_kind"
		} else {
			return "Five_of_a_kind"
		}
	}

	return handPower
}
