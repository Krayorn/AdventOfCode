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

	grid := make([][]int, len(lines))

	for j, line := range lines {
		chars := strings.Split(line, "")
		row := make([]int, len(chars))
		for i, char := range chars {
			n, _ := strconv.Atoi(char)
			row[i] = n
		}
		grid[j] = row
	}

	sum := 0
	sum2 := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				uniq := make(map[string]bool)
				results := launchSearch(grid, i, j, 0)
				for _, result := range results {
					uniq[fmt.Sprintf("%d-%d", result[0], result[1])] = true
				}
				sum += len(uniq)
				sum2 += len(results)
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func launchSearch(grid [][]int, i, j, target int) [][]int {
	total := make([][]int, 0)
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) {
		return total
	}

	if grid[i][j] == 9 && target == 9 {
		return [][]int{{i, j}}
	}

	if grid[i][j] == target {
		total = append(total, launchSearch(grid, i+1, j, target+1)...)
		total = append(total, launchSearch(grid, i, j+1, target+1)...)
		total = append(total, launchSearch(grid, i-1, j, target+1)...)
		total = append(total, launchSearch(grid, i, j-1, target+1)...)
	}

	return total
}
