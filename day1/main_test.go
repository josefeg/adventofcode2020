package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testNumbers = []int{1721, 979, 366, 299, 675, 1456}

func TestFindAndMultiplyTwoNumbers(t *testing.T) {
	assert.Equal(t, 514579, findAndMultiplyTwoNumbers(testNumbers))
}

func TestFindAndMultiplyThreeNumbers(t *testing.T) {
	assert.Equal(t, 241861950, findAndMultiplyThreeNumbers(testNumbers))
}
