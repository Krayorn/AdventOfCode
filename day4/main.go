package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p Passport) isValidBytes() bool {
	return true
}

func (p Passport) isValid() bool {
	if p.byr == "" || len(p.byr) < 4 {
		return false
	}
	byr, err := strconv.Atoi(p.byr)
	if err != nil || byr > 2002 || byr < 1920 {
		return false
	}

	if p.iyr == "" || len(p.iyr) < 4 {
		return false
	}
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil || iyr > 2020 || iyr < 2010 {
		return false
	}

	if p.eyr == "" || len(p.eyr) < 4 {
		return false
	}
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil || eyr > 2030 || eyr < 2020 {
		return false
	}

	if !strings.HasSuffix(p.hgt, "cm") && !strings.HasSuffix(p.hgt, "in") {
		return false
	}

	if strings.HasSuffix(p.hgt, "cm") {
		hgt, err := strconv.Atoi(strings.Replace(p.hgt, "cm", "", 1))
		if err != nil || hgt < 150 || hgt > 193 {
			return false
		}
	}

	if strings.HasSuffix(p.hgt, "in") {
		hgt, err := strconv.Atoi(strings.Replace(p.hgt, "in", "", 1))
		if err != nil || hgt < 59 || hgt > 76 {
			return false
		}
	}

	matched, err := regexp.Match(`^#[a-z\d]{6}$`, []byte(p.hcl))
	if !matched || err != nil {
		return false
	}

	colors := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	if !stringInSlice(p.ecl, colors) {
		return false
	}

	newmatched, err := regexp.Match(`^\d{9}$`, []byte(p.pid))
	if !newmatched || err != nil {
		return false
	}

	return true
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func regexpValid(passport string) bool {
	rpid := regexp.MustCompile(`pid:\d{9}\b`)
	pidMatches := rpid.FindStringSubmatchIndex(passport)
	if len(pidMatches) == 0 || pidMatches[1]-pidMatches[0] != 13 {
		return false
	}

	hcl, err := regexp.Match(`hcl:#([[a-f\d]{6})`, []byte(passport))
	if !hcl || err != nil {
		return false
	}

	ecl, err := regexp.Match(`ecl:(brn|blu|amb|gry|grn|hzl|oth)`, []byte(passport))
	if !ecl || err != nil {
		return false
	}

	rbyr := regexp.MustCompile(`byr:(?P<byr>\d{4})`)
	byrMatches := rbyr.FindStringSubmatchIndex(passport)
	if len(byrMatches) == 0 {
		return false
	}

	byr, err := strconv.Atoi(passport[byrMatches[0]+4 : byrMatches[1]])
	if err != nil || byr > 2002 || byr < 1920 {
		return false
	}

	riyr := regexp.MustCompile(`iyr:(?P<iyr>\d{4})`)
	iyrMatches := riyr.FindStringSubmatchIndex(passport)
	if len(iyrMatches) == 0 {
		return false
	}

	iyr, err := strconv.Atoi(passport[iyrMatches[0]+4 : iyrMatches[1]])
	if err != nil || iyr > 2020 || iyr < 2010 {
		return false
	}

	reyr := regexp.MustCompile(`eyr:(?P<eyr>\d{4})`)
	eyrMatches := reyr.FindStringSubmatchIndex(passport)
	if len(eyrMatches) == 0 {
		return false
	}

	eyr, err := strconv.Atoi(passport[eyrMatches[0]+4 : eyrMatches[1]])
	if err != nil || eyr > 2030 || eyr < 2020 {
		return false
	}

	rhgt := regexp.MustCompile(`hgt:(\d{2,3}(in|cm))`)
	hgtMatches := rhgt.FindStringSubmatchIndex(passport)
	if len(hgtMatches) == 0 {
		return false
	}

	hgt, err := strconv.Atoi(passport[hgtMatches[2] : hgtMatches[3]-2])
	if err != nil {
		return false
	}

	if passport[hgtMatches[4]:hgtMatches[5]] == "cm" {
		if hgt < 150 || hgt > 193 {
			return false
		}
	}

	if passport[hgtMatches[4]:hgtMatches[5]] == "in" {
		if hgt < 59 || hgt > 76 {
			return false
		}
	}

	return true
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n\n")

	total := 0

	for _, v := range values {
		cleanStr := strings.Replace(v, "\n", " ", -1)
		passport := strings.Split(cleanStr, " ")

		var vPassport Passport
		for _, v := range passport {
			elem := strings.Split(v, ":")
			if elem[0] == "byr" {
				vPassport.byr = elem[1]
			}
			if elem[0] == "iyr" {
				vPassport.iyr = elem[1]
			}
			if elem[0] == "eyr" {
				vPassport.eyr = elem[1]
			}
			if elem[0] == "hgt" {
				vPassport.hgt = elem[1]
			}
			if elem[0] == "hcl" {
				vPassport.hcl = elem[1]
			}
			if elem[0] == "ecl" {
				vPassport.ecl = elem[1]
			}
			if elem[0] == "pid" {
				vPassport.pid = elem[1]
			}
			if elem[0] == "cid" {
				vPassport.cid = elem[1]
			}
		}
		if vPassport.isValidBytes() {
			total++
		}
	}
	fmt.Println(total)
}
