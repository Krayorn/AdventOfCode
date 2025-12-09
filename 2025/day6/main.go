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
	lines := strings.Split(content, "\n")
	sum := 0

	numbers := make([]int, 0)
	operator := "+"
	for i := 0; i < len(lines[0]); i++ {
		number := ""
		for j, line := range lines {
			relevant := line[i]
			if j == len(lines)-1 {
				if string(relevant) != " " {
					operator = string(relevant)
				}
			} else {
				number += string(relevant)
			}
		}

		trimmed := strings.Trim(number, " ")
		if trimmed == "" {
			operationSum := 0
			if operator == "*" {
				operationSum = 1
			}
			for _, n := range numbers {
				if operator == "*" {
					operationSum *= n
				} else {
					operationSum += n
				}
			}
			fmt.Println(numbers, operationSum)
			sum += operationSum
			numbers = make([]int, 0)
			number = ""
			continue
		}
		n, _ := strconv.Atoi(trimmed)
		numbers = append(numbers, n)
		number = ""
	}

	operationSum := 0
	if operator == "*" {
		operationSum = 1
	}
	for _, n := range numbers {
		if operator == "*" {
			operationSum *= n
		} else {
			operationSum += n
		}
	}
	fmt.Println(numbers, operationSum)
	sum += operationSum

	fmt.Println(sum)
}
