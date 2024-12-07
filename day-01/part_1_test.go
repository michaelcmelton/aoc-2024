package day01_test

import (
	day01 "aoc-2024/day-01"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProcessAnswer_Sample(t *testing.T) {
	answer := day01.ProcessAnswer("./part_1_sample.txt")
	assert.Equal(t, 11, answer)
}

func Test_ProcessAnswer_Actual(t *testing.T) {
	answer := day01.ProcessAnswer("./part_1_input.txt")
	assert.Equal(t, 1830467, answer)
}
