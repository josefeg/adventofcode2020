package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc",
}

func TestFindValidAmountWithCount(t *testing.T) {
	assert.Equal(t, 2, validAmountWithCount(input))
}

func TestFindValidAmountWithPosition(t *testing.T) {
	assert.Equal(t, 1, validAmountWithPosition(input))
}
