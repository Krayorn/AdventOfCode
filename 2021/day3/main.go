package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, "\n")

	gammaRate := ""
	epsilonRate := ""
	for i := range values[0] {
		bitCountPositive := 0
		for _, v := range values {
			if v[i] == '1' {
				bitCountPositive++
			} else {
				bitCountPositive--
			}
		}

		if bitCountPositive > 0 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gammaRateValue, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateValue, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println(fmt.Sprintf("The power consumption of the submarine is %d", gammaRateValue*epsilonRateValue))

	oxygenGeneratorRatingValue, _ := strconv.ParseInt(recursiveFind(0, values, '1'), 2, 64)
	co2ScrubberRatingValue, _ := strconv.ParseInt(recursiveFind(0, values, '0'), 2, 64)

	fmt.Println(fmt.Sprintf("The power consumption of the submarine is %d", oxygenGeneratorRatingValue*co2ScrubberRatingValue))
}

func recursiveFind(index int, list []string, search rune) string {
	if len(list) == 1 {
		return list[0]
	}
	all0, all1 := make([]string, 0), make([]string, 0)
	for _, v := range list {
		if v[index] == '0' {
			all0 = append(all0, v)
		} else {
			all1 = append(all1, v)
		}
	}
	if search == '1' {
		if len(all1) < len(all0) {
			return recursiveFind(index+1, all0, search)
		}
		return recursiveFind(index+1, all1, search)
	}

	if len(all1) < len(all0) {
		return recursiveFind(index+1, all1, search)
	}
	return recursiveFind(index+1, all0, search)
}
