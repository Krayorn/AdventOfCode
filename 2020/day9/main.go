package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
	var invalidNumber int

	numbers := make([]int, len(values))
	for i, v := range values {
		number, _ := strconv.Atoi(v)
		numbers[i] = number
	}

	for i := 25; i < len(numbers); i++ {
		number := numbers[i]
		isValid := false
		for j := i - 25; j < i; j++ {
			n1 := numbers[j]
			for k := j + 1; k < i; k++ {
				n2 := numbers[k]
				if n1+n2 == number {
					isValid = true
				}
			}
		}
		if !isValid {
			invalidNumber = number
			break
		}
	}

	firstIndex, lastIndex := 0, 0
	sum := 0
	for i := 0; i < len(numbers); i++ {
		number := numbers[i]
		sum += number

		for sum > invalidNumber {
			numberToRemove := numbers[firstIndex]
			sum -= numberToRemove
			firstIndex++
		}
		if sum == invalidNumber {
			lastIndex = i
			break
		}
	}

	list := numbers[firstIndex : lastIndex+1]
	sort.Ints(list)
	fmt.Println(list[0] + list[len(list)-1])
}
