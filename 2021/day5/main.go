package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

type Wind struct {
	x1, y1, x2, y2 int
}

func main() {
	values := strings.Split(content, "\n")

	maxX, maxY := 0, 0
	winds := make([]Wind, len(values))
	for i, v := range values {
		parts := strings.Split(v, "->")
		firsts := strings.Split(strings.Trim(parts[0], " "), ",")
		seconds := strings.Split(strings.Trim(parts[1], " "), ",")
		x1, _ := strconv.Atoi(firsts[1])
		y1, _ := strconv.Atoi(firsts[0])
		x2, _ := strconv.Atoi(seconds[1])
		y2, _ := strconv.Atoi(seconds[0])

		if y1 > maxY {
			maxY = y1
		}
		if y2 > maxY {
			maxY = y2
		}
		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		winds[i] = Wind{x1, y1, x2, y2}
	}
	fmt.Println(fmt.Sprintf("At least two lines overlap at %d points when not counting diagonals winds.", getNumberOfPointsWhereTwoLinesOverlap(winds, maxX, maxY, false)))
	fmt.Println(fmt.Sprintf("At least two lines overlap at %d points when counting diagonals winds.", getNumberOfPointsWhereTwoLinesOverlap(winds, maxX, maxY, true)))
}

func getNumberOfPointsWhereTwoLinesOverlap(winds []Wind, maxX, maxY int, coundDiagonals bool) int {
	grid := make([][]int, maxY+1)
	for i := range grid {
		line := make([]int, maxX+1)
		grid[i] = line
	}

	for _, wind := range winds {
		x1, x2, y1, y2 := wind.x1, wind.x2, wind.y1, wind.y2
		for {
			if !coundDiagonals && x1 != x2 && y1 != y2 {
				break
			}

			grid[y1][x1] += 1

			if x1 == x2 && y1 == y2 {
				break
			}

			if x1 < x2 {
				x1 += 1
			} else if x1 > x2 {
				x1 += -1
			}

			if y1 < y2 {
				y1 += 1
			} else if y1 > y2 {
				y1 += -1
			}
		}
	}

	count := 0
	for _, line := range grid {
		for _, cell := range line {
			if cell >= 2 {
				count++
			}
		}
	}
	return count
}
