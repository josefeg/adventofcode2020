package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func separateGroups(input []string) []string {
	var groups []string
	group := ""
	for _, line := range input {
		if line == "" {
			groups = append(groups, group)
			group = ""
		}
		if group == "" {
			group = line
			continue
		}
		group = fmt.Sprintf("%s %s", group, line)
	}
	return groups
}

func removeDuplicates(groups []string) []map[string]bool {
	var result []map[string]bool
	for _, group := range groups {
		s := strings.Split(group, "")
		g := make(map[string]bool)
		for _, s := range s {
			if s == " " {
				continue
			}
			g[s] = true
		}
		result = append(result, g)
	}
	return result
}

func countAnyoneYes(input []string) int {
	groups := separateGroups(input)
	questions := removeDuplicates(groups)

	count := 0
	for _, q := range questions {
		count += len(q)
	}

	return count
}

func keepDuplicates(groups []string) []map[rune]bool {
	var result []map[rune]bool
	for _, group := range groups {
		people := strings.Split(group, " ")
		duplicates := getDuplicatesInGroup(people)
		result = append(result, duplicates)

	}
	return result
}

func toMap(str string) map[rune]bool {
	m := make(map[rune]bool)
	for _, s := range str {
		m[s] = true
	}
	return m
}

func getDuplicatesInGroup(people []string) map[rune]bool {
	duplicates := toMap(people[0])

	for i := 1; i < len(people); i++ {
		m := make(map[rune]bool)
		for k, _ := range duplicates {
			p := toMap(people[i])
			v := p[k]
			if !v {
				continue
			}
			m[k] = true
		}
		duplicates = m
	}
	return duplicates
}

func countEveryoneYes(input []string) int {
	groups := separateGroups(input)
	duplicates := keepDuplicates(groups)

	count := 0
	for _, d := range duplicates {
		count += len(d)
	}

	return count
}

func main() {
	contents, err := ioutil.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	fmt.Println(countAnyoneYes(lines))
	fmt.Println(countEveryoneYes(lines))
}
