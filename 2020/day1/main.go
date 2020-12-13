package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	numbers := make([]int, len(values))
	for i, v := range values {
		number, _ := strconv.Atoi(v)
		numbers[i] = number
	}

	fmt.Println(getSumProduct(2020, numbers, 2))
	fmt.Println(getSumProduct(2020, numbers, 3))
}

// 2 first implem => 40 200
// 2 second case => 20 300
// 3 first implem => 8 040 200
// 3 second case => 1 373 700
func getSumProduct(target int, numbers []int, factorsNumber int) int {
	for i, n1 := range numbers {
		for j := i; j < len(numbers); j++ {
			n2 := numbers[j]
			if factorsNumber == 2 && n1+n2 == 2020 {
				return n1 * n2
			}
			if factorsNumber >= 3 {
				for k := j; k < len(numbers); k++ {
					n3 := numbers[k]
					if factorsNumber == 3 && n1+n2+n3 == 2020 {
						return n1 * n2 * n3
					}
				}
			}
		}
	}

	return -1
}
