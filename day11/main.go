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

	grid := make([][]string, len(values))
	newGrid := make([][]string, len(values))

	for i, v := range values {
		row := make([]string, len(v))
		newRow := make([]string, len(v))
		newGrid[i] = newRow
		for j, letter := range v {
			row[j] = string(letter)
		}
		grid[i] = row
	}

	fmt.Println("Number of seat that will be occupied if people move once 4 seats around them are occupied =>", countNumberOfSeat(grid, newGrid, true, 4))
	for i, v := range values {
		row := make([]string, len(v))
		newRow := make([]string, len(v))
		newGrid[i] = newRow
		for j, letter := range v {
			row[j] = string(letter)
		}
		grid[i] = row
	}
	fmt.Println("Number of seat that will be occupied if people move once they can see 5 occupied seats in any direction =>", countNumberOfSeat(grid, newGrid, false, 5))
}

func countNumberOfSeat(grid [][]string, newGrid [][]string, countAround bool, treshold int) int {
	for {
		update := false
		for i, row := range grid {
			for j, letter := range row {
				if letter == "L" || letter == "#" {
					occupiedSeatCount := 0
					startRow, startColumn, endRow, endColumn := 0, 0, 0, 0
					if i == 0 {
						startRow = 0
					} else {
						startRow = i - 1
					}
					if j == 0 {
						startColumn = 0
					} else {
						startColumn = j - 1
					}
					if i+1 == len(grid) {
						endRow = i
					} else {
						endRow = i + 1
					}
					if j+1 == len(row) {
						endColumn = j
					} else {
						endColumn = j + 1
					}
					for k := startRow; k <= endRow; k++ {
						for l := startColumn; l <= endColumn; l++ {
							if k == i && l == j {
								continue
							}
							if countAround {
								if grid[k][l] == "#" {
									occupiedSeatCount++
								}
							} else {
								offsetK, offsetL := k, l
								for {
									if grid[offsetK][offsetL] == "#" {
										occupiedSeatCount++
										break
									}
									if grid[offsetK][offsetL] == "L" {
										break
									}
									if k < i {
										offsetK--
									}
									if k > i {
										offsetK++
									}
									if l < j {
										offsetL--
									}
									if l > j {
										offsetL++
									}
									if offsetK == -1 || offsetL == -1 || offsetK == len(grid) || offsetL == len(row) {
										break
									}
								}
							}
						}
					}
					if letter == "L" {
						if occupiedSeatCount == 0 {
							update = true
							newGrid[i][j] = "#"
						}
					} else if letter == "#" {
						if occupiedSeatCount >= treshold {
							update = true
							newGrid[i][j] = "L"
						}
					}
				} else {
					newGrid[i][j] = grid[i][j]
				}
			}
		}

		for i := range newGrid {
			for j := range newGrid[i] {
				grid[i][j] = newGrid[i][j]
			}
		}

		if !update {
			break
		}
	}
	count := 0
	for _, row := range grid {
		for _, seat := range row {
			if seat == "#" {
				count++
			}
		}
	}

	return count
}

func pprint(grid [][]string) {
	fmt.Println("---------------")
	for _, row := range grid {
		for _, seat := range row {
			fmt.Print(seat)
		}
		fmt.Println("")
	}
}
