package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

//go:embed input.txt
var content string

func main() {
	games := strings.Split(content, "\n")

	sum := 0

	re := regexp.MustCompile(`^Game \d+: `)

	//GAME: 
	for _, game := range games {
		game = re.ReplaceAllString(game, "")
		plays := strings.Split(game, "; ")

		minBlue, minRed, minGreen := -1, -1, -1
		for _, play := range plays {
			cubes := strings.Split(play, ", ")
			fmt.Println(cubes)
			for _, cube := range cubes {
				sp := strings.Split(cube, " ")
				n, _ := strconv.Atoi(sp[0])
				
				if sp[1] == "green" && (minGreen < n || n == -1){
					minGreen = n
				}

				if sp[1] == "blue" && (minBlue < n || n == -1){
					minBlue = n
				}

				if sp[1] == "red" && (minRed < n || n == -1){
					minRed = n
				}

				// if sp[1] == "green" && n > 13{
				// 	continue GAME
				// }

				// if sp[1] == "blue" && n > 14{
				// 	continue GAME
				// }

				// if sp[1] == "red" && n > 12{
				// 	continue GAME
				// }
			}
		}
		
		fmt.Println(minRed * minGreen * minBlue, minRed, minGreen, minBlue)
		sum += minRed * minGreen * minBlue
	}
		
	fmt.Println(sum)
}