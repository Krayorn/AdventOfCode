package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	lines := strings.Split(content, "\n")

	w, h := 101, 103
	//w, h := 11, 7

	r, _ := regexp.Compile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	iter := 0

	lowestSafety := -1
	for {
		iter++
		robots := make([][]int, 0)
		q1, q2, q3, q4 := 0, 0, 0, 0

		for _, line := range lines {
			matches := r.FindStringSubmatch(line)

			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])

			dirX, _ := strconv.Atoi(matches[3])
			dirY, _ := strconv.Atoi(matches[4])

			endX := (x + iter*dirX) % w
			endY := (y + iter*dirY) % h

			if endX < 0 {
				endX += w
			}

			if endY < 0 {
				endY += h
			}

			robots = append(robots, []int{endX, endY})

			if endX < (w-1)/2 {
				if endY < (h-1)/2 {
					q1++
				} else if endY > (h-1)/2 {
					q3++
				}
			} else if endX > (w-1)/2 {
				if endY < (h-1)/2 {
					q2++
				} else if endY > (h-1)/2 {
					q4++
				}
			}
		}

		safety := q1 * q2 * q3 * q4
		if lowestSafety == -1 || safety < lowestSafety {
			lowestSafety = safety
			fmt.Println(iter, lowestSafety)
		}

		if iter == 10000 {
			break
		}
	}
}
