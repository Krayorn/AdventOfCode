package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type comb struct {
	addr   []rune
	binary string
}

var memory = make(map[string]string)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	var mask string
	var pending []comb
	var floatings = make(map[int]bool)

	for _, v := range values {
		if strings.HasPrefix(v, "mask") {
			if len(pending) != 0 {
				keys := make([]int, 0, len(floatings))
				for k := range floatings {
					keys = append(keys, k)
				}
				assignAllMemoryAddresses(keys, pending, 0)
				pending = []comb{}
				floatings = map[int]bool{}
			}
			mask = v[7:]
		} else {
			split := strings.Split(v, " = ")
			memoryAddr := split[0][4 : len(split[0])-1]
			number, _ := strconv.Atoi(split[1])
			binary := strconv.FormatInt(int64(number), 2)
			maskedBinary := make([]rune, 36)
			offset := 35 - len(binary)
			for i := range mask {
				if i <= offset {
					maskedBinary[i] = '0'
				} else {
					maskedBinary[i] = rune(binary[i-offset-1])
				}
			}
			n, _ := strconv.Atoi(memoryAddr)
			binaryAddr := strconv.FormatInt(int64(n), 2)
			maskedAddr := make([]rune, 36)
			offset = 35 - len(binaryAddr)
			for i, c := range mask {
				if c == 'X' {
					floatings[i] = true
					maskedAddr[i] = c
				} else if c == '1' {
					maskedAddr[i] = c
				} else if i <= offset {
					maskedAddr[i] = '0'
				} else {
					maskedAddr[i] = rune(binaryAddr[i-offset-1])
				}
			}
			pending = append(pending, comb{addr: maskedAddr, binary: string(maskedBinary)})
		}
	}
	keys := make([]int, 0, len(floatings))
	for k := range floatings {
		keys = append(keys, k)
	}
	assignAllMemoryAddresses(keys, pending, 0)
	pending = []comb{}
	floatings = map[int]bool{}
	total := 0
	for _, v := range memory {
		n, _ := strconv.ParseInt(v, 2, 64)
		total += int(n)
	}
	fmt.Println(total)
}

func assignAllMemoryAddresses(floatings []int, combs []comb, depth int) {
	if depth >= len(floatings) {
		for _, comb := range combs {
			n, _ := strconv.ParseInt(string(comb.addr), 2, 64)
			memory[strconv.Itoa(int(n))] = comb.binary
		}
		return
	}
	for _, comb := range combs {
		comb.addr[floatings[depth]] = '0'
	}
	assignAllMemoryAddresses(floatings, combs, depth+1)
	for _, comb := range combs {
		comb.addr[floatings[depth]] = '1'
	}
	assignAllMemoryAddresses(floatings, combs, depth+1)
}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	textFile, err := os.Open("./input.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	byteValue, _ := ioutil.ReadAll(textFile)
// 	values := strings.Split(string(byteValue), "\n")

// 	memory := make(map[string]string)
// 	var mask string
// 	for _, v := range values {
// 		if strings.HasPrefix(v, "mask") {
// 			mask = v[7:]
// 		} else {
// 			split := strings.Split(v, " = ")
// 			memoryAddr := split[0][4 : len(split[0])-1]
// 			number, _ := strconv.Atoi(split[1])
// 			binary := strconv.FormatInt(int64(number), 2)
// 			maskedBinary := make([]rune, 36)
// 			offset := 35 - len(binary)
// 			for i, c := range mask {
// 				if c != 'X' {
// 					maskedBinary[i] = c
// 				} else if i <= offset {
// 					maskedBinary[i] = '0'
// 				} else {
// 					maskedBinary[i] = rune(binary[i-offset-1])
// 				}
// 			}
// 			memory[memoryAddr] = string(maskedBinary)
// 		}
// 	}
// 	total := 0
// 	for _, v := range memory {
// 		n, _ := strconv.ParseInt(v, 2, 64)
// 		total += int(n)
// 	}
// 	fmt.Println(total)
// }
