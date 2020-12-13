package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}

var colors = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n\n")

	total := 0
	total2 := 0
	for _, v := range values {
		cleanStr := strings.ReplaceAll(v, "\n", " ")
		field := strings.Split(cleanStr, " ")

		var passport passport
		for _, v := range field {
			elem := strings.Split(v, ":")
			switch elem[0] {
			case "byr":
				byr, err := strconv.Atoi(elem[1])
				if err != nil {
					continue
				}
				passport.byr = byr
			case "eyr":
				eyr, err := strconv.Atoi(elem[1])
				if err != nil {
					continue
				}
				passport.eyr = eyr
			case "iyr":
				iyr, err := strconv.Atoi(elem[1])
				if err != nil {
					continue
				}
				passport.iyr = iyr
			case "hgt":
				passport.hgt = elem[1]
			case "hcl":
				passport.hcl = elem[1]
			case "ecl":
				passport.ecl = elem[1]
			case "pid":
				passport.pid = elem[1]
			}
		}
		if passport.hasAllField() {
			total++
		}
		if passport.isValid() {
			total2++
		}
	}

	fmt.Println("The number of passports with all fields except cid is =>", total)
	fmt.Println("The number of valid passports is =>", total2)
}

func (p passport) hasAllField() bool {
	if p.byr == 0 || p.iyr == 0 || p.eyr == 0 || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "" {
		return false
	}
	return true
}

func (p passport) isValid() bool {
	if p.byr > 2002 || p.byr < 1920 {
		return false
	}
	if p.iyr > 2020 || p.iyr < 2010 {
		return false
	}
	if p.eyr > 2030 || p.eyr < 2020 {
		return false
	}

	if len(p.hgt) < 4 {
		return false
	}

	unit := p.hgt[len(p.hgt)-2 : len(p.hgt)]
	hgtValue, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
	if err != nil {
		return false
	}
	switch unit {
	case "cm":
		if hgtValue < 150 || hgtValue > 193 {
			return false
		}
	case "in":
		if hgtValue < 59 || hgtValue > 76 {
			return false
		}
	default:
		return false
	}

	matched, err := regexp.Match(`^#[a-z\d]{6}$`, []byte(p.hcl))
	if !matched || err != nil {
		return false
	}

	if _, ok := colors[p.ecl]; !ok {
		return false
	}

	if len(p.pid) != 9 {
		return false
	}
	for _, c := range p.pid {
		if c < '0' && c > '9' {
			return false
		}
	}

	return true
}
