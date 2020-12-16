package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	input := "9,3,1,0,8,4"
	values := strings.Split(input, ",")

	max := 30000000

	numbers := make([]int, max)
	seen := make(map[int]int)

	for i, v := range values {
		n, _ := strconv.Atoi(v)
		numbers[i] = n
		seen[n] = i
	}

	for i := len(values); i < max-1; i++ {
		if nd, ok := seen[numbers[i]]; ok {
			numbers[i+1] = i - nd
		}
		seen[numbers[i]] = i
	}

	fmt.Println(numbers[len(numbers)-1])
	totalCallTime := time.Since(startTime).Truncate(1 * time.Millisecond)
	fmt.Println("Total call time : ", totalCallTime)
}
