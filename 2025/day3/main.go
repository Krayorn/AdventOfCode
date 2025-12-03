package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	banks := strings.Split(content, "\n")

	n := 12

	sum := 0
	for _, bank := range banks {
		batteries := strings.Split(bank, "")
		bigPowers := make([]int, n)

		j := 0
		for i := len(batteries) - n; i < len(batteries); i++ {
			battery := batteries[i]
			power, _ := strconv.Atoi(battery)
			bigPowers[j] = power
			j++
		}
		fmt.Println(bigPowers)
		for i := len(batteries) - 1 - n; i >= 0; i-- {
			battery := batteries[i]
			power, _ := strconv.Atoi(battery)

			if power >= bigPowers[0] {
				slidingBigger := bigPowers[0]
				for j := 1; j < len(bigPowers); j++ {
					if slidingBigger >= bigPowers[j] {
						bigPowers[j], slidingBigger = slidingBigger, bigPowers[j]
					} else {
						break
					}
				}

				bigPowers[0] = power
			}
		}

		bankPowerStr := ""
		for _, power := range bigPowers {
			bankPowerStr += strconv.Itoa(power)
		}
		bankPower, _ := strconv.Atoi(bankPowerStr)
		fmt.Println(bank, bankPower)
		sum += bankPower
	}

	fmt.Println(sum)

}
