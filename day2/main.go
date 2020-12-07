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
	total := 0

	// for _, v := range values {
	// 	if v == "" {
	// 		continue
	// 	}
	// 	split := strings.Split(v, ":")
	// 	rules := strings.Split(split[0], " ")
	// 	letter := rules[1]
	// 	numbers := strings.Split(rules[0], "-")

	// 	min, _ := strconv.Atoi(numbers[0])
	// 	max, _ := strconv.Atoi(numbers[1])

	// 	count := strings.Count(split[1], letter)
	// 	if count >= min && count <= max {
	// 		total++
	// 	}
	// }

	for _, v := range values {
		split := strings.Split(v, ":")
		rules := strings.Split(split[0], " ")
		letter := rules[1]
		numbers := strings.Split(rules[0], "-")

		index1, _ := strconv.Atoi(numbers[0])
		index2, _ := strconv.Atoi(numbers[1])
		if (string(split[1][index1]) == letter || string(split[1][index2]) == letter) && (split[1][index1] != split[1][index2]) {
			total++
		}
	}

	fmt.Println(total)
}
