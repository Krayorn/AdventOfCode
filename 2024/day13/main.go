package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

//go:embed input.txt
var content string

func main() {
	machinesInput := strings.Split(content, "\n\n")

	rButton, _ := regexp.Compile(`X\+(\d+), Y\+(\d+)`)
	rPrize, _ := regexp.Compile(`X=(\d+), Y=(\d+)`)

	sum := 0

	for _, machineInput := range machinesInput {
		lines := strings.Split(machineInput, "\n")
		matchesA := rButton.FindStringSubmatch(lines[0])
		matchesB := rButton.FindStringSubmatch(lines[1])
		matchesPrize := rPrize.FindStringSubmatch(lines[2])

		aX, _ := strconv.Atoi(matchesA[1])
		aY, _ := strconv.Atoi(matchesA[2])

		bX, _ := strconv.Atoi(matchesB[1])
		bY, _ := strconv.Atoi(matchesB[2])

		prizeX, _ := strconv.Atoi(matchesPrize[1])
		prizeY, _ := strconv.Atoi(matchesPrize[2])

		A := mat.NewDense(2, 2, []float64{float64(aX), float64(bX), float64(aY), float64(bY)})
		b := mat.NewVecDense(2, []float64{float64(prizeX + 10000000000000), float64(prizeY + 10000000000000)})

		var x mat.VecDense
		if err := x.SolveVec(A, b); err != nil {
			fmt.Println(err)
			continue
		}

		aCount := x.RawVector().Data[0]
		bCount := x.RawVector().Data[1]
		if math.Abs(math.Round(aCount)-aCount) <= 0.001 && math.Abs(math.Round(bCount)-bCount) <= 0.001 {
			sum += int(math.Round(aCount))*3 + int(math.Round(bCount))
		}
	}

	fmt.Println(sum)

}
