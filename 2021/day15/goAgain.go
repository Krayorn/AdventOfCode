package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var contents string

func main() {
	contents = strings.TrimSpace(contents)
	lines := strings.Split(contents, "\n")

	grid := make(map[xy]int)
	var length, height = len(lines), len(lines[0])
	for x, line := range lines {
		for y, v := range numbers(line) {
			grid[xy{x, y}] = v
		}
	}

	start := xy{0, 0}
	target := xy{length*5 - 1, height*5 - 1}

	fmt.Println(start, target)

	risk := func(pos xy) int {
		og := xy{
			pos.x % length,
			pos.y % height,
		}

		mul := pos.x/length + pos.y/height
		risk := grid[og] + mul
		if risk > 9 {
			risk = risk - 9
		}

		return risk
	}

	shortestAt := make(map[xy]int)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(qi{pos: start, riskLevel: 0})

	for pq.Len() > 0 {
		head := heap.Pop(&pq).(qi)
		for i := 0; i < 4; i++ {
			next := xy{
				head.pos.x + dx[i],
				head.pos.y + dy[i],
			}

			if next.x >= length*5 || next.x < 0 || next.y >= height*5 || next.y < 0 {
				continue
			}

			nextRisk := head.riskLevel + risk(next)
			if sAt, ok := shortestAt[next]; ok && sAt <= nextRisk {
				continue
			} else {
				shortestAt[next] = nextRisk
			}

			pq.Push(qi{
				pos:       next,
				riskLevel: nextRisk,
			})
		}
	}

	fmt.Println(shortestAt[target])
}

type qi struct {
	pos       xy
	riskLevel int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.

type PriorityQueue []qi

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].riskLevel < pq[j].riskLevel
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(qi)
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

var dx = [4]int{0, 0, -1, 1}
var dy = [4]int{-1, 1, 0, 0}

func numbers(in string) []int {
	var list []int
	for _, word := range strings.Split(in, "") {
		list = append(list, atoi(word))
	}
	return list
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

type xy struct {
	x int
	y int
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}
