package main

import (
	"fmt"
	"strconv"
	"strings"
)

type numberData struct {
	lastTime     int
	lastLastTime int
}

func main() {
	input := "9,3,1,0,8,4"
	values := strings.Split(input, ",")

	numbers := make([]int, len(values))
	numberHistory := make(map[int]numberData)

	for i, v := range values {
		n, _ := strconv.Atoi(v)
		numbers[i] = n
		numberHistory[n] = numberData{lastTime: i + 1, lastLastTime: i + 1}
	}

	lastNumber := numbers[len(numbers)-1]
	// runs in ~3.5s
	for i := len(numbers) + 1; i < 30000001; i++ {
		nd, ok := numberHistory[lastNumber]
		if !ok || nd.lastTime == nd.lastLastTime {
			lastNumber = 0
			if data, ok := numberHistory[0]; !ok {
				numberHistory[0] = numberData{lastTime: i, lastLastTime: i}
			} else {
				numberHistory[0] = numberData{lastTime: i, lastLastTime: data.lastTime}
			}
		} else {
			lastNumber = nd.lastTime - nd.lastLastTime
			if data, ok := numberHistory[lastNumber]; !ok {
				numberHistory[lastNumber] = numberData{lastTime: i, lastLastTime: i}
			} else {
				numberHistory[lastNumber] = numberData{lastTime: i, lastLastTime: data.lastTime}
			}
		}
	}
	fmt.Println(lastNumber)
}
