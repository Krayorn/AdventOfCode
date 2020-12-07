package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	maxID := 0
	ids := make([]int, len(values))

	for i, v := range values {
		minR, maxR := 0, 127
		minC, maxC := 0, 7
		for _, c := range v {
			if c == 'F' {
				maxR = int(math.Floor(float64(minR + (maxR-minR)/2)))
			}
			if c == 'B' {
				minR = int(math.Ceil(float64(minR + (maxR-minR)/2 + 1)))
			}
			if c == 'L' {
				maxC = int(math.Floor(float64(minC + (maxC-minC)/2)))
			}
			if c == 'R' {
				minC = int(math.Ceil(float64(minC + (maxC-minC)/2 + 1)))
			}
		}
		id := minR*8 + minC
		ids[i] = id
		if id > maxID {
			maxID = id
		}
	}
	fmt.Println("maxId", maxID)
	sort.Ints(ids)

	var currentID int
	for _, id := range ids {
		if currentID == 0 {
			currentID = id
			continue
		}
		if id-currentID != 1 {
			fmt.Println("MyId between", currentID, id)
		}
		currentID = id
	}
}
