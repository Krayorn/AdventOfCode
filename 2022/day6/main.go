package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var content string

func main() {
	//letters := strings.Split(content, "")

	lastFour := make([]rune, 0, 0)

	for i, letter := range content {
		lastFour = append(lastFour, letter)
		if len(lastFour) == 14 {
			if allDiff(lastFour) {
				fmt.Println(i + 1)
				break
			}

			lastFour = lastFour[1:]
			continue
		}

	}
}

func allDiff(letters []rune) bool {
	s := make(map[rune]bool)

	for _, l := range letters {
		s[l] = true
	}

	if len(s) == len(letters) {
		return true
	}
	return false
}
