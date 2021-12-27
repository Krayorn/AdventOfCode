package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, "\n")

	incompleteScores := make([]int, 0)
	score := 0

	for _, v := range values {
		last := make([]rune, 0)
		lastScore := score
		for _, c := range v {
			if c == '{' || c == '(' || c == '<' || c == '[' {
				last = append(last, c)
				continue
			}
			if c == '}' && (len(last) == 0 || last[len(last)-1] != '{') {
				score += 1197
				break
			}
			if c == ')' && (len(last) == 0 || last[len(last)-1] != '(') {
				score += 3
				break
			}
			if c == '>' && (len(last) == 0 || last[len(last)-1] != '<') {
				score += 25137
				break
			}
			if c == ']' && (len(last) == 0 || last[len(last)-1] != '[') {
				score += 57
				break
			}
			last = remove(last, len(last)-1)
		}
		if score == lastScore {
			lineScore := 0
			for i := len(last) - 1; i >= 0; i-- {
				lineScore = lineScore * 5

				c := last[i]
				if c == '{' {
					lineScore += 3
				}
				if c == '(' {
					lineScore += 1
				}
				if c == '<' {
					lineScore += 4
				}
				if c == '[' {
					lineScore += 2
				}
			}
			incompleteScores = append(incompleteScores, lineScore)
		}

		lastScore = score
	}

	sort.Slice(incompleteScores, func(i, j int) bool {
		return incompleteScores[i] > incompleteScores[j]
	})

	fmt.Println(fmt.Sprintf("The total syntax error score is %d.", score))
	fmt.Println(fmt.Sprintf("The middle score of the incomplete lines is %d.", incompleteScores[(len(incompleteScores)-1)/2]))
}

func remove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}
