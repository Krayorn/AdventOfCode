package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
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
}

type Zone struct {
	dist, h, l int
	index      int
}

type PriorityQueue []Zone

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].dist == pq[j].dist {
		if pq[i].h == pq[j].h {
			return pq[i].l > pq[j].l
		}
		return pq[i].h > pq[j].h
	}
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(Zone)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	// old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]

	return item
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

	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	queue.Push(Zone{0, 0, 0, 0})

	for queue.Len() > 0 {
		var zone Zone
		zone = heap.Pop(&queue).(Zone)
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
			queue.Push(Zone{distances[zone.h][zone.l], newHeight, newLength, 0})
		}
		// fmt.Println("PUSHED", queue)
	}

	return distances[tiles*height-1][tiles*length-1] - cave[0][0]
}
