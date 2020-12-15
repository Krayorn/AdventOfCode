package main

import (
	"fmt"
	"strconv"
	"strings"
)

type numberData struct {
	turns []int
	count int
}

func main() {
	input := "9,3,1,0,8,4"
	values := strings.Split(input, ",")

	numbers := make([]int, len(values))
	numberHistory := make(map[int]numberData)

	for i, v := range values {
		n, _ := strconv.Atoi(v)
		numbers[i] = n
		numberHistory[n] = numberData{turns: []int{i + 1}, count: 1}
	}

	lastNumber := numbers[len(numbers)-1]
	// runs in 6s
	for i := len(numbers) + 1; i < 30000001; i++ {
		nd, ok := numberHistory[lastNumber]
		if !ok || nd.count == 1 {
			lastNumber = 0
			if data, ok := numberHistory[0]; !ok {
				numberHistory[0] = numberData{turns: []int{i}, count: 1}
			} else {
				data.turns = append(data.turns, i)
				data.count++
				numberHistory[0] = data
			}
		} else {
			lastNumber = nd.turns[len(nd.turns)-1] - nd.turns[len(nd.turns)-2]
			if data, ok := numberHistory[lastNumber]; !ok {
				numberHistory[lastNumber] = numberData{turns: []int{i}, count: 1}
			} else {
				data.turns = append(data.turns, i)
				data.count++
				numberHistory[lastNumber] = data
			}
		}
	}
	fmt.Println(lastNumber)
}
