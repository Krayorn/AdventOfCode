package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	name string
	dist int
	arcs map[string]int
}

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

	bigCave := make([][]int, len(cave)*5)
	for i := range bigCave {
		bigLine := make([]int, len(cave[0])*5)
		bigCave[i] = bigLine
		for j := range bigLine {
			risk := 0
			if i-len(cave) >= 0 {
				risk = bigCave[i-len(cave)][j] + 1
			} else if j-len(cave[i]) >= 0 {
				risk = bigCave[i][j-len(cave[0])] + 1
			} else {
				risk = cave[i][j]
			}
			if risk == 10 {
				risk = 1
			}
			bigCave[i][j] = risk
		}

	}
	// printCave(cave)
	// printCave(bigCave)
	verticles := make(map[string]Vertex, 0)

	for i, line := range bigCave {
		for j := range line {
			verticle := Vertex{name: strconv.Itoa(j) + "-" + strconv.Itoa(i), arcs: make(map[string]int)}
			if i+1 < len(bigCave) {
				verticle.arcs[strconv.Itoa(j)+"-"+strconv.Itoa(i+1)] = bigCave[j][i+1]
			}
			if i-1 >= 0 {
				verticle.arcs[strconv.Itoa(j)+"-"+strconv.Itoa(i-1)] = bigCave[j][i-1]
			}
			if j+1 < len(line) {
				verticle.arcs[strconv.Itoa(j+1)+"-"+strconv.Itoa(i)] = bigCave[j+1][i]
			}
			if j-1 >= 0 {
				verticle.arcs[strconv.Itoa(j-1)+"-"+strconv.Itoa(i)] = bigCave[j-1][i]
			}
			verticles[verticle.name] = verticle
		}
	}

	fmt.Println(short("0-0", strconv.Itoa(len(bigCave)-1)+"-"+strconv.Itoa(len(bigCave[0])-1), verticles))
}

func short(src string, dest string, verticles map[string]Vertex) int {
	nodes := make([]Vertex, 0)
	for _, node := range verticles {
		nodes = append(nodes, node)
	}

	shortestPath := make(map[string]int)

	max := math.MaxInt64
	for _, node := range nodes {
		shortestPath[node.name] = max
	}
	shortestPath[src] = 0

	for len(nodes) > 0 {
		minNode := nodes[0]
		minI := 0
		for i, node := range nodes {
			if shortestPath[node.name] < shortestPath[minNode.name] {
				minNode = node
				minI = i
			}
		}

		neighbors := minNode.arcs
		for neighborName, distance := range neighbors {
			tentValue := shortestPath[minNode.name] + distance
			if tentValue < shortestPath[neighborName] {
				shortestPath[neighborName] = tentValue
			}
		}

		nodes = append(nodes[:minI], nodes[minI+1:]...)
	}

	return shortestPath[dest]
}

func printCave(cave [][]int) {
	for y := range cave {
		for x := range cave[y] {
			fmt.Print(cave[y][x])
		}
		fmt.Println("")
	}
}
