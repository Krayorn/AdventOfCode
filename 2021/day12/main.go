package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Location struct {
	name         string
	destinations []*Location
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	locations := make(map[string]Location, 0)
	for _, str := range values {
		split := strings.Split(str, "-")
		if _, ok := locations[split[0]]; !ok {
			dests := make([]*Location, 0)
			locations[split[0]] = Location{name: split[0], destinations: dests}
		}
		if _, ok := locations[split[1]]; !ok {
			dests := make([]*Location, 0)
			locations[split[1]] = Location{name: split[1], destinations: dests}
		}

		from := locations[split[0]]
		to := locations[split[1]]
		from.destinations = append(from.destinations, &to)
		to.destinations = append(to.destinations, &from)
		locations[split[0]] = from
		locations[split[1]] = to
	}

	count := 0
	count2 := 0
	trail := make([]string, 0)
	explore(locations["start"], locations, &count2, trail, true)
	explore(locations["start"], locations, &count, trail, false)
	fmt.Println(fmt.Sprintf("The number of paths available when there is no time to visit a single small cave more than once is %d.", count))
	fmt.Println(fmt.Sprintf("The number of paths available when there is time to visit a single small cave twice is %d.", count2))
}

func explore(location Location, locations map[string]Location, count *int, trail []string, withTime bool) {
	if location.name == "end" {
		*count++
		return
	}
	trail = append(trail, location.name)

	for _, dest := range location.destinations {
		if dest.isAvailable(trail, withTime) {
			explore(locations[dest.name], locations, count, trail, withTime)
		}
	}

}

func (loca Location) isAvailable(trail []string, withTime bool) bool {
	if strings.ToLower(loca.name) == loca.name {
		if !withTime {
			for _, name := range trail {
				if name == loca.name {
					return false
				}
			}
			return true
		}
		if loca.name == "start" || loca.name == "end" {
			for _, name := range trail {
				if name == loca.name {
					return false
				}
			}
		} else {
			countLoca := make(map[string]int, 0)
			for _, name := range trail {
				if _, ok := countLoca[name]; !ok {
					countLoca[name] = 0
				}
				countLoca[name]++
			}
			multipleVisitsInSmall := false
			for name, count := range countLoca {
				if count > 1 && strings.ToLower(name) == name {
					multipleVisitsInSmall = true
				}
			}
			if _, ok := countLoca[loca.name]; ok && multipleVisitsInSmall {
				return false
			}
		}
	}
	return true
}

func printMap(locations map[string]Location) {
	for _, location := range locations {
		fmt.Println("name", location.name, " - ")
		fmt.Print("destinations - ")
		for _, dest := range location.destinations {
			fmt.Print(dest.name, " - ")
		}
		fmt.Print("\n")
	}
}
