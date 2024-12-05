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

	sumXmas := 0
	sumMas := 0

	following := []string{"X", "M", "A", "S"}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	followingMas := []string{"M", "A", "S"}
	directionsMas := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	aLocation := make(map[string]int)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "X" {

			DIRECTION:
				for _, direction := range directions {
					for letterIdx := 1; letterIdx <= 3; letterIdx++ {
						if direction[0]*letterIdx+i < 0 || direction[0]*letterIdx+i >= len(grid) {
							continue DIRECTION
						}
						if direction[1]*letterIdx+j < 0 || direction[1]*letterIdx+j >= len(grid[i]) {
							continue DIRECTION
						}

						if grid[direction[0]*letterIdx+i][direction[1]*letterIdx+j] != following[letterIdx] {
							continue DIRECTION
						}

						continue
					}
					sumXmas++
				}
			}

			if grid[i][j] == "M" {

			DIRECTIONM:
				for _, direction := range directionsMas {
					theA := ""
					for letterIdx := 1; letterIdx <= 2; letterIdx++ {
						if direction[0]*letterIdx+i < 0 || direction[0]*letterIdx+i >= len(grid) {
							continue DIRECTIONM
						}
						if direction[1]*letterIdx+j < 0 || direction[1]*letterIdx+j >= len(grid[i]) {
							continue DIRECTIONM
						}

						if grid[direction[0]*letterIdx+i][direction[1]*letterIdx+j] != followingMas[letterIdx] {
							continue DIRECTIONM
						}

						if followingMas[letterIdx] == "A" {
							theA = fmt.Sprintf("%d-%d", direction[0]*letterIdx+i, direction[1]*letterIdx+j)
						}

						continue
					}
					if _, ok := aLocation[theA]; !ok {
						aLocation[theA] = 0
					}

					aLocation[theA] = aLocation[theA] + 1
				}
			}
		}
	}

	for _, val := range aLocation {
		if val == 2 {
			sumMas++
		}
	}

	fmt.Println(sumXmas)
	fmt.Println(sumMas)

}
