package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

func TestExecute(t *testing.T) {
	assert.Equal(t, 5, execute(input))
}

func TestExecuteWithChanges(t *testing.T) {
	acc, _ := executeWithChanges(input, 0, 0, true)
	assert.Equal(t, 8, acc)
}
