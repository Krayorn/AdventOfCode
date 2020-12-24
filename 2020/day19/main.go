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
	values := strings.Split(string(byteValue), "\n\n")

	rules := make(map[int][][]int)
	end := make(map[int]string)
	for _, rule := range strings.Split(values[0], "\n") {
		split := strings.Split(rule, ": ")

		ruleNumber, _ := strconv.Atoi(split[0])

		if strings.Contains(split[1], "\"") {
			end[ruleNumber] = string(split[1][1])
			continue
		}

		paths := make([][]int, 0)
		for _, opt := range strings.Split(split[1], " | ") {
			path := make([]int, 0)
			for _, ruleID := range strings.Split(opt, " ") {
				n, _ := strconv.Atoi(ruleID)
				path = append(path, n)
			}
			paths = append(paths, path)
		}
		rules[ruleNumber] = paths

	}
	fmt.Println(rules)
	fmt.Println(end)
	total := 0
	for _, message := range strings.Split(values[1], "\n") {
		fmt.Println(message)
		if isValid(message, rules, end, []int{0}) {
			total++
		}
	}
	fmt.Println("total =>", total)
}

func isValid(message string, rules map[int][][]int, end map[int]string, rule []int) bool {
	if len(rule) == 0 {
		return len(message) == 0
	}

	var letter string
	if _, ok := rules[rule[0]]; !ok {
		letter = end[rule[0]]
	}

	currentRule := rules[rule[0]]
	rule = rule[1:]

	if letter != "" {
		if message == "" {
			return false
		}
		return string(message[0]) == letter && isValid(message[1:], rules, end, rule)
	}
	for _, currentRule := range currentRule {
		newRules := append(currentRule, rule...)
		if isValid(message, rules, end, newRules) {
			return true
		}
	}
	return false
}

// func isValid(message string, rules map[int]string, ruleIDStart int) int {
// 	rule := rules[ruleIDStart]
// 	if strings.Contains(rule, "\"") {
// 		if message[0] == rule[1] {
// 			return 1
// 		}
// 		return -1
// 	}
// 	for _, opt := range strings.Split(rule, " | ") {
// 		offset := 0
// 		match := true
// 		for _, ruleID := range strings.Split(opt, " ") {
// 			n, _ := strconv.Atoi(ruleID)
// 			validity := isValid(message[offset:], rules, n)
// 			if validity == -1 {
// 				match = false
// 				break
// 			}
// 			offset += validity
// 			if offset == len(message) {
// 				return offset
// 			}
// 		}
// 		if match == true {
// 			return offset
// 		}
// 	}

// 	return -1
// }
