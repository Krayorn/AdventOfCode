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

	antennasByKind := make(map[string][][]int)
	antiNodes := make(map[string]bool)

	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			if char != "." {
				if _, ok := antennasByKind[char]; !ok {
					antennasByKind[char] = make([][]int, 0)
				}
				antennasByKind[char] = append(antennasByKind[char], []int{i, j})
			}
		}
	}

	for _, antennas := range antennasByKind {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}
				first := antennas[i]
				second := antennas[j]
				
				shiftI := second[0] - first[0]
				shiftJ := second[1] - first[1]

				newI := first[0]
				newJ := first[1]
				for {
					if newI >= 0 && newJ >= 0 && newI < len(lines) && newJ < len(lines[0]) {
						antiNodes[fmt.Sprintf("%d-%d", newI, newJ)] = true
					} else {
						break
					}
					newI += shiftI
					newJ += shiftJ
				}
			}
		}
	}
	fmt.Println(len(antiNodes))

}
