package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	lines := strings.Split(content, "\n")

	grid := make([][]string, len(lines))

	for j, line := range lines {
		chars := strings.Split(line, "")
		row := make([]string, len(chars))
		for i, char := range chars {
			row[i] = char
		}
		grid[j] = row
	}

	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	sum := 0

	for {
		initialSum := sum
		coords := make([][]int, 0)
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == "@" {
					count := 0
					for _, dir := range directions {
						x := i + dir[1]
						y := j + dir[0]
						if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[i]) {
							if grid[x][y] == "@" {
								count++
							}
						}
					}
					if count < 4 {
						sum++
						coords = append(coords, []int{i, j})
					}
				}
			}
		}
		if sum == initialSum {
			break
		}
		for _, coord := range coords {
			grid[coord[0]][coord[1]] = "."
		}
	}

	fmt.Println(sum)
}
