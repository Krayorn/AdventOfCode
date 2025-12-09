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
	lines := strings.Split(content, "\n")

	points := make([][]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, []int{x, y})
	}

	edges := make([][]int, 0)
	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		p2 := points[i+1]
		edges = append(edges, []int{int(math.Min(float64(p1[0]), float64(p2[0]))), int(math.Min(float64(p1[1]), float64(p2[1]))), int(math.Max(float64(p1[0]), float64(p2[0]))), int(math.Max(float64(p1[1]), float64(p2[1])))})
	}

	p1 := points[0]
	p2 := points[len(points)-1]
	edges = append(edges, []int{int(math.Min(float64(p1[0]), float64(p2[0]))), int(math.Min(float64(p1[1]), float64(p2[1]))), int(math.Max(float64(p1[0]), float64(p2[0]))), int(math.Max(float64(p1[1]), float64(p2[1])))})

	area := -1
	areaConstraints := -1

	for i := 0; i < len(points); i++ {
	COMBINATIONS:
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			xDiff := math.Abs(float64(p2[0]-p1[0])) + 1
			yDiff := math.Abs(float64(p2[1]-p1[1])) + 1

			newArea := int(xDiff * yDiff)
			if newArea > area {
				area = newArea
			}
			if newArea > areaConstraints {
				minX, minY, maxX, maxY := int(math.Min(float64(p1[0]), float64(p2[0]))), int(math.Min(float64(p1[1]), float64(p2[1]))), int(math.Max(float64(p1[0]), float64(p2[0]))), int(math.Max(float64(p1[1]), float64(p2[1])))

				for _, edge := range edges {
					if minX < edge[2] && maxX > edge[0] && minY < edge[3] && maxY > edge[1] {
						continue COMBINATIONS
					}
				}
				areaConstraints = newArea
			}
		}
	}

	fmt.Println(area)
	fmt.Println(areaConstraints)
}
