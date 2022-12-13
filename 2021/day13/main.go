package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	points := make([]Point, 0)
	dots := true
	inventory := make(map[string]bool, 0)

	for _, str := range values {
		if str == "" {
			dots = false
			continue
		}
		if dots {
			split := strings.Split(str, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			points = append(points, Point{x: x, y: y})
		} else {
			instruc := strings.Split(strings.TrimLeft(str, "fold along "), "=")
			loca := instruc[0]
			line, _ := strconv.Atoi(instruc[1])

			for i, point := range points {
				if loca == "y" && point.y > line {
					point = Point{x: point.x, y: point.y - ((point.y - line) * 2)}
				} else if loca == "x" && point.x > line {
					point = Point{y: point.y, x: point.x - ((point.x - line) * 2)}
				}
				points[i] = point
			}

			inventory = make(map[string]bool, 0)
			for _, point := range points {
				inventory[strconv.Itoa(point.x)+","+strconv.Itoa(point.y)] = true
			}
			// fmt.Println(len(inventory))
			// break
		}
	}
	highestY, highestX := 0, 0
	for _, point := range points {
		if point.x > highestX {
			highestX = point.x
		}
		if point.y > highestY {
			highestY = point.x
		}
	}

	printInventory(inventory, highestX, highestY)
}

func printInventory(inventory map[string]bool, x int, y int) {
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			_, ok := inventory[strconv.Itoa(j)+","+strconv.Itoa(i)]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
