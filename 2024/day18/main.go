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

	size := 71            // 71
	initialNumber := 1023 // 1023

	grid := make([][]string, size)

	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for i, line := range lines {
		nums := strings.Split(line, ",")
		locI, _ := strconv.Atoi(nums[0])
		locJ, _ := strconv.Atoi(nums[1])

		grid[locJ][locI] = "#"

		if i == initialNumber {
			break
		}
	}
	for i := initialNumber + 1; i < len(lines); i++ {
		nums := strings.Split(lines[i], ",")
		locI, _ := strconv.Atoi(nums[0])
		locJ, _ := strconv.Atoi(nums[1])
		grid[locJ][locI] = "#"

		escaped := false
		queue := make([][]int, 0)
		queue = append(queue, []int{0, 0, 0})
		visited := make(map[string]int)

		for len(queue) > 0 {
			// maybe sort the queue before to get the one with the smallest distance, but wasn't needed to complete the challenge
			loc := queue[0]
			queue = queue[1:]

			if loc[0] == size-1 && loc[1] == size-1 {
				escaped = true
				fmt.Println("Escaped", loc)
				break
			}

			if val, ok := visited[fmt.Sprintf("%d-%d", loc[0], loc[1])]; ok && val <= loc[2] {
				continue
			}
			visited[fmt.Sprintf("%d-%d", loc[0], loc[1])] = loc[2]

			directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
			for _, dir := range directions {
				newI := loc[0] + dir[0]
				newJ := loc[1] + dir[1]

				if newI < 0 || newI >= size || newJ < 0 || newJ >= size {
					continue
				}

				if grid[newI][newJ] != "#" {
					queue = append(queue, []int{newI, newJ, loc[2] + 1})
				}

			}
		}

		if !escaped {
			fmt.Println("not escaped on", lines[i])
			break
		}
	}

	fmt.Println("done")
}
