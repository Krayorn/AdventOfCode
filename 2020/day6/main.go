package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n\n")

	questionsWithSomeYesCount := 0
	questionsWithOnlyYesCount := 0

	for _, group := range values {
		groupYes := map[rune]int{}
		persons := strings.Split(group, "\n")
		for _, person := range persons {
			for _, question := range person {
				if _, ok := groupYes[question]; ok {
					groupYes[question]++
				} else {
					groupYes[question] = 1
				}
			}
		}
		groupTotal := 0
		for _, count := range groupYes {
			if count == len(persons) {
				groupTotal++
			}
		}
		questionsWithSomeYesCount += len(groupYes)
		questionsWithOnlyYesCount += groupTotal
	}
	fmt.Println("The sum of the counts of questions to which anyone answered yes is =>", questionsWithSomeYesCount)
	fmt.Println("The sum of the counts of questions to which everyone answered yes is =>", questionsWithOnlyYesCount)
}
