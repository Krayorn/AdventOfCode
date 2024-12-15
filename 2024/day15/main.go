package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	lines := strings.Split(content, "\n\n")

	gridRaw := strings.Split(lines[0], "\n")
	grid := make([][]string, len(gridRaw))

	currentPosition := []int{}

	for i, line := range gridRaw {
		chars := strings.Split(line, "")
		row := make([]string, len(chars)*2)
		for j, char := range chars {
			if char == "." || char == "#" {
				row[j*2] = char
				row[j*2+1] = char
			}
			if char == "O" {
				row[j*2] = "["
				row[j*2+1] = "]"
			}
			if char == "@" {
				currentPosition = []int{i, j * 2}
				row[j*2] = char
				row[j*2+1] = "."
			}
		}
		grid[i] = row
	}

	moves := strings.Split(strings.ReplaceAll(lines[1], "\n", ""), "")

	instructs := map[string][]int{"^": {-1, 0}, "<": {0, -1}, ">": {0, 1}, "v": {1, 0}}

	printGrid(grid)

	for _, move := range moves {
		dir := instructs[move]

		newGrid, moved := applyMove(currentPosition, dir, grid)
		grid = newGrid
		if moved {
			currentPosition[0] = currentPosition[0] + dir[0]
			currentPosition[1] = currentPosition[1] + dir[1]
		}

		printGrid(grid)
	}

	sum := 0
	for i, row := range grid {
		for j, char := range row {
			if char == "[" {
				sum += i*100 + j
			}
		}
	}
	fmt.Println(sum)
}

func applyMove(currentPosition, dir []int, original [][]string) ([][]string, bool) {
	grid := make([][]string, len(original))
	for i := range original {
		grid[i] = make([]string, len(original[i]))
		copy(grid[i], original[i])
	}

	newX, newY := currentPosition[0]+dir[0], currentPosition[1]+dir[1]
	if grid[newX][newY] == "." {
		grid[newX][newY] = grid[currentPosition[0]][currentPosition[1]]
		grid[currentPosition[0]][currentPosition[1]] = "."
		return grid, true
	}

	if grid[newX][newY] == "#" {
		return grid, false
	}

	if grid[newX][newY] == "[" || grid[newX][newY] == "]" {
		if dir[1] == 0 {
			val := grid[newX][newY]
			newGrid, moved := applyMove([]int{newX, newY}, dir, grid)
			var newGrid2 [][]string
			var moved2 bool
			if val == "[" {
				newGrid2, moved2 = applyMove([]int{newX, newY + 1}, dir, newGrid)
			} else {
				newGrid2, moved2 = applyMove([]int{newX, newY - 1}, dir, newGrid)
			}
			if !moved || !moved2 {
				return grid, false
			}

			newGrid2[newX][newY] = newGrid2[currentPosition[0]][currentPosition[1]]
			newGrid2[currentPosition[0]][currentPosition[1]] = "."
			return newGrid2, true

		}
		newGrid, moved := applyMove([]int{newX, newY}, dir, grid)
		if !moved {
			return grid, false
		}
		newGrid[newX][newY] = newGrid[currentPosition[0]][currentPosition[1]]
		newGrid[currentPosition[0]][currentPosition[1]] = "."
		return newGrid, true
	}

	return grid, false
}

func printGrid(grid [][]string) {
	for _, line := range grid {
		fmt.Println(line)
	}

	fmt.Println("-----------")
}
