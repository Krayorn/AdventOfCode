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

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		ns := make([]int, len(numbers))
		for i, number := range numbers {
			n, _ := strconv.Atoi(number)
			ns[i] = n
		}

		safe := isSafe(ns)

		if !safe {
			for i := 0; i < len(ns); i++ {
				newNs := append([]int{}, ns[:i]...)
				newNs = append(newNs, ns[i+1:]...)
				if isSafe(newNs) {
					safe = true
					break
				}
			}
		}

		if safe {
			sum++
		}
	}

	fmt.Println(sum)
}

func isSafe(report []int) bool {
	if report[0] < report[1] {
		for i := 1; i < len(report); i++ {
			if report[i] <= report[i-1] {
				return false
			}

			if report[i]-report[i-1] > 3 {
				return false
			}
		}
	} else {
		for i := 1; i < len(report); i++ {
			if report[i-1] <= report[i] {
				return false
			}

			if report[i-1]-report[i] > 3 {
				return false
			}
		}
	}
	return true
}
