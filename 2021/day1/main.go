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

	numbers := make([]int, len(values))
	for i, v := range values {
		n, _ := strconv.Atoi(v)
		numbers[i] = n
	}

	count := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] < numbers[i] {
			count++
		}
	}

	fmt.Println(fmt.Sprintf("There is %d measurements larger than the previous measurement", count))

	count = 0
	for i := 3; i < len(numbers); i++ {
		if numbers[i-1]+numbers[i-2]+numbers[i-3] < numbers[i]+numbers[i-1]+numbers[i-2] {
			count++
		}
	}

	fmt.Println(fmt.Sprintf("There is %d sums larger than the previous sum", count))
}
