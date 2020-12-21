package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	op      string
	arg     int
	visited bool
}

func execute(input []string) int {
	program := parse(input)

	ip := 0
	acc := 0
	for {
		i := program[ip]
		if i.visited {
			return acc
		}

		i.visited = true
		if i.op == "nop" {
			ip++
			continue
		}
		if i.op == "acc" {
			ip++
			acc += i.arg
			continue
		}
		if i.op == "jmp" {
			ip += i.arg
			continue
		}
	}
}

func reset(program []*instruction) {
	for _, i := range program {
		i.visited = false
	}
}

func executeWithChanges(input []string, ip, acc int, shouldChange bool) (int, bool) {
	program := parse(input)

	for ip < len(program) {
		i := program[ip]
		if i.visited {
			return acc, false
		}

		i.visited = true
		if i.op == "nop" {
			if shouldChange {
				result, finishedSuccessfully := executeWithChanges(input, ip+i.arg, acc, false)
				if finishedSuccessfully {
					return result, true
				}
			}

			ip++
			continue
		}
		if i.op == "acc" {
			ip++
			acc += i.arg
			continue
		}
		if i.op == "jmp" {
			if shouldChange {
				result, finishedSuccessfully := executeWithChanges(input, ip+1, acc, false)
				if finishedSuccessfully {
					return result, true
				}
			}
			ip += i.arg
			continue
		}

	}
	return acc, true
}

func parse(input []string) []*instruction {
	var instructions []*instruction
	for _, line := range input {
		parts := strings.Split(line, " ")
		arg, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		i := instruction{
			op:  parts[0],
			arg: arg,
		}
		instructions = append(instructions, &i)
	}
	return instructions
}

func main() {
	contents, err := ioutil.ReadFile("day8/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	fmt.Println(execute(lines))

	result, _ := executeWithChanges(lines, 0, 0, true)
	fmt.Println(result)
}
