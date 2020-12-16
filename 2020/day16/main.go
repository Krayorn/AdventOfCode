package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	lowerBoundary  int
	higherBoundary int
}

type enhancedField struct {
	value        int
	passingRules map[int]bool
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n\n")

	rulesInput := strings.Split(values[0], "\n")
	rules := make([][]rule, len(rulesInput))

	for i, v := range rulesInput {
		split := strings.Split(v, ": ")
		strrules := strings.Split(split[1], " or ")
		lineRules := make([]rule, 2)
		for j, strRange := range strrules {
			boundaries := strings.Split(strRange, "-")

			lowerBoundary, _ := strconv.Atoi(boundaries[0])
			higherBoundary, _ := strconv.Atoi(boundaries[1])

			lineRules[j] = rule{lowerBoundary: lowerBoundary, higherBoundary: higherBoundary}

		}
		rules[i] = lineRules
	}

	nearbyTickets := strings.Split(values[2], "\n")[1:]
	total := 0

	var validTickets [][]enhancedField
	for _, ticket := range nearbyTickets {
		fields := strings.Split(ticket, ",")
		fieldsList := make([]enhancedField, len(fields))
		isTicketValid := true
		for i, field := range fields {
			number, _ := strconv.Atoi(field)
			fieldData := enhancedField{value: number, passingRules: map[int]bool{}}
			isNumberValid := false
			for j, ruleOptions := range rules {
				isNumberValidForRule := false
				for _, rule := range ruleOptions {
					if number >= rule.lowerBoundary && number <= rule.higherBoundary {
						isNumberValidForRule = true
					}
				}
				fieldData.passingRules[j] = isNumberValidForRule
				if isNumberValidForRule {
					isNumberValid = true
				}
			}
			if !isNumberValid {
				isTicketValid = false
				total += number
			}
			fieldsList[i] = fieldData
		}
		if isTicketValid {
			validTickets = append(validTickets, fieldsList)
		}
	}

	foundRules := make(map[int]map[int]bool, len(rules))
	for _, ticket := range validTickets {
		for fieldID, field := range ticket {
			for passingRuleID, passingRule := range field.passingRules {
				if _, ok := foundRules[fieldID]; !ok {
					foundRules[fieldID] = make(map[int]bool)
				}
				if value, ok := foundRules[fieldID][passingRuleID]; !ok {
					foundRules[fieldID][passingRuleID] = passingRule
				} else {
					if value == true {
						foundRules[fieldID][passingRuleID] = passingRule
					}
				}
			}
		}
	}

	mapping := make(map[int]int, 20)

	currentCount := 1
	for currentCount < 20 {
		for fieldID, rules := range foundRules {
			count := 0
			for _, rule := range rules {
				if rule {
					count++
				}
			}
			if count == currentCount {
				for ruleID, rule := range rules {
					if rule {
						if _, ok := mapping[ruleID]; !ok {
							mapping[ruleID] = fieldID
						}
					}
				}
				currentCount++
			}
		}
	}

	myTicket := make(map[int]int)
	fields := strings.Split(values[1][1:], ",")
	for fieldID, field := range fields {
		number, _ := strconv.Atoi(field)
		myTicket[fieldID] = number
	}

	fmt.Println("total 1", total)
	fmt.Println(myTicket[mapping[0]] * myTicket[mapping[1]] * myTicket[mapping[2]] * myTicket[mapping[3]] * myTicket[mapping[4]] * myTicket[mapping[5]])
}
