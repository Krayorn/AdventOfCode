package main

import (
	_ "embed"
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

//go:embed input.txt
var content string

func main() {

	space := regexp.MustCompile(`\s+`)
	
	data := strings.Split(content, "\n")
	times := strings.Split(strings.Trim(space.ReplaceAllString(data[0], ""), "Time: "), " ")
	distances := strings.Split(strings.Trim(space.ReplaceAllString(data[1], ""), "Distance: "), " ")

	sum := -1
	
	for i := range times {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
	
		count := 0
		evaluate(time, distance)
		for i:= 1; i <= time; i++ {
			if evaluate(i, time) > distance {
				count++
			}
		}

		if sum == -1 {
			sum = count
		} else {
			sum *= count
		}
	}

	fmt.Println(sum)
}

func evaluate(timePressed int, time int) int {
	return (time - timePressed) * timePressed
} 