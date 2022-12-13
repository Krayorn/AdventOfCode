package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	template := values[0]
	rules := make(map[string]string, 0)

	for i := 2; i < len(values); i++ {
		split := strings.Split(values[i], " -> ")
		pair := strings.Split(split[0], "")
		rules[split[0]] = pair[0] + split[1] + pair[1]
	}

	pairCounts := make(map[string]int)
	for i := 1; i < len(template); i++ {
		pair := template[i-1 : i+1]
		if _, ok := pairCounts[pair]; !ok {
			pairCounts[pair] = 0
		}
		pairCounts[pair]++
	}

	for j := 0; j < 10; j++ {
		newpairCounts := make(map[string]int)
		for pair, count := range pairCounts {
			if replacement, ok := rules[pair]; ok {
				newpairCounts[replacement[0:2]] += count
				newpairCounts[replacement[1:3]] += count
			} else {
				newpairCounts[pair] += count
			}
		}
		pairCounts = newpairCounts
	}

	fmt.Println(pairCounts)

	counter := make(map[string]int)
	for pair, count := range pairCounts {
		for _, c := range pair {
			counter[string(c)] += count
		}
	}
	fmt.Println(counter)
	low, max := -1, -1
	counter[template[0:1]]++
	counter[template[len(template)-1:]]++
	for _, v := range counter {
		v = v / 2
		if low > v || low == -1 {
			low = v
		}
		if max < v || max == -1 {
			max = v
		}
	}
	fmt.Println(max - low)
}
