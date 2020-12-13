package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
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

	time, _ := strconv.Atoi(values[0])

	var ids []int
	max := 0
	maxIndex := 0
	for i, v := range strings.Split(values[1], ",") {
		if v != "x" {
			id, _ := strconv.Atoi(v)
			if max == 0 || id > max {
				max = id
				maxIndex = i
			}
			ids = append(ids, id)
		} else {
			ids = append(ids, -1)
		}
	}

	fmt.Println("Product of earliest bus and number of minutes ot wait =>", findID(time, ids))
	fmt.Println("Earliest timestamp for all listed bus to depart with correct offset =>", findSuit(ids, max, maxIndex))
	fmt.Println("Earliest timestamp for all listed bus to depart with correct offset =>", findSuitStupid(ids))
}

func findSuitStupid(ids []int) int {
	cleaned := make(map[int]bool)

	loop := 1
	iter := 0
	for {
		goodRun := true
		for i, id := range ids {
			if id == -1 {
				continue
			}
			if (iter+i)%id != 0 {
				goodRun = false
				continue
			} else {
				if _, ok := cleaned[i]; ok {
					continue
				}
				cleaned[i] = true
				loop *= id
				if loop < 0 {
					fmt.Println("error")
					return -1
				}
			}
		}
		if goodRun {
			return iter
		}
		iter += loop
	}
}

var one = big.NewInt(1)

// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go implem
func chineseRemainerTheorem(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func findSuit(ids []int, max int, maxIndex int) *big.Int {
	var newIds []*big.Int
	var modulos []*big.Int
	for i, id := range ids {
		if id == -1 {
			continue
		}
		newIds = append(newIds, big.NewInt(int64(id)))
		modulos = append(modulos, big.NewInt(int64(id-i%id)))
	}
	res, _ := chineseRemainerTheorem(modulos, newIds)
	return res
}

func findID(time int, ids []int) int {
	start := time
	for {
		for _, id := range ids {
			if id == -1 {
				continue
			}
			if start%id == 0 {
				return (start - time) * id
			}
		}
		start++
	}
}
