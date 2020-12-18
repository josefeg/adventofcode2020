package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")

	_ = lines
}