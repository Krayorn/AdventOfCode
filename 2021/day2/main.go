package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, "\n")

	strat1horizontal, strat1depth := 0, 0
	horizontal, depth, aim := 0, 0, 0

	for _, v := range values {
		data := strings.Split(v, " ")
		number, _ := strconv.Atoi(data[1])

		if data[0] == "forward" {
			strat1horizontal += number
			horizontal += number
			depth += aim * number
		}
		if data[0] == "down" {
			strat1depth += number
			aim += number
		}
		if data[0] == "up" {
			strat1depth -= number
			aim -= number
		}
	}

	fmt.Println(fmt.Sprintf("The product of the final depth and horizontal position with the first strategy is %d ", strat1depth*strat1horizontal))
	fmt.Println(fmt.Sprintf("The product of the final depth and horizontal position with the second strategy is %d ", depth*horizontal))
}
