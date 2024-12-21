package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

type Item struct {
	I, J int
}

func main() {
	lines := strings.Split(content, "\n")
	queue := make([]Item, 0)

	grid := make([][]string, len(lines))
	for i, line := range lines {
		row := make([]string, len(line))
		for j, char := range line {
			row[j] = string(char)

			if char == 'S' {
				queue = append(queue, Item{I: i, J: j})
			}
		}
		grid[i] = row
	}

	visited := make(map[string]int)

	for len(queue) > 0 {
		loc := queue[0]
		queue = queue[1:]

		directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
		for _, dir := range directions {
			newI := loc.I + dir[0]
			newJ := loc.J + dir[1]

			if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) {
				continue
			}

			if grid[newI][newJ] != "#" {
				if _, ok := visited[fmt.Sprintf("%d-%d", newI, newJ)]; !ok {
					visited[fmt.Sprintf("%d-%d", newI, newJ)] = visited[fmt.Sprintf("%d-%d", loc.I, loc.J)] + 1
					queue = append(queue, Item{I: newI, J: newJ})
				}
			}
		}
	}

	maxBlink := 20
	timeSkips := make(map[int]int)
	for location, originalTime := range visited {
		for i := -maxBlink; i <= maxBlink; i++ {
			for j := -maxBlink; j <= maxBlink; j++ {
				nums := strings.Split(location, "-")
				currentI, _ := strconv.Atoi(nums[0])
				currentJ, _ := strconv.Atoi(nums[1])
				newI := currentI + i
				newJ := currentJ + j

				skip := math.Abs(float64(i)) + math.Abs(float64(j))
				if skip == 0 || skip > float64(maxBlink) {
					continue
				}

				if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) {
					continue
				}

				if newTime, ok := visited[fmt.Sprintf("%d-%d", newI, newJ)]; ok {
					timeSaved := newTime - originalTime - int(skip)
					if timeSaved > 0 {
						timeSkips[timeSaved]++
					}
				}
			}
		}
	}

	res := 0
	for saved, count := range timeSkips {
		if saved >= 100 {
			res += count
		}
	}

	fmt.Println(res)
}
