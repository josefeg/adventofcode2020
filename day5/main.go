package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func calc(directions string, up, down rune, max int) int {
	pos := 0
	lowerBound := 0
	upperBound := max - 1
	for i, r := range directions {
		bound := (max / int(math.Pow(2, float64(i)))) / 2
		if r == up {
			upperBound -= bound
			pos = lowerBound
		} else if r == down {
			lowerBound += bound
			pos = upperBound
		}
	}
	return pos
}

func getSeatID(input string) (int, int, int) {
	row := calc(input[:7], 'F', 'B', 128)
	col := calc(input[7:], 'L', 'R', 8)

	return row, col, (row * 8) + col
}

func main() {
	contents, err := ioutil.ReadFile("day5/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	var ids []int
	for _, line := range lines {
		_, _, id := getSeatID(line)
		ids = append(ids, id)
	}
	sort.Ints(ids)
	fmt.Printf("find max => %d\n", ids[len(ids)-1])

	for i := 1; i < len(ids); i++ {
		if ids[i] == ids[i-1]+2 {
			fmt.Println("seatID =>", ids[i]-1)
			return
		}
	}

}
