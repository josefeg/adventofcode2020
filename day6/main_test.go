package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
	"",
}

func TestCountAnyoneYes(t *testing.T) {
	assert.Equal(t, 11, countAnyoneYes(input))
}

func TestCountEveryoneYes(t *testing.T) {
	assert.Equal(t, 6, countEveryoneYes(input))
}
