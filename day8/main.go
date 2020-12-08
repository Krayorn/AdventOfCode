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

	fmt.Println("acc before code execution start to loop => ", executeInstructions(values, false))
	fmt.Println("acc once faulty instruction is repaired => ", executeInstructions(values, true))
}

func executeInstructions(values []string, withSwitching bool) int {
	visitedIds := make(map[int]bool)

	id := 0
	acc := 0

	switchedIds := make(map[int]bool)
	switching := false

	for {
		if id == len(values) {
			break
		}

		line := values[id]
		instruction := line[0:3]
		sign := line[4]
		value, _ := strconv.Atoi(line[5:len(line)])

		if withSwitching && !switching && (instruction == "jmp" || instruction == "nop") && !switchedIds[id] {
			switchedIds[id] = true
			switching = true
			if instruction == "jmp" {
				instruction = "nop"
			} else {
				instruction = "jmp"
			}
		}

		if visitedIds[id] {
			if !withSwitching {
				break
			}
			switching = false
			visitedIds = make(map[int]bool)
			id = 0
			acc = 0
			continue
		}

		visitedIds[id] = true
		if instruction == "acc" {
			if sign == '+' {
				acc += value
			} else {
				acc -= value
			}
			id++
			continue
		}

		if instruction == "jmp" {
			if sign == '+' {
				id += value
			} else {
				id -= value
			}
			continue
		}

		if instruction == "nop" {
			id++
			continue
		}

	}
	return acc
}
