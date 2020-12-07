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

	total := 0
	total2 := 0

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
		total += len(groupYes)
		total2 += groupTotal
	}
	fmt.Println(total)
	fmt.Println(total2)
}
