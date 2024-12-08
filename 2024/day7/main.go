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
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])

		ns := strings.Split(parts[1], " ")
		options := make([]int, 0)
		for _, n := range ns {
			number, _ := strconv.Atoi(n)
			if len(options) == 0 {
				options = append(options, number)
				continue
			}
			newOptions := make([]int, 0)
			for _, option := range options {
				concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", option, number))
				newOptions = append(newOptions, option+number, option*number, concat)
			}

			options = newOptions
		}

		for _, option := range options {
			if option == target {
				sum += target
				break
			}
		}
	}
	fmt.Print(sum)
}
