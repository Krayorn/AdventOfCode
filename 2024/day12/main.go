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

	areas := make([][][]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "." {
				continue
			}
			area := make([][]int, 0)
			grid, area = searchAround(grid[i][j], i, j, grid, area)
			areas = append(areas, area)
		}
	}

	sum := 0
	sum2 := 0
	for _, area := range areas {
		perimeter := 0
		corners := 0
		fmt.Println(area)
		for i, plot := range area {
			perimeter += 4
			up, left, right, down, up_left, up_right, down_left, down_right := false, false, false, false, false, false, false, false
			for j, plotCompare := range area {
				if i == j {
					continue
				}

				if plot[0]+1 == plotCompare[0] && plot[1] == plotCompare[1] {
					down = true
				}
				if plot[0]-1 == plotCompare[0] && plot[1] == plotCompare[1] {
					up = true
				}
				if plot[0] == plotCompare[0] && plot[1]-1 == plotCompare[1] {
					left = true
				}
				if plot[0] == plotCompare[0] && plot[1]+1 == plotCompare[1] {
					right = true
				}

				if plot[0]+1 == plotCompare[0] && plot[1]+1 == plotCompare[1] {
					down_right = true
				}
				if plot[0]-1 == plotCompare[0] && plot[1]+1 == plotCompare[1] {
					up_right = true
				}
				if plot[0]-1 == plotCompare[0] && plot[1]-1 == plotCompare[1] {
					up_left = true
				}
				if plot[0]+1 == plotCompare[0] && plot[1]-1 == plotCompare[1] {
					down_left = true
				}

				if plotCompare[0] == plot[0] && (plotCompare[1]-plot[1] == 1 || plotCompare[1]-plot[1] == -1) {
					perimeter--
				}

				if plotCompare[1] == plot[1] && (plotCompare[0]-plot[0] == 1 || plotCompare[0]-plot[0] == -1) {
					perimeter--
				}
			}

			if up {
				if left && !up_left {
					corners++
				}
				if right && !up_right {
					corners++
				}
			} else {
				if !left {
					corners++
				}
				if !right {
					corners++
				}
			}

			if down {
				if left && !down_left {
					corners++
				}
				if right && !down_right {
					corners++
				}
			} else {
				if !left {
					corners++
				}
				if !right {
					corners++
				}
			}
			fmt.Println("left", left, "right", right, "up", up, "down", down, "up_left", up_left, "up_right", up_right, "down_left", down_left, "down_right", down_right, corners)
		}

		sum += len(area) * perimeter
		sum2 += len(area) * corners
	}

	fmt.Println(sum, sum2)
}

func searchAround(letter string, i, j int, grid [][]string, area [][]int) ([][]string, [][]int) {
	grid[i][j] = "."
	area = append(area, []int{i, j})

	if i > 0 && grid[i-1][j] == letter {
		grid, area = searchAround(letter, i-1, j, grid, area)
	}

	if i < len(grid)-1 && grid[i+1][j] == letter {
		grid, area = searchAround(letter, i+1, j, grid, area)
	}

	if j < len(grid[i])-1 && grid[i][j+1] == letter {
		grid, area = searchAround(letter, i, j+1, grid, area)
	}

	if j > 0 && grid[i][j-1] == letter {
		grid, area = searchAround(letter, i, j-1, grid, area)
	}

	return grid, area
}
