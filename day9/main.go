package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func findFirstNotSumNumber(input []int, premableSize int) int {
	for i := premableSize; i < len(input); i++ {
		if !canResultFromSumOfTwoNumbers(input[i], input[i-premableSize:i]) {
			return input[i]
		}
	}
	return -1
}

func canResultFromSumOfTwoNumbers(target int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == target {
				return true
			}
		}
	}
	return false
}

func findSeries(target int, input []int) (int, int) {
	for i := 0; i < len(input); i++ {
		count := input[i]
		for j := i + 1; j < len(input); j++ {
			count += input[j]
			if count == target {
				return i, j
			}
			if count > target {
				break
			}
		}
	}
	panic(errors.New("could not find series"))
}

func findEncryptionWeakness(target int, input []int) int {
	lower, upper := findSeries(target, input)
	series := input[lower : upper+1]
	sort.Ints(series)
	min := series[0]
	max := series[len(series)-1]
	return min + max
}

func main() {
	contents, err := ioutil.ReadFile("day9/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	var input []int
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		input = append(input, i)
	}

	firstNotSumNumber := findFirstNotSumNumber(input, 25)
	fmt.Println(firstNotSumNumber)
	fmt.Println(findEncryptionWeakness(firstNotSumNumber, input))
}
