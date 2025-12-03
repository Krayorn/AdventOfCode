package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var content string

func main() {
	r, _ := regexp.Compile("mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\)")

	matches := r.FindAllStringSubmatch(content, -1)

	sum := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "don't()" {
			enabled = false
			continue
		}

		if match[0] == "do()" {
			enabled = true
			continue
		}

		if !enabled {
			continue
		}

		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		sum += n1 * n2
	}

	fmt.Println(sum)
}
