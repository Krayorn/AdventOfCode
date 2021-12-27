package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Location struct {
	i, j int
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	cave := make([][]int, len(values))
	for i, l := range values {
		line := make([]int, len(l))
		cave[i] = line
		lineValues := strings.Split(l, "")
		for j, v := range lineValues {
			number, _ := strconv.Atoi(v)
			line[j] = number
		}
	}

	low := make([]Location, 0)
	for i, line := range cave {
		for j, value := range line {
			if i+1 < len(cave) && cave[i+1][j] <= value {
				continue
			}
			if i-1 >= 0 && cave[i-1][j] <= value {
				continue
			}
			if j+1 < len(line) && cave[i][j+1] <= value {
				continue
			}
			if j-1 >= 0 && cave[i][j-1] <= value {
				continue
			}
			low = append(low, Location{i, j})
		}
	}

	count := 0
	for _, v := range low {
		count += cave[v.i][v.j] + 1
	}

	bassinsValues := make([]int, 0)
	for _, v := range low {
		bassin := make(map[string]bool, 0)
		bassin = addLocationToBassin(v.i, v.j, bassin, cave)
		bassinsValues = append(bassinsValues, len(bassin))
	}

	sort.Slice(bassinsValues, func(i, j int) bool {
		return bassinsValues[i] > bassinsValues[j]
	})
	mult := 1
	for i := 0; i < 3; i++ {
		mult *= bassinsValues[i]
	}

	fmt.Println(fmt.Sprintf("The sum of the risks levels of all low points on the heightmap is %d.", count))
	fmt.Println(fmt.Sprintf("The size of the three largets basins multiplies by 3 is %d.", mult))
}

func addLocationToBassin(i int, j int, bassin map[string]bool, cave [][]int) map[string]bool {
	value := cave[i][j]
	bassin[strconv.Itoa(i)+","+strconv.Itoa(j)] = true

	if i+1 < len(cave) && cave[i+1][j] > value && cave[i+1][j] != 9 {
		bassin = addLocationToBassin(i+1, j, bassin, cave)
	}
	if i-1 >= 0 && cave[i-1][j] > value && cave[i-1][j] < 9 {
		bassin = addLocationToBassin(i-1, j, bassin, cave)
	}
	if j+1 < len(cave[i]) && cave[i][j+1] > value && cave[i][j+1] < 9 {
		bassin = addLocationToBassin(i, j+1, bassin, cave)
	}
	if j-1 >= 0 && cave[i][j-1] > value && cave[i][j-1] < 9 {
		bassin = addLocationToBassin(i, j-1, bassin, cave)
	}

	return bassin
}
