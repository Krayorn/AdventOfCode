package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	lines := strings.Split(content, "\n")

	listA := make([]int, len(lines))
	listB := make([]int, len(lines))
	mapB := make(map[int]int)
	for i, line := range lines {
		numbers := strings.Split(line, "   ")
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])
		listA[i] = a
		listB[i] = b
		if _, ok := mapB[b]; !ok {
			mapB[b] = 0
		}
		mapB[b] = mapB[b] + 1
	}

	sort.Ints(listA)
	sort.Ints(listB)

	sum := 0
	sim := 0
	for i := range listA {
		sum += int(math.Abs(float64(listA[i]) - float64(listB[i])))

		if _, ok := mapB[listA[i]]; ok {
			sim += listA[i] * mapB[listA[i]]
		}
	}

	fmt.Println(sum)
	fmt.Println(sim)
}
