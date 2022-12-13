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
	values := strings.Split(content, "\n")

	sum := 0
	sum2 := 0
	for _, line := range values {
		elves := strings.Split(line, ",")

		bounds1 := strings.Split(elves[0], "-")
		start1, _ := strconv.Atoi(bounds1[0])
		end1, _ := strconv.Atoi(bounds1[1])

		bounds2 := strings.Split(elves[1], "-")
		start2, _ := strconv.Atoi(bounds2[0])
		end2, _ := strconv.Atoi(bounds2[1])

		if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
			sum += 1
		}

		if (start2 >= start1 && start2 <= end1) || (end2 >= start1 && end2 <= end1) || (start1 >= start2 && start1 <= end2) || (end1 >= start2 && end1 <= end2) {
			sum2 += 1
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
