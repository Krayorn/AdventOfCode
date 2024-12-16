package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var content string

type Item struct {
	Pos   []int
	Dir   int
	Score int
	Path  [][]int
}

func main() {
	lines := strings.Split(content, "\n")

	grid := make([][]string, len(lines))
	queue := make([]Item, 1)

	for i, line := range lines {
		chars := strings.Split(line, "")
		row := make([]string, len(chars))
		for j, char := range chars {
			row[j] = char
			if char == "S" {
				queue[0] = Item{Pos: []int{i, j}, Dir: 0, Score: 0}
			}
		}
		grid[i] = row
	}

	visited := make(map[string]int)
	directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	target := 107468
	saves := make([][][]int, 0)
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if grid[item.Pos[0]][item.Pos[1]] == "E" {
			fmt.Println(item.Score)
			if item.Score == target {
				saves = append(saves, item.Path)
			}
			continue
		}

		if item.Score > target {
			continue
		}

		if val, ok := visited[fmt.Sprintf("%d-%d-%d", item.Pos[0], item.Pos[1], item.Dir)]; ok && val < item.Score {
			continue
		}
		visited[fmt.Sprintf("%d-%d-%d", item.Pos[0], item.Pos[1], item.Dir)] = item.Score

		newI, newJ := item.Pos[0]+directions[item.Dir][0], item.Pos[1]+directions[item.Dir][1]
		if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] != "#" {
			newPath := make([][]int, len(item.Path))
			for i := range item.Path {
				newPath[i] = make([]int, len(item.Path[i]))
				copy(newPath[i], item.Path[i])
			}

			queue = append(queue, Item{Pos: []int{newI, newJ}, Dir: item.Dir, Score: item.Score + 1, Path: append(newPath, []int{newI, newJ})})
		}

		queue = append(queue,
			Item{Pos: item.Pos, Dir: (item.Dir + 1) % 4, Score: item.Score + 1000, Path: item.Path},
			Item{Pos: item.Pos, Dir: (item.Dir + 3) % 4, Score: item.Score + 1000, Path: item.Path},
		)
	}

	goodPlaces := make(map[string]bool)

	for _, save := range saves {
		for _, coords := range save {
			goodPlaces[fmt.Sprintf("%d-%d", coords[0], coords[1])] = true
		}
	}

	fmt.Println(len(goodPlaces) + 1)

}
