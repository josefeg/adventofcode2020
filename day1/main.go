package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func findAndMultiplyTwoNumbers(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				return numbers[i] * numbers[j]
			}
		}
	}
	return 0
}

func findAndMultiplyThreeNumbers(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					return numbers[i] * numbers[j] * numbers[k]
				}
			}
		}
	}
	return 0
}

func main() {
	contents, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	var numbers []int
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	fmt.Println(findAndMultiplyTwoNumbers(numbers))
	fmt.Println(findAndMultiplyThreeNumbers(numbers))
}
