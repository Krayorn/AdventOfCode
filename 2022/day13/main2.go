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
	lines := strings.Split(content, "\n")

	lastPacket := ""
	pair := 1
	sum := 0

	for _, line := range lines {
		if line == "" {
			pair++
			lastPacket = ""

			// if pair == 7 {
			// 	break
			// }

			continue
		}

		if lastPacket == "" {
			lastPacket = line
			continue
		}

		if compareList(lastPacket, line) == 1 {
			fmt.Println("PAIR", pair, "IN ORDER")
			sum += pair
		}

	}

	fmt.Println(sum)

	allPackets := make([]string, 0, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		allPackets = append(allPackets, line)
	}
	allPackets = append(allPackets, "[[2]]", "[[6]]")

	sort.Slice(allPackets, func(i, j int) bool {
		return compareList(allPackets[i], allPackets[j]) > 0
	})

	var twoDividerIdx int
	var sixDividerIdx int

	for i, curr := range allPackets {
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

func compareList(left, right string) int {
	fmt.Println(left, right)
	left = left[1:]
	right = right[1:]
	fmt.Println("start", left, right)

	if left == "" && right == "" {
		return 0
	}

	if left == "" {
		return 1
	}

	if right == "" {
		return -1
	}

	for {
		if left[0] == ']' && right[0] != ']' {
			return 1
		}
		if right[0] == ']' && left[0] == ']' {
			return -1
		}
		if left[0] == ']' && right[0] == ']' {
			return 0
		}

		if left[0] == '[' && right[0] != '[' {
			toL1 := strings.Split(right, "]")
			toL2 := strings.Split(right, ",")
			if len(toL1[0]) > len(toL2[0]) {
				right = "[" + toL1[0] + "]"
				if len(toL1) > 1 {
					right = right + strings.Join(toL1[1:], "]")
				}
			} else {
				right = "[" + toL2[0] + "]"
				if len(toL2) > 1 {
					right = right + strings.Join(toL2[1:], ",")
				}
			}
			fmt.Println("right toList", left, right)
		}

		if left[0] != '[' && right[0] == '[' {
			toL1 := strings.Split(left, "]")
			toL2 := strings.Split(left, ",")
			if len(toL1[0]) > len(toL2[0]) {
				left = "[" + toL1[0] + "]"
				if len(toL1) > 1 {
					left = left + strings.Join(toL1[1:], "]")
				}
			} else {
				left = "[" + toL2[0] + "]"
				if len(toL2) > 1 {
					left = left + strings.Join(toL2[1:], ",")
				}
			}
			fmt.Println("left toList", left, right)
		}

		if left[0] == '[' && right[0] == '[' {
			l1 := strings.Split(left, "]")
			l2 := strings.Split(right, "]")
			fmt.Println("list", l1, l2)

			res := compareList(l1[0], l2[0])
			if res != 0 {
				return res
			}
			left = strings.TrimLeft(strings.Join(l1[1:], "]"), ",")
			right = strings.TrimLeft(strings.Join(l2[1:], "]"), ",")

			if left == "" && right == "" {
				return 0
			}

			fmt.Println("join list", left, right)
			continue
		}

		s1 := strings.Split(left, ",")
		s2 := strings.Split(right, ",")
		fmt.Println("split", s1, s2)

		n1, err1 := strconv.Atoi(s1[0])
		n2, err2 := strconv.Atoi(s2[0])
		fmt.Println("num", n1, err1, "-", n2, err2)

		if n1 < n2 {
			return 1
		} else if n1 > n2 {
			return -1
		}

		if len(s1) == 1 && len(s2) == 1 {
			return 0
		} else if len(s1) == 1 {
			return 1
		} else if len(s2) == 1 {
			return -1
		}

		left = strings.Join(s1[1:], ",")
		right = strings.Join(s2[1:], ",")

		fmt.Println("join split", left, right)
	}
}
