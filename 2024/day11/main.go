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
	inputs := strings.Split(content, " ")

	numbers := make(map[int]int, len(inputs))
	for _, input := range inputs {
		n, _ := strconv.Atoi(input)
		numbers[n]++
	}

	for i := 0; i < 75; i++ {
		newNumbers := make(map[int]int, 0)
		for stone, count := range numbers {
			if stone == 0 {
				newNumbers[1] += count
				continue
			}
			str := strconv.Itoa(stone)
			if len(str)%2 == 0 {
				n1, _ := strconv.Atoi(str[0 : len(str)/2])
				n2, _ := strconv.Atoi(str[len(str)/2:])
				newNumbers[n1] += count
				newNumbers[n2] += count
				continue
			}
			newNumbers[stone*2024] += count
		}
		numbers = newNumbers
	}
	sum := 0
	for _, count := range numbers {
		sum += count
	}
	fmt.Println(sum)

}
