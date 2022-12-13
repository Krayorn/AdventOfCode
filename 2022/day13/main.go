package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	parts := strings.Split(content, "\n\n")

	sum := 0

	for i := range parts {
		packets := strings.Split(parts[i], "\n")
		a := parse(packets[0])
		b := parse(packets[1])

		if compare(a, b) == -1 {
			sum += i + 1
		}
	}

	fmt.Println(sum)

	rawLines := strings.Split(content, "\n")
	var realLines []string
	for _, curr := range rawLines {
		if curr != "" {
			realLines = append(realLines, curr)
		}
	}

	realLines = append(realLines, "[[2]]", "[[6]]")

	sort.Slice(realLines, func(i, j int) bool {
		return compare(parse(realLines[i]), parse(realLines[j])) < 0
	})

	var twoDividerIdx int
	var sixDividerIdx int

	for i, curr := range realLines {
		if curr == "[[2]]" {
			twoDividerIdx = i + 1
		}
		if curr == "[[6]]" {
			sixDividerIdx = i + 1
		}
	}

	fmt.Printf("%d\n", twoDividerIdx*sixDividerIdx)
	fmt.Println(twoDividerIdx, sixDividerIdx)
}

func isList(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), "[]")
}

func compare(a any, b any) int {
	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	aIsList := isList(aType)
	bIsList := isList(bType)

	if !aIsList && !bIsList {
		// both numbers
		aNum := a.(float64)
		bNum := b.(float64)

		if aNum < bNum {
			return -1
		} else if aNum > bNum {
			return 1
		} else {
			return 0
		}
	}

	if aIsList && bIsList {
		// both lists
		aList := a.([]any)
		bList := b.([]any)

		aLen := len(aList)
		bLen := len(bList)

		var compareLen int
		if aLen < bLen {
			compareLen = aLen
		} else {
			compareLen = bLen
		}

		for i := 0; i < compareLen; i++ {
			c := compare(aList[i], bList[i])
			if c == -1 {
				return -1
			} else if c == 1 {
				return 1
			}
		}

		if aLen < bLen {
			return -1
		} else if aLen > bLen {
			return 1
		} else {
			return 0
		}
	}

	if aIsList && !bIsList {
		bNum := b.(float64)
		return compare(a, []any{bNum})
	}

	if !aIsList && bIsList {
		aNum := a.(float64)
		return compare([]any{aNum}, b)
	}

	panic("should not have made it this far")
}

func parse(input string) []any {
	var data []any
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}

	return data
}
