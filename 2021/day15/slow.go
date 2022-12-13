package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	cave := make([][]int, len(values))
	for i, l := range values {
		line := make([]int, len(l))
		risks := strings.Split(l, "")
		for j, v := range risks {
			n, _ := strconv.Atoi(v)
			line[j] = n
		}
		cave[i] = line
	}

	height := len(cave)
	length := len(cave[0])
	movesHeight := []int{-1, 0, 1, 0}
	movesLength := []int{0, 1, 0, -1}

	fmt.Println(solve(1, cave, height, length, movesHeight, movesLength))
	fmt.Println(solve(5, cave, height, length, movesHeight, movesLength))

	f, _ = os.Create("mem.prof")
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatalf("aaaa")
	}
	defer f.Close()
}

type Zone struct {
	dist, h, l int
}

func pop(slice []Zone) (Zone, []Zone) {
	return slice[0], slice[1:]
}

func push(slice []Zone, zone Zone) []Zone {
	slice = append(slice, zone)
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].dist < slice[j].dist {
			return true
		}
		if slice[i].h < slice[j].h {
			return true
		}
		return slice[i].l < slice[j].l
	})
	return slice
}

func solve(tiles int, cave [][]int, height, length int, movesHeight, movesLength []int) int {
	distances := make([][]int, height*tiles)
	for i := range distances {
		lineDistances := make([]int, length*tiles)
		distances[i] = lineDistances
		for j := range lineDistances {
			distances[i][j] = -1
		}
	}

	queue := make([]Zone, 1)
	queue[0] = Zone{0, 0, 0}

	for len(queue) > 0 {
		// fmt.Println(queue)
		var zone Zone
		zone, queue = pop(queue)
		if zone.h < 0 || zone.h >= tiles*height || zone.l < 0 || zone.l >= tiles*length {
			continue
		}

		value := cave[zone.h%height][zone.l%length] + int(math.Floor(float64(zone.h/height))) + int(math.Floor(float64(zone.l/length)))

		for value > 9 {
			value -= 9
		}
		// fmt.Println(value)
		newZoneDist := zone.dist + value

		if distances[zone.h][zone.l] == -1 || distances[zone.h][zone.l] > newZoneDist {
			distances[zone.h][zone.l] = newZoneDist
		} else {
			continue
		}

		if zone.h == tiles*height-1 && zone.l == tiles*length-1 {
			break
		}

		// fmt.Println("BEFORE", queue)
		for i := 0; i < 4; i++ {
			newHeight := zone.h + movesHeight[i]
			newLength := zone.l + movesLength[i]
			queue = push(queue, Zone{distances[zone.h][zone.l], newHeight, newLength})
		}
		// fmt.Println("PUSHED", queue)
	}

	return distances[tiles*height-1][tiles*length-1] - cave[0][0]
}
