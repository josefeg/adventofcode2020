package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func validAmountWithCount(passwordsAndPolicies []string) int {
	amount := 0
	for _, passwordAndPolicy := range passwordsAndPolicies {
		pp := strings.Split(passwordAndPolicy, ":")
		if isValidCount(pp[0], strings.TrimSpace(pp[1])) {
			amount += 1
		}
	}
	return amount
}

func isValidCount(policy, password string) bool {
	policyParts := strings.Split(policy, " ")
	constraints := strings.Split(policyParts[0], "-")

	min, err := strconv.Atoi(constraints[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(constraints[1])
	if err != nil {
		panic(err)
	}

	count := 0
	target := []rune(policyParts[1])[0]
	for _, r := range password {
		if r == target {
			count += 1
		}
	}

	return (count >= min) && (count <= max)
}

func validAmountWithPosition(passwordsAndPolicies []string) int {
	amount := 0
	for _, passwordAndPolicy := range passwordsAndPolicies {
		pp := strings.Split(passwordAndPolicy, ":")
		if isValidPosition(pp[0], strings.TrimSpace(pp[1])) {
			amount += 1
		}
	}
	return amount
}

func isValidPosition(policy, password string) bool {
	policyParts := strings.Split(policy, " ")
	constraints := strings.Split(policyParts[0], "-")

	startIndex, err := strconv.Atoi(constraints[0])
	if err != nil {
		panic(err)
	}
	startIndex -= 1

	endIndex, err := strconv.Atoi(constraints[1])
	if err != nil {
		panic(err)
	}
	endIndex -= 1

	target := []rune(policyParts[1])[0]
	passwordRunes := []rune(password)

	return (passwordRunes[startIndex] == target && passwordRunes[endIndex] != target) ||
		(passwordRunes[startIndex] != target && passwordRunes[endIndex] == target)
}

func main() {
	contents, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	fmt.Println(validAmountWithCount(lines))
	fmt.Println(validAmountWithPosition(lines))
}
