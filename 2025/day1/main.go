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
	lines := strings.Split(content, "\n")

	dial := 50
	sum := 0
	for _, line := range lines {
		side := string(line[0])
		number, _ := strconv.Atoi(line[1:])

		n := int(math.Floor(float64(number / 100)))
		sum += n

		number = number % 100

		if side == "L" {
			dial -= number
			if dial < 0 && dial != -number {
				sum++
			}
		} else {
			dial += number
			if dial > 100 && dial-number != 100 {
				sum++
			}
		}

		dial = (dial + 100) % 100

		if dial == 0 {
			sum++
		}
	}

	fmt.Println(sum)

}
