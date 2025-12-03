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
	ranges := strings.Split(content, ",")
	sum := 0
	for _, r := range ranges {
		boundaries := strings.Split(r, "-")

		min, _ := strconv.Atoi(boundaries[0])
		max, _ := strconv.Atoi(boundaries[1])

		for i := min; i <= max; i++ {
			id := strconv.Itoa(i)
		DIVIDER:
			for divider := 1; divider < len(id); divider++ {
				if len(id)%divider != 0 {
					continue
				}
				chunks := make([]string, 0)

				for j := 0; j < len(id); j += divider {
					chunks = append(chunks, id[j:j+divider])
				}

				for _, chunk := range chunks[1:] {
					if chunk != chunks[0] {
						continue DIVIDER
					}
				}
				sum += i
				break DIVIDER
			}
		}
	}

	fmt.Println(sum)
}
