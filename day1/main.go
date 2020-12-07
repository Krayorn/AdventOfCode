package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	for _, v := range values {
		base, _ := strconv.Atoi(v)
		for _, v := range values {
			add1, _ := strconv.Atoi(v)
			for _, v := range values {
				add2, _ := strconv.Atoi(v)
				if base+add1+add2 == 2020 {
					fmt.Println(base * add1 * add2)
				}
			}
		}
	}
}
