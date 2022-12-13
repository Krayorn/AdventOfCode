package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

var draw []string

func main() {
	lines := strings.Split(content, "\n")

	draw = make([]string, 6, 6)

	cycle := 0
	x := 1

	sum := 0
	for _, instruction := range lines {
		if instruction == "noop" {
			cycle += 1
			sum = execCycle(cycle, x, sum)

			continue
		}

		parts := strings.Split(instruction, " ")
		value, _ := strconv.Atoi(parts[1])
		cycle += 1
		sum = execCycle(cycle, x, sum)

		cycle += 1
		sum = execCycle(cycle, x, sum)
		x += value

	}

	for _, line := range draw {
		fmt.Println(line)
	}
}

func execCycle(cycle int, x int, sum int) int {
	saves := []int{20, 60, 100, 140, 180, 220}

	for _, round := range saves {
		if cycle == round {
			sum = sum + x*cycle
			break
		}
	}

	for i, line := range draw {
		if len(line) == 40 {
			continue
		}

		fmt.Println(x, cycle)
		if cycle-(40*i) == x || cycle-(40*i) == x+1 || cycle-(40*i) == x+2 {
			draw[i] = draw[i] + "#"
		} else {
			draw[i] = draw[i] + "."
		}
		fmt.Println(draw[i])

		break
	}

	return sum
}
