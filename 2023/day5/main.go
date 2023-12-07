package main

import (
	_ "embed"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

//go:embed input.txt
var content string

func main() {
	groups := strings.Split(content, "\n\n")

	seedsStack := strings.Split(groups[0], ": ")
	seedsStr := strings.Split(seedsStack[1], " ")

	seeds := make([]int, len(seedsStr))
	for i, seed := range seedsStr {
		n, _ := strconv.Atoi(seed)
		seeds[i] = n
	}

	groups = groups[1:]
	// for _, group := range groups {
	// 	rules := strings.Split(group, "\n")
	// 	rules = rules[1:]
		
	// 	SEED:
	// 	for i, seed := range seeds {
	// 		for _, rule := range rules {
	// 			nums := strings.Split(rule, " ")
	// 			destinationStart, _ := strconv.Atoi(nums[0])
	// 			sourceStart, _ := strconv.Atoi(nums[1])
	// 			rangeLength, _ := strconv.Atoi(nums[2])
				
	// 			if seed >= sourceStart && seed < sourceStart + rangeLength {
	// 				seeds[i] =  seed + (destinationStart - sourceStart) 
	// 				continue SEED
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(seeds)
	// }

	// minLoc := -1
	// for _, seed := range seeds {
	// 	if minLoc == -1 || seed < minLoc {
	// 		minLoc = seed
	// 	}
	// } 

	// fmt.Println(minLoc)

	slices.Reverse(groups)
	i := -1
	INVERSE:
	for  {

		
		i++
		if i % 100_000 == 0 {
			fmt.Println("------ TRYING FROM", i)
		}
		location := i
		GROUP:
		for _, group := range groups {
			rules := strings.Split(group, "\n")
			rules = rules[1:]
			
			for _, rule := range rules {
				nums := strings.Split(rule, " ")
				
				destinationStart, _ := strconv.Atoi(nums[0])
				sourceStart, _ := strconv.Atoi(nums[1])
				rangeLength, _ := strconv.Atoi(nums[2])
				
				if location >= destinationStart && location < destinationStart + rangeLength {
					location = location + (sourceStart - destinationStart) 
					continue GROUP
				}
			}
		}

		start := -1 
		for _, seed := range seeds {
			if start == -1 {
				start = seed
			} else {
				if location >= start && location < start + seed {
					fmt.Println(location)
					break INVERSE
				}
				start = -1
			}
		}
	}
	fmt.Println(i)
}
