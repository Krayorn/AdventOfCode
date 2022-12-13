package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var content string

var shortcuts map[string]int = map[string]int{
	"A": 0,
	"X": 0,
	"B": 1,
	"Y": 1,
	"C": 2,
	"Z": 2,
}

func main() {
	values := strings.Split(content, "\n")

	score := 0
	score2 := 0
	for _, line := range values {
		val := strings.Split(line, " ")

		e := shortcuts[val[0]]
		m := shortcuts[val[1]]

		score += m + 1
		outcome := e%3 - m%3
		if outcome == -1 || outcome == 2 {
			score += 6
		}
		if outcome == 0 {
			score += 3
		}
		if outcome == 1 {
			score += 0
		}

		if m == 0 {
			score2 += (e+2)%3 + 1
		} else if m == 1 {
			score2 += e + 1
		} else {
			score2 += (e+1)%3 + 1
		}
		score2 += m * 3
	}

	fmt.Println(score)
	fmt.Println(score2)

	score = 0
	score2 = 0
	for _, line := range values {
		val := strings.Split(line, " ")
		enemy := val[0]
		me := val[1]

		if me == "X" {
			score += 1
		}
		if me == "Y" {
			score += 2
		}
		if me == "Z" {
			score += 3
		}

		if enemy == "A" {
			if me == "X" {
				score += 3
			}
			if me == "Y" {
				score += 6
			}
			if me == "Z" {
				score += 0
			}
		}
		if enemy == "B" {
			if me == "X" {
				score += 0
			}
			if me == "Y" {
				score += 3
			}
			if me == "Z" {
				score += 6
			}
		}
		if enemy == "C" {
			if me == "X" {
				score += 6
			}
			if me == "Y" {
				score += 0
			}
			if me == "Z" {
				score += 3
			}
		}

		if me == "X" {
			score2 += 0
			if enemy == "A" {
				score2 += 3
			}
			if enemy == "B" {
				score2 += 1
			}
			if enemy == "C" {
				score2 += 2
			}
		}
		if me == "Y" {
			score2 += 3
			if enemy == "A" {
				score2 += 1
			}
			if enemy == "B" {
				score2 += 2
			}
			if enemy == "C" {
				score2 += 3
			}
		}
		if me == "Z" {
			score2 += 6
			if enemy == "A" {
				score2 += 2
			}
			if enemy == "B" {
				score2 += 3
			}
			if enemy == "C" {
				score2 += 1
			}
		}
	}

	fmt.Println(score)
	fmt.Println(score2)
}
