package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

func TestFindFirstNotSumNumber(t *testing.T) {
	assert.Equal(t, 27, findFirstNotSumNumber(input, 5))
}

func TestFindEncryptionWeakness(t *testing.T) {
	assert.Equal(t, 62, findEncryptionWeakness(127, input))
}
