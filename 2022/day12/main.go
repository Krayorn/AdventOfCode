package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var content string

type Coord struct {
	y, x int
}

var distances map[Coord]int

func main() {
	lines := strings.Split(content, "\n")

	myMap := make([][]int, len(lines), len(lines))
	distances = make(map[Coord]int)

	allStarts := make([]Coord, 0, 0)

	var start, end Coord

	for y, line := range lines {
		myLine := make([]int, len(line), len(line))
		for x, letter := range line {
			n := int(strings.ToLower(string(letter))[0]) - 96
			if letter == 'S' {
				start = Coord{y, x}
				allStarts = append(allStarts, start)
				n = 1
			}
			if letter == 'a' {
				allStarts = append(allStarts, Coord{y, x})
			}
			if letter == 'E' {
				end = Coord{y, x}
				n = 26
			}
			myLine[x] = n
		}
		myMap[y] = myLine
	}

	var visited []Coord
	fmt.Println(explore(myMap, start, end, 0, visited))

	lowest := -1
	for _, s := range allStarts {
		res := explore(myMap, s, end, 0, visited)
		if (lowest == -1 || res < lowest) && res != -1 {
			lowest = res
		}
	}
	fmt.Println(lowest)
}

func explore(myMap [][]int, start, end Coord, step int, visited []Coord) int {
	if count, ok := distances[start]; ok && count <= step {
		return -1
	}
	distances[start] = step

	for _, c := range visited {
		if start.x == c.x && start.y == c.y {
			return -1
		}
	}

	if start.x == end.x && start.y == end.y {
		return step
	}
	visited = append(visited, start)
	potentialSteps := -1

	if start.x-1 >= 0 {
		if explorable(myMap[start.y][start.x-1], myMap[start.y][start.x]) {
			res := explore(myMap, Coord{start.y, start.x - 1}, end, step+1, visited)

			if res != -1 && res <= potentialSteps || potentialSteps == -1 {
				potentialSteps = res
			}
		}
	}

	if start.y-1 >= 0 {
		if explorable(myMap[start.y-1][start.x], myMap[start.y][start.x]) {
			res := explore(myMap, Coord{start.y - 1, start.x}, end, step+1, visited)
			if res != -1 && res <= potentialSteps || potentialSteps == -1 {
				potentialSteps = res
			}
		}
	}

	if start.x+1 < len(myMap[0]) {
		if explorable(myMap[start.y][start.x+1], myMap[start.y][start.x]) {
			res := explore(myMap, Coord{start.y, start.x + 1}, end, step+1, visited)
			if res != -1 && res <= potentialSteps || potentialSteps == -1 {
				potentialSteps = res
			}
		}
	}

	if start.y+1 < len(myMap) {
		if explorable(myMap[start.y+1][start.x], myMap[start.y][start.x]) {
			res := explore(myMap, Coord{start.y + 1, start.x}, end, step+1, visited)
			if res != -1 && res <= potentialSteps || potentialSteps == -1 {
				potentialSteps = res
			}
		}
	}

	return potentialSteps
}

func explorable(a, b int) bool {
	if a >= b {
		return a-b <= 1
	}
	return true
}
