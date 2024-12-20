package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	lines := strings.Split(content, "\n")

	//A, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	B, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	C, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])

	nums := strings.Split(strings.ReplaceAll(lines[4], "Program: ", ""), ",")

	a := 0
	for pos := len(nums) - 1; pos >= 10; pos-- {
		a <<= 3
		for !slices.Equal(run(a, B, C, nums), nums[pos:]) {
			fmt.Println(run(a, B, C, nums), a, nums[pos:])
			a++
		}
		fmt.Println("success", a, nums[pos:])
	}

	fmt.Println(a)
}

func run(a, b, c int, nums []string) []string {
	output := make([]string, 0)
	for i := 0; i < len(nums); i += 2 {
		opcode, _ := strconv.Atoi(nums[i])
		operand, _ := strconv.Atoi(nums[i+1])

		comboOperand := operand
		if operand == 4 {
			comboOperand = a
		}
		if operand == 5 {
			comboOperand = b
		}
		if operand == 6 {
			comboOperand = c
		}

		if opcode == 0 {
			a = int(math.Floor(float64(a) / math.Pow(float64(2), float64(comboOperand))))
		} else if opcode == 1 {
			b = operand ^ b
		} else if opcode == 2 {
			b = comboOperand % 8
		} else if opcode == 3 {
			if a != 0 {
				i = operand - 2
			}
		} else if opcode == 4 {
			b = b ^ c
		} else if opcode == 5 {
			output = append(output, strconv.Itoa(comboOperand%8))
		} else if opcode == 6 {
			b = int(math.Floor(float64(a) / math.Pow(float64(2), float64(comboOperand))))
		} else if opcode == 7 {
			c = int(math.Floor(float64(a) / math.Pow(float64(2), float64(comboOperand))))
		}
	}

	return output
}
