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

	params := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}

	maxRight := len(values[0])
	res := make([]int, len(params))

	for paramID, v := range params {
		total := 0
		right, down := v[0], v[1]
		for i, j := 0, 0; i < len(values); i, j = i+down, j+right {
			if values[i][j%maxRight] == '#' {
				total++
			}
		}
		res[paramID] = total
	}

	total2 := 1
	for _, v := range res {
		total2 *= v
	}

	fmt.Println("The number of trees encoutered for the slope right 3, down 1 is =>", res[1])
	fmt.Println("The product of all the trees encoutered for each specified parameters is =>", total2)
}
