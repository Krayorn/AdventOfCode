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

	totalNoOrder := 0
	totalAddFirst := 0
	for _, v := range values {
		totalNoOrder += executeInstructions(v, false)
		totalAddFirst += executeInstructions(v, true)
	}

	fmt.Println(totalNoOrder)
	fmt.Println(totalAddFirst)
}

func executeInstructions(calcul string, withAddFirst bool) int {
	for {
		startParenthesis := -1
		for i, c := range calcul {
			if c == '(' {
				startParenthesis = i
			}
			if c == ')' {
				res := executeInstructions(calcul[startParenthesis+1:i], withAddFirst)
				calcul = strings.ReplaceAll(calcul, calcul[startParenthesis:i+1], fmt.Sprintf("%d", res))
				break
			}
		}
		if startParenthesis == -1 {
			break
		}
	}

	instructions := strings.Split(calcul, " ")
	if withAddFirst {
		for i := 0; i < len(instructions); i++ {
			instruction := instructions[i]
			if instruction == "+" {
				n1, _ := strconv.Atoi(instructions[i-1])
				n2, _ := strconv.Atoi(instructions[i+1])
				newInstructions := instructions[0 : i-1]
				newInstructions = append(newInstructions, fmt.Sprintf("%d", n1+n2))
				newInstructions = append(newInstructions, instructions[i+2:]...)
				i--
				instructions = newInstructions
			}
		}
	}

	total := 0
	lastOperator := "+"
	for _, c := range instructions {
		if c == "+" || c == "*" {
			lastOperator = c
			continue
		}

		n, _ := strconv.Atoi(c)
		if lastOperator == "+" {
			total += n
		} else if lastOperator == "*" {
			total *= n
		}
	}
	return total
}
