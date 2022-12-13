package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, "\n")
	lines := make([]string, len(values))
	for i, line := range values {
		lines[i] = line
	}

	res := lines[0]

	for i := 1; i < len(lines); i++ {
		line := "[" + res + "," + lines[i] + "]"
		active := true
		for active {
			line, active = explode(line)
			if !active {
				line, active = split(line)
			}
		}
		res = line
	}

	fmt.Println(getMagnitude(res))

	bestMagnitude := -1
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			line := "[" + lines[i] + "," + lines[j] + "]"
			active := true
			for active {
				line, active = explode(line)
				if !active {
					line, active = split(line)
				}
			}
			sumMagnitude := getMagnitude(line)
			if sumMagnitude > bestMagnitude {
				bestMagnitude = sumMagnitude
			}
		}
	}
	fmt.Println(bestMagnitude)
}

func getMagnitude(line string) int {
	for deps := 4; deps > 0; deps-- {
		dep := 0
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char == "[" {
				dep++
			} else if char == "]" {
				dep--
			}

			if dep == deps {
				stop := -1
				for j := i; j < len(line); j++ {
					if string(line[j]) == "]" {
						stop = j
						break
					}
				}

				split := strings.Split(line[i+1:stop], ",")
				left, _ := strconv.Atoi(split[0])
				right, _ := strconv.Atoi(split[1])

				magnitude := strconv.Itoa(3*left + 2*right)

				line = line[:i] + magnitude + line[stop+1:]
				dep--
			}
		}
	}

	total, _ := strconv.Atoi(line)
	return total
}

func explode(line string) (string, bool) {
	deps := 0
	for i, r := range line {
		char := string(r)
		if char == "[" {
			deps++
		} else if char == "]" {
			deps--
		}

		if deps == 5 {
			stop := -1
			for j := i; j < len(line); j++ {
				if string(line[j]) == "]" {
					stop = j
					break
				}
			}

			split := strings.Split(line[i+1:stop], ",")
			left, _ := strconv.Atoi(split[0])
			right, _ := strconv.Atoi(split[1])
			shiftLeft := 0
			for j := stop + 1; j < len(line); j++ {
				char := string(line[j])
				if char != "]" && char != "[" && char != "," {
					nextChar := string(line[j+1])
					if nextChar != "]" && nextChar != "[" && nextChar != "," {
						n, _ := strconv.Atoi(line[j : j+2])
						line = line[:j] + strconv.Itoa(n+right) + line[j+2:]
					} else {
						n, _ := strconv.Atoi(char)

						line = line[:j] + strconv.Itoa(n+right) + line[j+1:]
					}
					break
				}
			}
			for j := i; j > 0; j-- {
				char := string(line[j])
				if char != "]" && char != "[" && char != "," {
					nextChar := string(line[j-1])
					if nextChar != "]" && nextChar != "[" && nextChar != "," {
						n, _ := strconv.Atoi(line[j-1 : j+1])
						if n <= 9 && n+left > 9 {
							shiftLeft++
						}
						line = line[:j-1] + strconv.Itoa(n+left) + line[j+1:]
					} else {
						n, _ := strconv.Atoi(char)
						if n <= 9 && n+left > 9 {
							shiftLeft++
						}
						line = line[:j] + strconv.Itoa(n+left) + line[j+1:]
					}
					break
				}
			}
			line = line[:i+shiftLeft] + "0" + line[stop+1+shiftLeft:]
			return line, true
		}

	}

	return line, false
}

func split(line string) (string, bool) {
	for i, r := range line {
		char := string(r)
		if char == "]" || char == "[" || char == "," {
			continue
		}
		nextChar := string(line[i+1])
		if nextChar == "]" || nextChar == "[" || nextChar == "," {
			continue
		}

		n, _ := strconv.Atoi(line[i : i+2])
		if n < 10 {
			continue
		}
		number := float64(n) / 2
		left := int(math.Floor(number))
		right := int(math.Ceil(number))

		line = line[:i] + "[" + fmt.Sprintf("%d", left) + "," + fmt.Sprintf("%d", right) + "]" + line[i+2:]
		return line, true
	}

	return line, false
}
