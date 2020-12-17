package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	maxLength := 12 + len(values)

	grid := make([][][]int, 13)

	for i := range grid {
		z := make([][]int, maxLength)
		for j := range z {
			x := make([]int, maxLength)
			z[j] = x
		}
		grid[i] = z
	}

	for i, v := range values {
		for j, c := range v {
			if c == '#' {
				grid[6][i+6][j+6] = 1
			}
		}
	}

	pprint(grid)
	for cycle := 0; cycle < 6; cycle++ {
		newGrid := make([][][]int, 13)
		for z, plan := range grid {
			newGridRow := make([][]int, maxLength)
			for x, row := range plan {
				newGridColumns := make([]int, maxLength)
				for y, cube := range row {
					count := countAround(z, x, y, grid)
					if cube == 0 && count == 3 {
						newGridColumns[y] = 1
					} else if cube == 1 && (count == 2 || count == 3) {
						newGridColumns[y] = 1
					} else {
						newGridColumns[y] = 0
					}
				}
				newGridRow[x] = newGridColumns
			}
			newGrid[z] = newGridRow
		}

		for z, plan := range grid {
			for x, row := range plan {
				for y := range row {
					grid[z][x][y] = newGrid[z][x][y]
				}
			}
		}
		pprint(grid)
	}

	total := 0
	for z, plan := range grid {
		for x, row := range plan {
			for y := range row {
				if grid[z][x][y] == 1 {
					total++
				}
			}
		}
	}
	fmt.Println(total)

}

func countAround(z int, x int, y int, grid [][][]int) int {
	count := 0

	var startPlan, startRow, startColumn, endPlan, endRow, endColumn int
	if z == 0 {
		startPlan = 0
	} else {
		startPlan = z - 1
	}
	if x == 0 {
		startRow = 0
	} else {
		startRow = x - 1
	}
	if y == 0 {
		startColumn = 0
	} else {
		startColumn = y - 1
	}
	if z+1 == len(grid) {
		endPlan = z
	} else {
		endPlan = z + 1
	}
	if x+1 == len(grid[0]) {
		endRow = x
	} else {
		endRow = x + 1
	}
	if y+1 == len(grid[0][0]) {
		endColumn = y
	} else {
		endColumn = y + 1
	}

	for i := startPlan; i <= endPlan; i++ {
		for j := startRow; j <= endRow; j++ {
			for k := startColumn; k <= endColumn; k++ {
				if z == i && x == j && y == k {
					continue
				}
				if grid[i][j][k] == 1 {
					count++
				}
			}
		}
	}
	return count
}

func pprint(grid [][][]int) {
	fmt.Println("-----------------------------")
	for i, plan := range grid {
		fmt.Println("z =", i-6)
		for _, row := range plan {
			for _, cube := range row {
				fmt.Print(cube)
			}
			fmt.Println("")
		}
	}
	fmt.Println("-----------------------------")
}
