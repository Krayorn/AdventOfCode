package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

type Rule struct {
	Before int
	After  int
}

func main() {
	sections := strings.Split(content, "\n\n")

	rulesStr := strings.Split(sections[0], "\n")

	rules := make([]Rule, len(rulesStr))
	for i, ruleStr := range rulesStr {
		parts := strings.Split(ruleStr, "|")
		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])

		rules[i] = Rule{Before: n1, After: n2}
	}

	updates := strings.Split(sections[1], "\n")

	sum := 0
	sum2 := 0
UPDATE:
	for _, update := range updates {
		parts := strings.Split(update, ",")
		numbers := make([]int, len(parts))
		for i, part := range parts {
			n, _ := strconv.Atoi(part)
			numbers[i] = n
		}

		for _, rule := range rules {
			foundLast := -1
			for i, n := range numbers {
				if rule.Before == n {
					if foundLast != -1 {
						tmp := numbers[foundLast]
						numbers[foundLast] = numbers[i]
						numbers[i] = tmp
					SWAP:
						for {
							for _, rule := range rules {
								foundLast := -1
								for i, n := range numbers {
									if rule.Before == n {
										if foundLast != -1 {
											tmp := numbers[foundLast]
											numbers[foundLast] = numbers[i]
											numbers[i] = tmp
											continue SWAP
										}
									}
									if rule.After == n {
										foundLast = i
									}
								}
							}
							break
						}
						sum2 += numbers[len(numbers)/2]
						continue UPDATE
					}
				}
				if rule.After == n {
					foundLast = i
				}
			}
		}

		sum += numbers[len(numbers)/2]
	}

	fmt.Println(sum)
	fmt.Println(sum2)

}
