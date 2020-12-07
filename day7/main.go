package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Content struct {
	Color  string
	Number int
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	total := 0

	bags := make(map[string][]Content)

	for _, v := range values {
		split := strings.Split(v, " bags contain ")
		bagColor := split[0]

		if split[1] == "no other bags." {
			bags[bagColor] = []Content{}
		} else {
			bagsSplit := strings.Split(split[1], ", ")
			contains := make([]Content, len(bagsSplit))
			for i, bc := range bagsSplit {
				splittedNamed := strings.Split(bc, " ")
				cleanedName := strings.Join(splittedNamed[1:len(splittedNamed)-1], " ")
				number, _ := strconv.Atoi(splittedNamed[0])
				contains[i] = Content{Color: cleanedName, Number: number}
			}
			bags[bagColor] = contains
		}
	}

	fmt.Println("BAGS IN SHINY GOLD", getNeededBagsIn("shiny gold", bags))

	exploredBags := make(map[string]bool)
	for mainColor := range bags {
		if exploredBags[mainColor] {
			continue
		}
		if canBeInBag(mainColor, bags) {
			total++
		}
		exploredBags[mainColor] = true
	}
	fmt.Println("Bags that can contains shiny gold", total)
}

func getNeededBagsIn(color string, bags map[string][]Content) int {
	add := 0

	for _, content := range bags[color] {
		nb := getNeededBagsIn(content.Color, bags)
		add += content.Number * nb
		add += content.Number
	}
	return add
}

func canBeInBag(color string, bags map[string][]Content) bool {
	if stringInSlice("shiny gold", bags[color]) {
		return true
	}
	for _, content := range bags[color] {
		if canBeInBag(content.Color, bags) {
			return true
		}
	}
	return false
}

func stringInSlice(a string, list []Content) bool {
	for _, b := range list {
		if b.Color == a {
			return true
		}
	}
	return false
}
