package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

//go:embed input.txt
var content string

func main() {
	cards := strings.Split(content, "\n")
	sum := 0

	scratchCards := make(map[int]int)

	re := regexp.MustCompile(`^Card \d+: `)

	for i:=1;i <= len(cards);i++ {
		scratchCards[i] = 1
	}

	//GAME: 
	for i, card := range cards {
		card = strings.Replace(card, "  ", " ", -1)
		card = re.ReplaceAllString(card, "")

		points := 0

		numbers := strings.Split(card, " | ")
		
		winnings := strings.Split(numbers[0], " ")
		winningNumbers := make([]int, len(winnings))
		for i, num := range winnings {
			n, _ := strconv.Atoi(num)
			winningNumbers[i] = n
		}

		playings := strings.Split(numbers[1], " ")
		count := 0
		for _, num := range playings {
			n, _ := strconv.Atoi(num)

			for _, winningNumber := range winningNumbers {
				if n == winningNumber {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
					count++
				}
			}
		}

		for j:=0; j < count; j++ {
			scratchCards[i+2+j] = scratchCards[i+2+j] + 1 * scratchCards[i+1]
		}

		sum += points

	}

	fmt.Println(sum)
	
	fmt.Println(scratchCards)
	sum2 := 0
	for _, count := range scratchCards {
		sum2 += count
	}
	fmt.Println(sum2)

}
