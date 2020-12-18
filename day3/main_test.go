package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var areaMap = [][]string{
	strings.Split("..##.......", ""),
	strings.Split("#...#...#..", ""),
	strings.Split(".#....#..#.", ""),
	strings.Split("..#.#...#.#", ""),
	strings.Split(".#...##..#.", ""),
	strings.Split("..#.##.....", ""),
	strings.Split(".#.#.#....#", ""),
	strings.Split(".#........#", ""),
	strings.Split("#.##...#...", ""),
	strings.Split("#...##....#", ""),
	strings.Split(".#..#...#.#", ""),
}



func TestCountTrees(t *testing.T) {
	assert.Equal(t, 7, countTrees(areaMap, 3, 1))
}

func TestMultiplyTrees(t *testing.T) {
	assert.Equal(t, 336, multiplyTrees(areaMap))
}

