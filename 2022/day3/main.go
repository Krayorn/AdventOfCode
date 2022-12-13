package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var content string

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func main() {
	values := strings.Split(content, "\n")

	sum := 0
	rumsack := make(map[string]int)

	for i, v := range values {
		items := strings.Split(v, "")

		elfNumber := i % 3

		if elfNumber == 0 {
			rumsack = make(map[string]int)
		}
		for _, item := range items {
			if _, ok := rumsack[item]; !ok && elfNumber == 0 {
				rumsack[item] = 1
				continue
			}

			if val, ok := rumsack[item]; ok && elfNumber != 0 {
				if val == 1 && elfNumber == 1 {
					rumsack[item] = 2
				}
				if val == 2 && elfNumber == 2 {
					rumsack[item] = 3
					p := int(strings.ToLower(item)[0]) - 96
					if unicode.IsUpper(rune(item[0])) {
						p += 26
					}
					sum += p
					break
				}
			}

		}
	}
	fmt.Println(sum)
}

// if i+1 > half {
// 	if _, ok := rumsack[item]; ok {
// p := int(strings.ToLower(item)[0]) - 96
// if unicode.IsUpper(rune(item[0])) {
// 	p += 26
// }
// sum += p
// 		break
// 	}
// 	continue
// } else {
// 	if _, ok := rumsack[item]; !ok {
// 		rumsack[item] = 1
// 	} else {
// 		rumsack[item] += 1
// 	}
// }
