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
	content := strings.Split(content, "\n")

	piles := make(map[int][]string)

	//nbOfCol := len(string(content[0])) / 3
	nbOfCol := 9

	for i := 1; i <= nbOfCol; i++ {
		piles[i] = make([]string, 0)
	}

	instructions := false
	count := 0
	for _, v := range content {
		line := string(v)
		if line == "" {
			continue
		}
		if instructions == false && line[1] == '1' {
			instructions = true
			continue
		}

		if instructions {
			breaking := strings.Split(line, " ")
			count, _ := strconv.Atoi(string(breaking[1]))
			from, _ := strconv.Atoi(string(breaking[3]))
			to, _ := strconv.Atoi(string(breaking[5]))

			//fmt.Println(count, from, to)
			//fmt.Println(piles)

			// for i := 0; i < count; i++ {
			// 	if len(piles[from]) <= 0 {
			// 		continue
			// 	}
			// 	//				fmt.Println(piles[from], piles[to])
			// 	var x string
			// 	x, piles[from] = piles[from][len(piles[from])-1], piles[from][:len(piles[from])-1]
			// 	piles[to] = append(piles[to], x)
			// }

			piles[to] = append(piles[to], piles[from][len(piles[from])-count:]...)
			piles[from] = piles[from][:len(piles[from])-count]

			continue
		}

		for i := 1; i <= nbOfCol; i++ {
			if line[1+(i-1)*4] != 32 {
				count++
				piles[i] = append([]string{string(line[1+(i-1)*4])}, piles[i]...)
			}
		}
	}

	rez := ""
	for i := 1; i <= nbOfCol; i++ {
		rez = rez + piles[i][len(piles[i])-1]
	}
	fmt.Println(rez)

}
