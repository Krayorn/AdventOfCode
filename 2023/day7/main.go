package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sort"
	"strconv"
)

//go:embed input.txt
var content string

func main() {
	hands := strings.Split(content, "\n")

	five := make([]string, 0)
	four := make([]string, 0)
	house := make([]string, 0)
	three := make([]string, 0)
	twotwo := make([]string, 0)
	two := make([]string, 0)
	worse := make([]string, 0)

	HAND:
	for _, hand := range hands {
		split := strings.Split(hand, " ")
		coll := make(map[rune]int)
		
		for _, r := range split[0] {
			if _, ok := coll[r]; !ok {
				coll[r] = 0
			}
			coll[r] += 1
		}
		
		JCount, _ := coll['J']

		hasThree := false
		hasTwo := false
		for k, count := range coll {
			if count == 5 || JCount == 4 || (k != 'J' && count + JCount == 5) {
				five = append(five, hand)
				continue HAND
			}
			if count == 4 || (k != 'J' && count + JCount == 4) {
				four = append(four, hand)
				continue HAND
			}
			if count == 3 || (k != 'J' && count + JCount == 3) {
				if hasTwo {
					house = append(house, hand)
					continue HAND
				}
				if hasThree && JCount == 1{
					house = append(house, hand)
					continue HAND
				}
				hasThree = true
				continue
			}
			if count == 2 && k != 'J' {
				if hasTwo && JCount == 1 { 
					house = append(house, hand)
				}
				if hasTwo {
					twotwo = append(twotwo, hand)
					continue HAND
				}
				if hasThree {
					house = append(house, hand)
					continue HAND
				}
				hasTwo = true
			}
		}

		if hasThree {
			three = append(three, hand)
			continue HAND
		}

		if hasTwo {
			two = append(two, hand)
			continue HAND
		}

		if JCount == 1 {
			two = append(two, hand)
			continue HAND
		}

		worse = append(worse, hand)
		continue HAND
	}


	totalWinnings := 0
	handsCount := 0

	sort.SliceStable(worse, func(i, j int) bool {
		for ind, r := range worse[i] {
			i1 := strings.Index("AKQT98765432J", string(r))
			i2 := strings.Index("AKQT98765432J", string(worse[j][ind]))

			if i1 == i2 {
				continue
			}
			//fmt.Println(i1, i2, string(r), string(three[j][ind]))
			return i1 > i2
		}
		fmt.Println("broken")
		return false
	})

	sort.SliceStable(two, func(i, j int) bool {
		for ind, r := range two[i] {
			i1 := strings.Index("AKQT98765432J", string(r))
			i2 := strings.Index("AKQT98765432J", string(two[j][ind]))

			if i1 == i2 {
				continue
			}
			//fmt.Println(i1, i2, string(r), string(three[j][ind]))
			return i1 > i2
		}
		fmt.Println("broken")
		return false
	})

	sort.SliceStable(twotwo, func(i, j int) bool {
		for ind, r := range twotwo[i] {
			i1 := strings.Index("AKQT98765432J", string(r))
			i2 := strings.Index("AKQT98765432J", string(twotwo[j][ind]))

			if i1 == i2 {
				continue
			}
			//fmt.Println(i1, i2, string(r), string(three[j][ind]))
			return i1 > i2
		}
		fmt.Println("broken")
		return false
	})

	sort.SliceStable(three, func(i, j int) bool {
		for ind, r := range three[i] {
			i1 := strings.Index("AKQT98765432J", string(r))
			i2 := strings.Index("AKQT98765432J", string(three[j][ind]))

			if i1 == i2 {
				continue
			}
			//fmt.Println(i1, i2, string(r), string(three[j][ind]))
			return i1 > i2
		}
		fmt.Println("broken")
		return false
	})

	sort.SliceStable(house, func(i, j int) bool {
		for ind, r := range house[i] {
			i1 := strings.Index("AKQT98765432J", string(r))
			i2 := strings.Index("AKQT98765432J", string(house[j][ind]))

			if i1 == i2 {
				continue
			}
			//fmt.Println(i1, i2, string(r), string(three[j][ind]))
			return i1 > i2
		}
		fmt.Println("broken")
		return false
	})

	sort.SliceStable(four, func(i, j int) bool {
		for ind, r := range four[i] {
			i1 := strings.Index("AKQT98765432J", string(r))
			i2 := strings.Index("AKQT98765432J", string(four[j][ind]))

			if i1 == i2 {
				continue
			}
			//fmt.Println(i1, i2, string(r), string(three[j][ind]))
			return i1 > i2
		}
		fmt.Println("broken")
		return false
	})

	sort.SliceStable(five, func(i, j int) bool {
		for ind, r := range five[i] {
			i1 := strings.Index("AKQT98765432J", string(r))
			i2 := strings.Index("AKQT98765432J", string(five[j][ind]))

			if i1 == i2 {
				continue
			}
			//fmt.Println(i1, i2, string(r), string(three[j][ind]))
			return i1 > i2
		}
		fmt.Println("broken")
		return false
	})

	fmt.Println(five)
	fmt.Println(four)
	fmt.Println(house)
	fmt.Println(three)
	fmt.Println(twotwo)
	fmt.Println(two)
	fmt.Println(worse)

	for i:= 0;i <len(worse) ;i++{
		split := strings.Split(worse[i], " ")
		fmt.Print(split[0], " ")
	}
	for i:= 0;i <len(two) ;i++{
		split := strings.Split(two[i], " ")
		fmt.Print(split[0], " ")
	}
	for i:= 0;i <len(twotwo) ;i++{
		split := strings.Split(twotwo[i], " ")
		fmt.Print(split[0], " ")
	}
	for i:= 0;i <len(three) ;i++{
		split := strings.Split(three[i], " ")
		fmt.Print(split[0], " ")
	}
	for i:= 0;i <len(house) ;i++{
		split := strings.Split(house[i], " ")
		fmt.Print(split[0], " ")
	}
	for i:= 0;i <len(four) ;i++{
		split := strings.Split(four[i], " ")
		fmt.Print(split[0], " ")
	}
	for i:= 0;i <len(five) ;i++{
		split := strings.Split(five[i], " ")
		fmt.Print(split[0], " ")
	}


	for _, hand := range worse {
		handsCount++
		split := strings.Split(hand, " ")
		n, _ := strconv.Atoi(split[1])
		totalWinnings += handsCount * n
	}
	
	for _, hand := range two {
		handsCount++
		split := strings.Split(hand, " ")
		n, _ := strconv.Atoi(split[1])
		totalWinnings += handsCount * n 
	}
	for _, hand := range twotwo {
		handsCount++
		split := strings.Split(hand, " ")
		n, _ := strconv.Atoi(split[1])
		totalWinnings += handsCount * n 
	}
	for _, hand := range three{
		handsCount++
		split := strings.Split(hand, " ")
		n, _ := strconv.Atoi(split[1])
		totalWinnings += handsCount * n 
	}
	for _, hand := range house {
		handsCount++
		split := strings.Split(hand, " ")
		n, _ := strconv.Atoi(split[1])
		totalWinnings += handsCount * n 
	}
	for _, hand := range four {
		handsCount++
		split := strings.Split(hand, " ")
		n, _ := strconv.Atoi(split[1])
		totalWinnings += handsCount * n 
	}
	for _, hand := range five {
		handsCount++
		split := strings.Split(hand, " ")
		n, _ := strconv.Atoi(split[1])
		totalWinnings += handsCount * n 
	}

	//fmt.Println(five, four, house, three, twotwo, two, worse)

	fmt.Println(totalWinnings, handsCount)
}
