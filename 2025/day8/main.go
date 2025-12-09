package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

const NUMBER_DIST = 1000

type Box struct {
	X, Y, Z int
}

type Distance struct {
	Dist float64
	A, B Box
}

func main() {
	lines := strings.Split(content, "\n")

	boxes := make([]Box, 0)

	distances := make([]Distance, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		newBox := Box{x, y, z}

		for _, existingBox := range boxes {
			distance := math.Sqrt(math.Pow(float64(newBox.X-existingBox.X), 2) + math.Pow(float64(newBox.Y-existingBox.Y), 2) + math.Pow(float64(newBox.Z-existingBox.Z), 2))

			distances = append(distances, Distance{distance, newBox, existingBox})
		}

		boxes = append(boxes, newBox)
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		if a.Dist-b.Dist < 0 {
			return -1
		}
		return 1
	})

	circuits := make([][]Box, 0)

	for _, distance := range distances {
		aInCircuit := -1
		bInCircuit := -1
		for i, circuit := range circuits {
			for _, box := range circuit {
				if box.Equal(distance.A) {
					aInCircuit = i
				}
				if box.Equal(distance.B) {
					bInCircuit = i
				}
			}
		}

		if aInCircuit != -1 && bInCircuit != -1 {
			if aInCircuit == bInCircuit {
				continue
			}
			circuits[aInCircuit] = append(circuits[aInCircuit], circuits[bInCircuit]...)
			circuits = slices.Delete(circuits, bInCircuit, bInCircuit+1)
		} else if aInCircuit != -1 {
			circuits[aInCircuit] = append(circuits[aInCircuit], distance.B)
		} else if bInCircuit != -1 {
			circuits[bInCircuit] = append(circuits[bInCircuit], distance.A)
		} else {
			circuits = append(circuits, []Box{distance.A, distance.B})
		}

		if len(circuits) == 1 && len(circuits[0]) == len(boxes) {
			fmt.Println(distance.A.X * distance.B.X)
			os.Exit(1)
		}
	}

	slices.SortFunc(circuits, func(a, b []Box) int {
		if len(b)-len(a) < 0 {
			return -1
		}
		return 1
	})

	fmt.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))

	// fmt.Println("------")
	// for _, circuit := range circuits {
	// 	fmt.Println(len(circuit), circuit)
	// }

}

func (me Box) Equal(other Box) bool {
	return me.X == other.X && me.Y == other.Y && me.Z == other.Z
}
