package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Zone struct {
	minX, maxX, minY, maxY int
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	data := strings.Split(strings.Trim(string(byteValue), "target area: "), ", ")
	xSplit := strings.Split(strings.Trim(data[0], "x="), "..")

	minX, _ := strconv.Atoi(xSplit[0])
	maxX, _ := strconv.Atoi(xSplit[1])

	ySplit := strings.Split(strings.Trim(data[1], "y="), "..")

	maxY, _ := strconv.Atoi(ySplit[0])
	minY, _ := strconv.Atoi(ySplit[1])

	defaultXVelocity := maxX

	zone := Zone{minX, maxX, minY, maxY}
	count := 0
	bestY := 0
	for i := defaultXVelocity; i > 0; i-- {
		for j := maxY; j < 1000; j++ {
			win, highY := goesThroughZone(i, j, zone)
			if win {
				count++
				if highY > bestY {
					bestY = highY
				}
			}
		}
	}
	fmt.Println(bestY, count)
}

func goesThroughZone(xVelocity, yVelocity int, zone Zone) (bool, int) {
	x, y := 0, 0
	bestY := 0
	for {
		if y > bestY {
			bestY = y
		}
		if x > zone.maxX || y < zone.maxY {
			return false, bestY
		}
		if x <= zone.maxX && x >= zone.minX && y >= zone.maxY && y <= zone.minY {
			return true, bestY
		}

		x += xVelocity
		y += yVelocity
		if xVelocity > 0 {
			xVelocity--
		} else if xVelocity < 0 {
			xVelocity++
		} else {
			if x < zone.minX {
				return false, bestY
			}
		}
		yVelocity--
	}
}
