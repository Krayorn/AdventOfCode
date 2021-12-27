package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

type Octopus struct {
	x, y    int
	energy  int
	flashed int
}

func main() {
	values := strings.Split(content, "\n")

	octopuses := make([][]Octopus, len(values))
	for y, str := range values {
		line := make([]Octopus, len(str))
		for x, v := range strings.Split(str, "") {
			energy, _ := strconv.Atoi(v)
			line[x] = Octopus{x, y, energy, -1}
		}
		octopuses[y] = line
	}

	counter := 0
	i := 0
	for {
		for j := range octopuses {
			for k := range octopuses[j] {
				octopuses = energizeOcto(octopuses, j, k, i, &counter)
			}
		}
		allFlashed := true
		for j := range octopuses {
			for k := range octopuses[j] {
				octo := octopuses[j][k]
				if octo.energy != 0 {
					allFlashed = false
				}
			}
		}
		i++

		if i == 100 {
			fmt.Println(fmt.Sprintf("After 100 steps, there was %d flashes.", counter))
		}

		if allFlashed {
			fmt.Println(fmt.Sprintf("The first step during which all octopuses flashes is %d.", i))
			break
		}
	}

}

func energizeOcto(octopuses [][]Octopus, j int, k int, i int, counter *int) [][]Octopus {
	octo := octopuses[j][k]

	if octo.flashed == i {
		return octopuses
	}

	octo.energy++
	if octo.energy > 9 {
		*counter++
		octo.flashed = i
		octo.energy = 0
		octopuses[j][k] = octo
		octopuses = flashAdjacent(octopuses, j, k, i, counter)
	}

	octopuses[j][k] = octo
	return octopuses
}

func flashAdjacent(octopuses [][]Octopus, j int, k int, i int, counter *int) [][]Octopus {
	for l := j - 1; l <= j+1; l++ {
		for m := k - 1; m <= k+1; m++ {
			if l < 0 || l >= len(octopuses) || m < 0 || m >= len(octopuses[j]) {
				continue
			}
			octopuses = energizeOcto(octopuses, l, m, i, counter)
		}
	}

	return octopuses
}

func printOctopuses(octopuses [][]Octopus) {
	for y := range octopuses {
		for x := range octopuses[y] {
			fmt.Print(octopuses[y][x].energy)
		}
		fmt.Println("")
	}
}
