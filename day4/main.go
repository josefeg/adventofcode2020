package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var eyeColors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func separatePassports(input []string) []string {
	var passports []string
	passport := ""
	for _, line := range input {
		if line == "" {
			passports = append(passports, passport)
			passport = ""
		}
		if passport == "" {
			passport = line
		} else {
			passport = fmt.Sprintf("%s %s", passport, line)
		}
	}
	return passports
}

func parsePassports(input []string) []map[string]string {
	var parsedPassports []map[string]string

	passports := separatePassports(input)
	for _, passport := range passports {
		kvps := strings.Split(passport, " ")

		parsedPassport := make(map[string]string)
		for _, kvp := range kvps {
			kv := strings.Split(kvp, ":")
			parsedPassport[kv[0]] = kv[1]
		}
		parsedPassports = append(parsedPassports, parsedPassport)
	}
	return parsedPassports
}

func isValidPassport(passport map[string]string) bool {
	if len(passport) <= 6 {
		return false
	}

	delete(passport, "cid")
	if len(passport) == 7 {
		return validateFields(passport)
	}
	return false
}

func validateFields(passport map[string]string) bool {
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := strings.TrimSpace(passport["hgt"])
	r := regexp.MustCompile(`^(\d+)(cm|in)`)
	matches := r.FindStringSubmatch(hgt)
	if len(matches) != 3 {
		return false
	}
	length, _ := strconv.Atoi(matches[1])
	unit := matches[2]
	if unit == "cm" && (length < 150 || length > 193) {
		return false
	}
	if unit == "in" && (length < 59 || length > 76) {
		return false
	}

	hcl := passport["hcl"]
	if matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, hcl); !matched {
		return false
	}

	ecl := passport["ecl"]
	if !eyeColors[ecl] {
		return false
	}

	pid := passport["pid"]
	if matched, _ := regexp.MatchString(`^\d{9}$`, pid); !matched {
		return false
	}

	return true
}

func countValidPassports(input []string) int {
	amount := 0
	passports := parsePassports(input)
	for _, p := range passports {
		if isValidPassport(p) {
			amount += 1
		}
	}
	return amount
}

func main() {
	contents, err := ioutil.ReadFile("day4/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	fmt.Println(countValidPassports(lines))
}
