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
	values := strings.Split(content, ",")

	fishsSchool := make(map[int]int, 8)
	for _, v := range values {
		value, _ := strconv.Atoi(v)
		if _, ok := fishsSchool[value]; !ok {
			fishsSchool[value] = 0
		}
		fishsSchool[value] += 1
	}

	fmt.Println(fmt.Sprintf("After 80 days, you'd get %d lanterfishs.", getFishSchoolSizeAfterDays(fishsSchool, 80)))
	fmt.Println(fmt.Sprintf("After 256 days, you'd get %d lanterfishs.", getFishSchoolSizeAfterDays(fishsSchool, 256)))
}

func getFishSchoolSizeAfterDays(fishsSchool map[int]int, days int) int {
	for i := 0; i < days; i++ {
		newSchool := make(map[int]int, 8)
		for daysLeft, counter := range fishsSchool {
			if daysLeft == 0 {
				newSchool[8] = counter
				if _, ok := newSchool[6]; !ok {
					newSchool[6] = 0
				}
				newSchool[6] += counter
				continue
			}

			if _, ok := newSchool[daysLeft-1]; !ok {
				newSchool[daysLeft-1] = 0
			}
			newSchool[daysLeft-1] += counter
		}
		fishsSchool = newSchool
	}

	count := 0
	for _, counter := range fishsSchool {
		count += counter
	}
	return count
}
