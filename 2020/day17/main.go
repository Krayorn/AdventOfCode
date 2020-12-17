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

	grid := make([][][][]int, 13)

	for h := range grid {
		z := make([][][]int, 13)
		for i := range z {
			x := make([][]int, maxLength)
			for j := range x {
				y := make([]int, maxLength)
				x[j] = y
			}
			z[i] = x
		}
		grid[h] = z
	}

	for i, v := range values {
		for j, c := range v {
			if c == '#' {
				grid[6][6][i+6][j+6] = 1
			}
		}
	}

	// pprint(grid)
	for cycle := 0; cycle < 6; cycle++ {
		newGrid := make([][][][]int, 13)
		for w, wPlan := range grid {
			newGridPlans := make([][][]int, 13)
			for z, plan := range wPlan {
				newGridRow := make([][]int, maxLength)
				for x, row := range plan {
					newGridColumns := make([]int, maxLength)
					for y, cube := range row {
						count := countAround(w, z, x, y, grid)
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
				newGridPlans[z] = newGridRow
			}
			newGrid[w] = newGridPlans
		}

		for z, plan := range grid {
			for x, row := range plan {
				for y := range row {
					grid[z][x][y] = newGrid[z][x][y]
				}
			}
		}
		// pprint(grid)
	}

	total := 0
	for w, wPlan := range grid {
		for z, plan := range wPlan {
			for x, row := range plan {
				for y := range row {
					if grid[w][z][x][y] == 1 {
						total++
					}
				}
			}
		}
	}
	fmt.Println(total)

}

func countAround(w int, z int, x int, y int, grid [][][][]int) int {
	count := 0

	var startWPlan, startPlan, startRow, startColumn, endWPlan, endPlan, endRow, endColumn int
	if w == 0 {
		startWPlan = 0
	} else {
		startWPlan = w - 1
	}
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
	if w+1 == len(grid) {
		endWPlan = w
	} else {
		endWPlan = w + 1
	}
	if z+1 == len(grid[0]) {
		endPlan = z
	} else {
		endPlan = z + 1
	}
	if x+1 == len(grid[0][0]) {
		endRow = x
	} else {
		endRow = x + 1
	}
	if y+1 == len(grid[0][0][0]) {
		endColumn = y
	} else {
		endColumn = y + 1
	}

	for h := startWPlan; h <= endWPlan; h++ {
		for i := startPlan; i <= endPlan; i++ {
			for j := startRow; j <= endRow; j++ {
				for k := startColumn; k <= endColumn; k++ {
					if w == h && z == i && x == j && y == k {
						continue
					}
					if grid[h][i][j][k] == 1 {
						count++
					}
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
