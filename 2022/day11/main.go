package main

import (
	_ "embed"
	"fmt"
	"sort"
)

//go:embed input.txt
var content string

type OperationFunc func(int) int

type Monkey struct {
	Number    int
	Activity  int
	Items     []int
	Operation OperationFunc
	Divide    int
	Success   int
	Failure   int
}

func main() {
	monkeysTest := []Monkey{
		{
			0,
			0,
			[]int{79, 98},
			func(old int) int {
				return old * 19
			},
			23,
			2,
			3,
		},
		{
			1,
			0,
			[]int{54, 65, 75, 74},
			func(old int) int {
				return old + 6
			},
			19,
			2,
			0,
		},
		{
			2,
			0,
			[]int{79, 60, 97},
			func(old int) int {
				return old * old
			},
			13,
			1,
			3,
		},
		{
			3,
			0,
			[]int{74},
			func(old int) int {
				return old + 3
			},
			17,
			0,
			1,
		},
	}
	fmt.Println(monkeysTest)
	monkeys := []Monkey{
		{
			0,
			0,
			[]int{72, 97},
			func(old int) int {
				return old * 13
			},
			19,
			5,
			6,
		},
		{
			1,
			0,
			[]int{55, 70, 90, 74, 95},
			func(old int) int {
				return old * old
			},
			7,
			5,
			0,
		},
		{
			2,
			0,
			[]int{74, 97, 66, 57},
			func(old int) int {
				return old + 6
			},
			17,
			1,
			0,
		},
		{
			3,
			0,
			[]int{86, 54, 53},
			func(old int) int {
				return old + 2
			},
			13,
			1,
			2,
		},
		{
			4,
			0,
			[]int{50, 65, 78, 50, 62, 99},
			func(old int) int {
				return old + 3
			},
			11,
			3,
			7,
		},
		{
			5,
			0,
			[]int{90},
			func(old int) int {
				return old + 4
			},
			2,
			4,
			6,
		},
		{
			6,
			0,
			[]int{88, 92, 63, 94, 96, 82, 53, 53},
			func(old int) int {
				return old + 8
			},
			5,
			4,
			7,
		},
		{
			7,
			0,
			[]int{70, 60, 71, 69, 77, 70, 98},
			func(old int) int {
				return old * 7
			},
			3,
			2,
			3,
		},
	}
	fmt.Println(monkeys)
	//monkeys = monkeysTest

	lcm := prod(23, 19, 13, 17)
	lcm = prod(19, 7, 17, 13, 11, 2, 5, 3)

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkeys[monkey.Number].Activity += 1
				//stress := monkey.Operation(item) / 3
				stress := monkey.Operation(item) % lcm
				if stress%monkey.Divide == 0 {
					monkeys[monkey.Success].Items = append(monkeys[monkey.Success].Items, stress)
				} else {
					monkeys[monkey.Failure].Items = append(monkeys[monkey.Failure].Items, stress)
				}
			}
			monkeys[monkey.Number].Items = []int{}
		}
	}

	fmt.Println(monkeys)

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Activity > monkeys[j].Activity
	})
	fmt.Println(monkeys[0].Activity * monkeys[1].Activity)
}

func prod(integers ...int) int {  
	result := 1

	for i := 0; i < len(integers); i++ {
		result *= integers[i]
	}

	return result
}
