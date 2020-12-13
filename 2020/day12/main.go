package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
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

	fmt.Println("Manhattan distance between initial position of the ship and final position =>", executeInstructions(values, false))
	fmt.Println("Manhattan distance between initial position of the ship and final position wuth waypoint navigation =>", executeInstructions(values, true))
}

func executeInstructions(instructions []string, withWaypoint bool) float64 {
	ns, oe := 0, 0
	wpns, wpoe := 1, 10
	direction := "E"
	right, left := "NESWNESW", "NWSENWSE"

	for _, v := range instructions {
		instruction := string(v[0:1])
		number, _ := strconv.Atoi(v[1:len(v)])

		if instruction == "F" {
			if withWaypoint {
				ns += number * wpns
				oe += number * wpoe
			} else {
				instruction = direction
			}
		}

		shiftIndex := 4 - (360-number)/90
		if instruction == "R" {
			if withWaypoint {
				for i := 0; i < shiftIndex; i++ {
					wpoe, wpns = wpns, wpoe*-1
				}
			} else {
				index := strings.Index(right, direction)
				direction = string(right[index+shiftIndex])
			}
		} else if instruction == "L" {
			if withWaypoint {
				for i := 0; i < shiftIndex; i++ {
					wpoe, wpns = wpns*-1, wpoe
				}
			} else {
				index := strings.Index(left, direction)
				direction = string(left[index+shiftIndex])
			}
		}

		if withWaypoint {
			if instruction == "N" {
				wpns += number
			} else if instruction == "S" {
				wpns -= number
			} else if instruction == "W" {
				wpoe -= number
			} else if instruction == "E" {
				wpoe += number
			}
		} else {
			if instruction == "N" {
				ns += number
			} else if instruction == "S" {
				ns -= number
			} else if instruction == "W" {
				oe -= number
			} else if instruction == "E" {
				oe += number
			}
		}

	}

	return math.Abs(float64(ns)) + math.Abs(float64(oe))
}
