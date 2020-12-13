package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	var maxID int
	ids := make([]int, len(values))
	for i, ticketID := range values {
		ticketID = strings.Map(func(r rune) rune {
			if r == 'B' || r == 'R' {
				return '1'
			}
			return '0'
		}, ticketID)

		id64, err := strconv.ParseInt(ticketID, 2, 64)
		if err != nil {
			break
		}
		id := int(id64)
		ids[i] = int(id)
		if int(id) > maxID {
			maxID = int(id)
		}
	}

	sort.Ints(ids)

	var currentID, myID int
	for _, id := range ids {
		if currentID == 0 {
			currentID = id
			continue
		}
		if id-currentID != 1 {
			myID = id - 1
			break
		}
		currentID = id
	}

	fmt.Println("The highest ticketID on the list is =>", maxID)
	fmt.Println("My ticketID is =>", myID)
}
