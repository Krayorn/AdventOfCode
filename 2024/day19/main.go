package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var content string

var cache map[string]int

func main() {
	lines := strings.Split(content, "\n")
	// re := regexp.MustCompile(fmt.Sprintf("^(%s)*$", strings.ReplaceAll(lines[0], ", ", "|")))

	tw := strings.Split(lines[0], ", ")
	availableTowels := make(map[string]bool)
	for _, t := range tw {
		availableTowels[t] = true
	}

	sum := 0

	for i := 2; i < len(lines); i++ {
		fmt.Println(lines[i])
		pattern := lines[i]
		cache = make(map[string]int)
		sum += arrange(pattern, availableTowels)

		// if re.MatchString(pattern) {
		// 	sum++
		// }
	}

	fmt.Println(sum)
}

func arrange(pattern string, availableTowels map[string]bool) int {
	if pattern == "" {
		return 1
	}

	count := 0
	for towel := range availableTowels {
		if strings.HasPrefix(pattern, towel) {
			if res, ok := cache[pattern[len(towel):]]; ok {
				count += res
				continue
			}

			res := arrange(pattern[len(towel):], availableTowels)
			cache[pattern[len(towel):]] = res
			count += res
		}
	}

	return count
}
