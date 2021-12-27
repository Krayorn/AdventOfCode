package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	count, count2 := 0, 0
	for _, v := range values {
		line := strings.Split(v, " | ")
		outputs := strings.Split(line[1], " ")
		mappingSignals := strings.Split(line[0], " ")
		count2 += decode(mappingSignals, outputs)
		for _, output := range outputs {
			if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7 {
				count++
			}
		}
	}

	fmt.Println(fmt.Sprintf("The digits 1, 4, 7, 8 appears %d times in the outputs.", count))
	fmt.Println(fmt.Sprintf("If you add up all of the output values, you get %d.", count2))
}

func decode(mappingSignals []string, outputs []string) int {
	possibilities := make(map[string][]string, 9)

	sort.Slice(mappingSignals, func(p, q int) bool {
		return len(mappingSignals[p]) < len(mappingSignals[q])
	})

	for _, signal := range mappingSignals {
		if len(signal) == 2 {
			possibilities["c"] = strings.Split(signal, "")
			possibilities["f"] = strings.Split(signal, "")
		}
		if len(signal) == 3 {
			for _, c := range signal {
				if !isInArray(possibilities["c"], string(c)) {
					possibilities["a"] = make([]string, 0)
					possibilities["a"] = append(possibilities["a"], string(c))
					break
				}
			}
		}
		if len(signal) == 4 {
			possibilities["b"] = make([]string, 0)
			possibilities["d"] = make([]string, 0)
			for _, c := range signal {
				if !isInArray(possibilities["c"], string(c)) {
					possibilities["b"] = append(possibilities["b"], string(c))
					possibilities["d"] = append(possibilities["d"], string(c))
				}
			}
		}

		if len(signal) == 7 {
			possibilities["e"] = make([]string, 0)
			possibilities["g"] = make([]string, 0)
			for _, c := range signal {
				if !isInArray(possibilities["c"], string(c)) && !isInArray(possibilities["a"], string(c)) && !isInArray(possibilities["b"], string(c)) {
					possibilities["e"] = append(possibilities["e"], string(c))
					possibilities["g"] = append(possibilities["g"], string(c))
				}
			}
		}
	}

	for _, signal := range mappingSignals {
		if len(signal) == 6 {
			hasCF := 0
			hasBD := 0
			hasEG := 0
			for _, c := range signal {
				if isInArray(possibilities["c"], string(c)) {
					hasCF++
				} else if isInArray(possibilities["f"], string(c)) {
					hasCF++
				}

				if isInArray(possibilities["b"], string(c)) {
					hasBD++
				} else if isInArray(possibilities["d"], string(c)) {
					hasBD++
				}

				if isInArray(possibilities["e"], string(c)) {
					hasEG++
				} else if isInArray(possibilities["g"], string(c)) {
					hasEG++
				}
			}

			if hasCF == 2 && hasEG == 2 {
				values := possibilities["d"]
				for _, c := range values {
					if !isInArray(strings.Split(signal, ""), c) {
						newPos := make([]string, 0)
						possibilities["d"] = append(newPos, c)
					}
				}
				values = possibilities["b"]
				for _, c := range values {
					if !isInArray(possibilities["d"], c) {
						newPos := make([]string, 0)
						possibilities["b"] = append(newPos, c)
					}
				}
			}

			if hasBD == 2 && hasEG == 2 {
				values := possibilities["c"]
				for _, c := range values {
					if !isInArray(strings.Split(signal, ""), c) {
						newPos := make([]string, 0)
						possibilities["c"] = append(newPos, c)
					}
				}
				values = possibilities["f"]
				for _, c := range values {
					if !isInArray(possibilities["c"], c) {
						newPos := make([]string, 0)
						possibilities["f"] = append(newPos, c)
					}
				}
			}

			if hasBD == 2 && hasCF == 2 {
				values := possibilities["e"]
				for _, c := range values {
					if !isInArray(strings.Split(signal, ""), c) {
						newPos := make([]string, 0)
						possibilities["e"] = append(newPos, c)
					}
				}
				values = possibilities["g"]
				for _, c := range values {
					if !isInArray(possibilities["e"], c) {
						newPos := make([]string, 0)
						possibilities["g"] = append(newPos, c)
					}
				}
			}
		}
	}

	match := map[int]string{
		0: possibilities["a"][0] + possibilities["b"][0] + possibilities["c"][0] + possibilities["e"][0] + possibilities["f"][0] + possibilities["g"][0],
		1: possibilities["c"][0] + possibilities["f"][0],
		2: possibilities["a"][0] + possibilities["c"][0] + possibilities["d"][0] + possibilities["e"][0] + possibilities["g"][0],
		3: possibilities["a"][0] + possibilities["c"][0] + possibilities["d"][0] + possibilities["f"][0] + possibilities["g"][0],
		4: possibilities["b"][0] + possibilities["c"][0] + possibilities["d"][0] + possibilities["f"][0],
		5: possibilities["a"][0] + possibilities["b"][0] + possibilities["d"][0] + possibilities["f"][0] + possibilities["g"][0],
		6: possibilities["a"][0] + possibilities["b"][0] + possibilities["d"][0] + possibilities["e"][0] + possibilities["f"][0] + possibilities["g"][0],
		7: possibilities["a"][0] + possibilities["c"][0] + possibilities["f"][0],
		8: possibilities["a"][0] + possibilities["b"][0] + possibilities["c"][0] + possibilities["d"][0] + possibilities["e"][0] + possibilities["f"][0] + possibilities["g"][0],
		9: possibilities["a"][0] + possibilities["b"][0] + possibilities["c"][0] + possibilities["d"][0] + possibilities["f"][0] + possibilities["g"][0],
	}

	count := ""
	for _, output := range outputs {
		for id, val := range match {
			if alphasort([]rune(val)) == alphasort([]rune(output)) {
				count += strconv.Itoa(id)
			}
		}
	}

	c, _ := strconv.Atoi(count)
	return c
}

func alphasort(str []rune) string {
	for x := range str {
		y := x + 1
		for y = range str {
			if str[x] < str[y] {
				str[x], str[y] = str[y], str[x]
			}
		}
	}
	return string(str)
}

func isInArray(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
