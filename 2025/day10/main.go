package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

//go:embed input.txt
var content string

type State struct {
	Mask  int
	Count int
}

const maxPresses = 1000

func main() {
	lines := strings.Split(content, "\n")

	sum := 0
	sum2 := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		switchesN := len(parts[0]) - 2
		var wantedMask, mask int

		for i, c := range parts[0][1 : switchesN+1] {
			if c == '#' {
				wantedMask |= (1 << i)
			}
		}

		wantedJolts := make([]int, switchesN)
		jolts := make([]int, switchesN)

		joltage := strings.Split(strings.Trim(parts[len(parts)-1], "{}"), ",")
		for i, joltString := range joltage {
			n, _ := strconv.Atoi(joltString)
			wantedJolts[i] = n
		}

		operations := make([]int, 0)
		operationsNumbers := make([][]int, 0)
		for _, button := range parts[1 : len(parts)-1] {
			numbers := strings.Split(strings.Trim(button, "()"), ",")
			buttonMask := 0
			operationNumbers := make([]int, 0)
			for _, nString := range numbers {
				n, _ := strconv.Atoi(nString)
				operationNumbers = append(operationNumbers, n)
				buttonMask |= 1 << n
			}

			operationsNumbers = append(operationsNumbers, operationNumbers)
			operations = append(operations, buttonMask)
		}

		fmt.Println("Want -", fmt.Sprintf("%0*b", switchesN, wantedMask), wantedJolts)
		fmt.Println("Has -", fmt.Sprintf("%0*b", switchesN, mask), jolts)
		fmt.Println("Operations => ")
		for _, action := range operations {
			fmt.Println("Has -", fmt.Sprintf("%0*b", switchesN, action))
		}
		fmt.Println()
		fmt.Println("----")
		fmt.Println()

		queue := make([]State, 0)
		queue = append(queue, State{mask, 0})
		for true {
			item := queue[0]
			queue = queue[1:]

			if item.Mask == wantedMask {
				sum += item.Count
				break
			}

			for _, action := range operations {
				queue = append(queue, State{item.Mask ^ action, item.Count + 1})
			}
		}

		if slices.Equal(wantedJolts, jolts) {
			continue
		}

		numButtons := len(operationsNumbers)
		numJoltages := len(wantedJolts)

		lp := golp.NewLP(0, numButtons)
		lp.SetVerboseLevel(golp.NEUTRAL)

		objectiveCoeffs := make([]float64, numButtons)
		for i := range numButtons {
			objectiveCoeffs[i] = 1.0
		}
		lp.SetObjFn(objectiveCoeffs)

		for i := range numButtons {
			lp.SetInt(i, true)
			lp.SetBounds(i, 0.0, float64(maxPresses))
		}

		for i := 0; i < numJoltages; i++ {
			var entries []golp.Entry
			for j, btn := range operationsNumbers {
				if slices.Contains(btn, i) {
					entries = append(entries, golp.Entry{Col: j, Val: 1.0})
				}
			}
			targetValue := float64(wantedJolts[i])
			if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
				panic(err)
			}
		}

		status := lp.Solve()

		if status != golp.OPTIMAL {
			continue
		}

		solution := lp.Variables()
		totalPresses := 0
		for _, val := range solution {
			totalPresses += int(val + 0.5)
		}

		sum2 += totalPresses
	}

	fmt.Println(sum, sum2)
}
