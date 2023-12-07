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
	values := strings.Split(content, "\n")

	sum := 0
	for _, v := range values {
		first, last := -1, -1
	
		buff := ""
		chars := strings.Split(v, "")
		for _, c := range chars {
			buff += c
			n, err := strconv.Atoi(c)
			if err == nil {
				if first == -1 {
					first = n
				}

				last = n
			}



			exist := strings.HasSuffix(buff, "one")
			if exist {
				if first == -1 {
					first = 1
				}

				last = 1
			} 

			exist = strings.HasSuffix(buff, "two")
			if exist {
				if first == -1 {
					first = 2
				}

				last = 2
			} 
			exist = strings.HasSuffix(buff, "three")
			if exist {
				if first == -1 {
					first = 3
				}

				last = 3
			} 
			exist = strings.HasSuffix(buff, "four")
			if exist {
				if first == -1 {
					first = 4
				}

				last = 4
			} 
			exist = strings.HasSuffix(buff, "five")
			if exist {
				if first == -1 {
					first = 5
				}

				last = 5
			} 
			exist = strings.HasSuffix(buff, "six")
			if exist {
				if first == -1 {
					first = 6
				}

				last = 6
			} 
			exist = strings.HasSuffix(buff, "seven")
			if exist {
				if first == -1 {
					first = 7
				}

				last = 7
			} 
			exist = strings.HasSuffix(buff, "eight")
			if exist {
				if first == -1 {
					first = 8
				}

				last = 8
			} 
			exist = strings.HasSuffix(buff, "nine")		
			if exist {
				if first == -1 {
					first = 9
				}

				last = 9
			} 
		}
		
		line := strconv.Itoa(first) + strconv.Itoa(last)
		fmt.Println(line)
		n, _ := strconv.Atoi(line)
		sum += n
	}

	fmt.Println(sum)
}