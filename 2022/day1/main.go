package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, "\n")

	sum := 0
	carry := make([]int, 0)
	for _, v := range values {
		if v == "" {
			carry = append(carry, sum)
			sum = 0
			continue
		}
		n, _ := strconv.Atoi(v)
		sum = sum + n
	}

	sort.Ints(carry)

	fmt.Println(carry[len(carry)-1])
	fmt.Println(carry[len(carry)-1] + carry[len(carry)-2] + carry[len(carry)-3])
}
