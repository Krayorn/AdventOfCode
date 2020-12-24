package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type camera struct {
	ID        int
	sides     map[string]string
	image     []string
	neighbors map[int]bool
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n\n")

	p1Deck := make([]int, len(strings.Split(values[0], "\n")[1:]))
	p2Deck := make([]int, len(strings.Split(values[0], "\n")[1:]))

	for i, c := range strings.Split(values[0], "\n")[1:] {
		card, _ := strconv.Atoi(c)
		p1Deck[i] = card
	}

	for i, c := range strings.Split(values[1], "\n")[1:] {
		card, _ := strconv.Atoi(c)
		p2Deck[i] = card
	}

	dst1 := make([]int, len(p1Deck))
	dst2 := make([]int, len(p2Deck))
	copy(dst1, p1Deck)
	copy(dst2, p2Deck)

	winningDeckPart1 := playRegularGame(dst1, dst2)
	score1 := 0
	for i, v := range winningDeckPart1 {
		score1 += (len(winningDeckPart1) - i) * v
	}

	_, winningDeckPart2 := playRecursiveGame(p1Deck, p2Deck)

	score2 := 0
	for i, v := range winningDeckPart2 {
		score2 += (len(winningDeckPart2) - i) * v
	}

	fmt.Println(score1, score2)
}

func playRecursiveGame(p1Deck []int, p2Deck []int) (int, []int) {
	memory := make(map[string]bool)

	for len(p1Deck) != 0 && len(p2Deck) != 0 {
		p1String := make([]string, len(p1Deck))
		p2String := make([]string, len(p2Deck))
		for i, n := range p1Deck {
			text := strconv.Itoa(n)
			p1String[i] = text
		}
		for i, n := range p2Deck {
			text := strconv.Itoa(n)
			p2String[i] = text
		}
		addr := strings.Join(p1String, ",") + "-" + strings.Join(p2String, ",")
		if _, ok := memory[addr]; ok {
			return 1, p1Deck
		}
		memory[addr] = true

		p1Card := p1Deck[0]
		p2Card := p2Deck[0]

		p1Deck = remove(p1Deck, 0)
		p2Deck = remove(p2Deck, 0)

		var winner int
		if len(p1Deck) >= p1Card && len(p2Deck) >= p2Card {
			dst1 := make([]int, p1Card)
			dst2 := make([]int, p2Card)
			copy(dst1, p1Deck[:p1Card])
			copy(dst2, p2Deck[:p2Card])
			winner, _ = playRecursiveGame(dst1, dst2)
		} else {
			if p1Card > p2Card {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			p1Deck = append(p1Deck, p1Card, p2Card)
		} else {
			p2Deck = append(p2Deck, p2Card, p1Card)
		}
	}

	if len(p1Deck) > len(p2Deck) {
		return 1, p1Deck
	}

	return 2, p2Deck
}

func playRegularGame(p1Deck []int, p2Deck []int) []int {
	for len(p1Deck) != 0 && len(p2Deck) != 0 {
		if p1Deck[0] > p2Deck[0] {
			p1Deck = append(p1Deck, p1Deck[0], p2Deck[0])
		} else {
			p2Deck = append(p2Deck, p2Deck[0], p1Deck[0])
		}

		p1Deck = remove(p1Deck, 0)
		p2Deck = remove(p2Deck, 0)
	}

	if len(p1Deck) > len(p2Deck) {
		return p1Deck
	}
	return p2Deck
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
