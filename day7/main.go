package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const target = "shiny gold"

var contentsRegex = regexp.MustCompile(`((\d+)\s+\w+\s+\w+)`)

func keys(m map[string]bool) []string {
	var ks []string
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}

func canContainBag(colors map[string]bool, rules []string, targets ...string) int {
	if len(targets) == 0 {
		return len(colors)
	}

	newColors := make(map[string]bool)
	for _, target := range targets {
		for _, rule := range rules {
			matches, err := regexp.MatchString(fmt.Sprintf(`\d+\s+%s`, target), rule)
			if err != nil {
				panic(err)
			}
			if matches {
				color := strings.Join(strings.Split(rule, " ")[:2], " ")
				newColors[color] = true
			}
		}
	}
	for k, _ := range newColors {
		colors[k] = true
	}

	return canContainBag(colors, rules, keys(newColors)...)
}

func countAllBags(rules []string, target string) int {
	count := 0
	for _, rule := range rules {
		matches, err := regexp.MatchString(fmt.Sprintf(`^%s`, target), rule)
		if err != nil {
			panic(err)
		}
		if !matches {
			continue
		}

		contents := contentsRegex.FindAllString(rule, -1)
		for _, content := range contents {
			x := strings.Split(content, " ")
			amount, _ := strconv.Atoi(x[0])
			color := strings.Join(x[1:], " ")

			count = count + amount + (amount * countAllBags(rules, color))
		}
	}

	return count
}

func main() {
	contents, err := ioutil.ReadFile("day7/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	state := make(map[string]bool)
	fmt.Println(canContainBag(state, lines, target))
	fmt.Println(countAllBags(lines, target))
}
