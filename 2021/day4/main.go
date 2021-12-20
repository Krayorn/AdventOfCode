package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	grid [][]int
	Won  bool
}

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, "\n")

	numbers := make([]int, len(strings.Split(values[0], ",")))
	for i, v := range strings.Split(values[0], ",") {
		value, _ := strconv.Atoi(v)
		numbers[i] = value
	}

	boards := make([]Board, 0)
	for i := 2; i < len(values); i += 6 {
		lines := make([][]int, 0)
		for j := 0; j < 5; j++ {
			line := make([]int, 5)
			elem := strings.Split(strings.ReplaceAll(strings.Trim(values[i+j], " "), "  ", " "), " ")
			for k, n := range elem {
				value, _ := strconv.Atoi(n)
				line[k] = value
			}
			lines = append(lines, line)
		}
		boards = append(boards, Board{grid: lines, Won: false})
	}
	firstBoardScore, lastBoardScore := -1, -1

	drawn := make(map[int]bool)
	for _, n := range numbers {
		drawn[n] = true
		for i, board := range boards {
			if board.Won == false && board.Win(drawn) {
				score := board.Score(drawn, n)
				if firstBoardScore == -1 {
					firstBoardScore = score
				}
				lastBoardScore = score
				boards[i].Won = true
			}
		}
	}

	fmt.Println(fmt.Sprintf("The score of the first board to win will be: %d.", firstBoardScore))
	fmt.Println(fmt.Sprintf("The score of the last board to win will be: %d.", lastBoardScore))
}

func (board Board) Win(drawn map[int]bool) bool {
	for _, line := range board.grid {
		count := 0
		for _, elem := range line {
			if _, ok := drawn[elem]; !ok {
				break
			}
			count++
		}
		if count == 5 {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		count := 0
		for j := 0; j < 5; j++ {
			if _, ok := drawn[board.grid[j][i]]; !ok {
				break
			}
			count++
		}
		if count == 5 {
			return true
		}
	}

	return false
}

func (board Board) Score(drawn map[int]bool, n int) int {
	sum := 0
	for _, line := range board.grid {
		for _, elem := range line {
			if _, ok := drawn[elem]; !ok {
				sum += elem
			}
		}
	}

	return sum * n
}
