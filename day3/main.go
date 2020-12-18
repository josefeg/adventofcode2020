package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var slopes = [][]int{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func countTrees(areaMap [][]string, right, down int) int {
	x := 0
	trees := 0
	for y := 0; y < len(areaMap); y += down {
		if areaMap[y][x] == "#" {
			trees += 1
		}
		x += right
		if x >= len(areaMap[y]) {
			x -= len(areaMap[y])
		}
	}

	return trees
}

func multiplyTrees(areaMap [][]string) int {
	total := 1
	for _, slope := range slopes {
		trees := countTrees(areaMap, slope[0], slope[1])
		total *= trees
	}
	return total
}

func main() {
	contents, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	var areaMap [][]string
	for _, line := range lines {
		areaMap = append(areaMap, strings.Split(line, ""))
	}
	fmt.Println(countTrees(areaMap, 3, 1))
	fmt.Println(multiplyTrees(areaMap))
}
