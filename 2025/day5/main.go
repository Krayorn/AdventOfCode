package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

type FreshnessRange struct {
	Min int
	Max int
}

func main() {
	parts := strings.Split(content, "\n\n")

	rangesStr := strings.Split(parts[0], "\n")
	ranges := make([]FreshnessRange, 0)

	sum := 0

	slices.Sort(rangesStr)
	for _, rangeStr := range rangesStr {
		boundaries := strings.Split(rangeStr, "-")
		min, _ := strconv.Atoi(boundaries[0])
		max, _ := strconv.Atoi(boundaries[1])

		r := FreshnessRange{min, max}

		found := false
		for i, existing := range ranges {
			if r.Min <= existing.Max && r.Min >= existing.Min {
				existing.Max = int(math.Max(float64(r.Max), float64(existing.Max)))
				found = true
			}
			if r.Max <= existing.Max && r.Max >= existing.Min {
				existing.Min = int(math.Min(float64(r.Min), float64(existing.Min)))
				found = true
			}
			if found {
				ranges[i] = existing
				break
			}
		}
		if !found {
			ranges = append(ranges, r)
		}
	}

	list := strings.Split(parts[1], "\n")

	for _, itemS := range list {
		item, _ := strconv.Atoi(itemS)

		for _, r := range ranges {
			if item <= r.Max && item >= r.Min {
				sum++
				break
			}
		}
	}

	sum2 := 0
	for _, r := range ranges {
		sum2 += r.Max - r.Min + 1
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
