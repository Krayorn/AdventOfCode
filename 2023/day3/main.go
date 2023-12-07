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
	lines := strings.Split(content, "\n")
	sum := 0
	sum2 := 0
	grid := make([][]string, len(lines))

	for i, line := range lines {
		row := make([]string, len(line))
		for j, char := range line {
			row[j] = string(char)
		}
		grid[i] = row
	}
		
	gears := make(map[string][]int)

	buff := ""
	buffCloseToSymbol := false
	buffCloseToStar := make([]string, 0)
	for i, row := range grid {
		for j, char := range row {
			n, err := strconv.Atoi(char)
			if err == nil {
				buff += strconv.Itoa(n)
				fmt.Println("buff", buff)

				// avoid process for part 1, gonna use all process for part 2
				// if buffCloseToSymbol {
				// 	continue
				// }

//				CHECK:
				for k:= i-1;k <= i+1;k++{
					for l:= j-1;l <= j+1;l++ {
						if k < 0 || k >= len(grid) || l < 0 || l >= len(grid[0]) {
							continue
						}
						c := grid[k][l]
						fmt.Println(c)
						_, err := strconv.Atoi(c)
						if err != nil {
							if c != "." {
								if c == "*" {
									buffCloseToStar = append(buffCloseToStar, fmt.Sprintf("%d-%d", k, l))
								}

								buffCloseToSymbol = true
								//break CHECK don't break for part 2
							}
						}
					}
				}
				fmt.Println(buffCloseToSymbol)
			} else if buff != "" {
				if buffCloseToSymbol {
					nums, _ := strconv.Atoi(buff)
					sum += nums

					buffCloseToStar = removeDuplicate(buffCloseToStar)
					for _, gear := range buffCloseToStar {
						_, ok := gears[gear]
						if !ok {
							gears[gear] = make([]int, 0)
						}
						gears[gear] = append(gears[gear], nums)
					}
				}
				

				buff = ""
				buffCloseToSymbol = false
				buffCloseToStar = make([]string, 0)
			}
		}
		
		if buffCloseToSymbol {
			nums, _ := strconv.Atoi(buff)
			sum += nums

			buffCloseToStar = removeDuplicate(buffCloseToStar)
			for _, gear := range buffCloseToStar {
				_, ok := gears[gear]
				if !ok {
					gears[gear] = make([]int, 0)
				}
				gears[gear] = append(gears[gear], nums)
			}
		}

		buff = ""
		buffCloseToSymbol = false
		buffCloseToStar = make([]string, 0)
	}
	
	for _, gear := range gears {
		if len(gear) == 2 {
			sum2 += gear[0] * gear[1]
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func removeDuplicate[T string | int](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}