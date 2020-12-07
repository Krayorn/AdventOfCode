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
	values := strings.Split(string(byteValue), "\n")

	maxRight := len(values[0])

	res := 1
	options := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}

	for _, v := range options {
		total := 0
		right, down := v[0], v[1]
		for i, j := 0, 0; i < len(values); i, j = i+down, j+right {
			if values[i][j%maxRight] == '#' {
				total++
			}
		}
		res = res * total
	}
	fmt.Println(res)
}
