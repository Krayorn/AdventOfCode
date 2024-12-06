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
	lines := strings.Split(content, "\n")

	grid := make([][]string, len(lines))

	currentGuardPosition := []int{0, 0}
	initialGuardPosition := []int{0, 0}

	positions := make(map[string]bool)

	for j, line := range lines {
		chars := strings.Split(line, "")
		row := make([]string, len(chars))
		for i, char := range chars {
			row[i] = char
			if char == "^" {
				currentGuardPosition = []int{j, i}
				initialGuardPosition = []int{j, i}
				positions[fmt.Sprintf("%d-%d", j, i)] = true
			}
		}
		grid[j] = row
	}

	directions := map[string][]int{"^": {-1, 0}, ">": {0, 1}, "v": {1, 0}, "<": {0, -1}}
	order := map[string]string{"^": ">", ">": "v", "v": "<", "<": "^"}

	for {
		currentDirection := directions[grid[currentGuardPosition[0]][currentGuardPosition[1]]]
		newI, newJ := currentGuardPosition[0]+currentDirection[0], currentGuardPosition[1]+currentDirection[1]

		if newI >= len(grid) || newI < 0 || newJ >= len(grid[0]) || newJ < 0 {
			grid[currentGuardPosition[0]][currentGuardPosition[1]] = "."
			break
		}

		if grid[newI][newJ] == "#" {
			grid[currentGuardPosition[0]][currentGuardPosition[1]] = order[grid[currentGuardPosition[0]][currentGuardPosition[1]]]
			continue
		}

		positions[fmt.Sprintf("%d-%d", newI, newJ)] = true
		grid[newI][newJ] = grid[currentGuardPosition[0]][currentGuardPosition[1]]
		grid[currentGuardPosition[0]][currentGuardPosition[1]] = "."
		currentGuardPosition = []int{newI, newJ}
	}

	fmt.Println(len(positions))
	sum := 0

	for position := range positions {
		parts := strings.Split(position, "-")
		i, _ := strconv.Atoi(parts[0])
		j, _ := strconv.Atoi(parts[1])

		grid[i][j] = "#"
		grid[initialGuardPosition[0]][initialGuardPosition[1]] = "^"
		currentGuardPosition[0] = initialGuardPosition[0]
		currentGuardPosition[1] = initialGuardPosition[1]

		histo := make(map[string][]string)
		histo[fmt.Sprintf("%d-%d", initialGuardPosition[0], initialGuardPosition[1])] = []string{"^"}

		loop := false
	TURN:
		for {
			currentDirection := directions[grid[currentGuardPosition[0]][currentGuardPosition[1]]]
			newI, newJ := currentGuardPosition[0]+currentDirection[0], currentGuardPosition[1]+currentDirection[1]

			if newI >= len(grid) || newI < 0 || newJ >= len(grid[0]) || newJ < 0 {
				grid[currentGuardPosition[0]][currentGuardPosition[1]] = "."
				break
			}

			if grid[newI][newJ] == "#" {
				grid[currentGuardPosition[0]][currentGuardPosition[1]] = order[grid[currentGuardPosition[0]][currentGuardPosition[1]]]
				continue
			}

			if _, ok := histo[fmt.Sprintf("%d-%d", newI, newJ)]; !ok {
				histo[fmt.Sprintf("%d-%d", newI, newJ)] = []string{}
			}
			allPos := histo[fmt.Sprintf("%d-%d", newI, newJ)]
			for _, pos := range allPos {
				if pos == grid[currentGuardPosition[0]][currentGuardPosition[1]] {
					loop = true
					break TURN
				}
			}
			histo[fmt.Sprintf("%d-%d", newI, newJ)] = append(allPos, grid[currentGuardPosition[0]][currentGuardPosition[1]])

			grid[newI][newJ] = grid[currentGuardPosition[0]][currentGuardPosition[1]]
			grid[currentGuardPosition[0]][currentGuardPosition[1]] = "."
			currentGuardPosition = []int{newI, newJ}
		}

		if loop {
			sum++
		}
		grid[i][j] = "."
	}

	fmt.Println(sum)
}
