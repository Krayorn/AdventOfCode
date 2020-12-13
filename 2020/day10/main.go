package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
	numbers = append(numbers, 0)

	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	diffCount := map[int]int{1: 0, 2: 0, 3: 0}

	last := 0
	for _, n := range numbers {
		diff := n - last
		diffCount[diff]++
		last = n
	}

	fmt.Println(diffCount[1] * diffCount[3])

	possibilities := make(map[int][]int)
	for i := 0; i < len(numbers)-1; i++ {
		possibilities[numbers[i]] = []int{}
		for j := 1; j < 4 && i+j < len(numbers); j++ {
			if numbers[i+j]-numbers[i] <= 3 {
				possibilities[numbers[i]] = append(possibilities[numbers[i]], numbers[i+j])
			}
		}
	}
	fmt.Println(getNumberOfPossibilities(0, possibilities))
}

var memory = make(map[int]int)

func getNumberOfPossibilities(entry int, possibilities map[int][]int) int {
	if _, ok := memory[entry]; ok {
		return memory[entry]
	}
	if _, ok := possibilities[entry]; !ok {
		return 1
	}
	count := 0
	for _, p := range possibilities[entry] {
		count += getNumberOfPossibilities(p, possibilities)
	}
	memory[entry] = count
	return count
}
