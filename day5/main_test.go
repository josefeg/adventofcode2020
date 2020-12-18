package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeatID(t *testing.T) {
	row, column, seatID := getSeatID("BFFFBBFRRR")
	assert.Equal(t, 70, row)
	assert.Equal(t, 7, column)
	assert.Equal(t, 567, seatID)

	row, column, seatID = getSeatID("FFFBBBFRRR")
	assert.Equal(t, 14, row)
	assert.Equal(t, 7, column)
	assert.Equal(t, 119, seatID)

	row, column, seatID = getSeatID("BBFFBBFRLL")
	assert.Equal(t, 102, row)
	assert.Equal(t, 4, column)
	assert.Equal(t, 820, seatID)
}
