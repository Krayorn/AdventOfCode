package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, ",")

	numbers := make([]int, len(values))
	value, _ := strconv.Atoi(values[0])
	min, max := value, value
	for i, v := range values {
		value, _ := strconv.Atoi(v)
		numbers[i] = value

		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	bestFuel, bestFuel2 := -1, -1
	for i := min; i <= max; i++ {
		fuelCount, fuelCount2 := 0, 0
		for _, crab := range numbers {
			number := int(math.Abs(float64(i - crab)))
			fuelCount += number
			fuelCount2 += (number*number + number) / 2
		}
		if fuelCount <= bestFuel || bestFuel == -1 {
			bestFuel = fuelCount
		}
		if fuelCount2 <= bestFuel2 || bestFuel2 == -1 {
			bestFuel2 = fuelCount2
		}
	}
	fmt.Println(fmt.Sprintf("The least fuel needed to align every crab submarine using regular engineering is %d", bestFuel))
	fmt.Println(fmt.Sprintf("The least fuel needed to align every crab submarine using crab engineering is %d", bestFuel2))
}
