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

	columns := make(map[int]int, 0)

	for j, line := range lines {
		chars := strings.Split(line, "")
		row := make([]string, len(chars))
		for i, char := range chars {
			row[i] = char
			if char == "S" {
				columns[i] = 1
			}
		}
		grid[j] = row
	}

	sum := 0
	for _, line := range grid {
		for key, val := range columns {
			if val == -1 {
				continue
			}

			if line[key] == "^" {
				sum++
				if key > 0 {
					if _, ok := columns[key-1]; !ok {
						columns[key-1] = columns[key]
					} else {
						columns[key-1] += columns[key]
					}
				}
				if key < len(line)-1 {
					if _, ok := columns[key+1]; !ok {
						columns[key+1] = columns[key]
					} else {
						columns[key+1] += columns[key]
					}
				}
				delete(columns, key)
			}
		}
	}

	sum2 := 0
	for _, val := range columns {
		if val > 0 {
			sum2 += val
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
