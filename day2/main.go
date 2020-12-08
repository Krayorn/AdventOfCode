package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
	firstPasswordPolicyValid := 0
	secondPasswordPolicyValid := 0

	r := regexp.MustCompile(`(\d*)-(\d*)\s([a-z]): ([a-z]*)`)
	for _, v := range values {
		matches := r.FindStringSubmatch(v)
		firstBoundary, secondBoundary, letter, password := matches[1], matches[2], matches[3], matches[4]

		index1, _ := strconv.Atoi(firstBoundary)
		index2, _ := strconv.Atoi(secondBoundary)

		count := strings.Count(password, letter)
		if count >= index1 && count <= index2 {
			firstPasswordPolicyValid++
		}

		isInFirstPosition := password[index1-1] == letter[0]
		isInSecondPosition := password[index2-1] == letter[0]
		if isInFirstPosition != isInSecondPosition {
			secondPasswordPolicyValid++
		}
	}

	fmt.Println("number of passwords valid with the first password policy =>", firstPasswordPolicyValid)
	fmt.Println("number of passwords valid with the second password policy =>", secondPasswordPolicyValid)
}
