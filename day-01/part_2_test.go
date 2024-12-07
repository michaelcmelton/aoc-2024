package day01_test

import (
	day01 "aoc-2024/day-01"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProcessAnswer_Sample_Part2(t *testing.T) {

	answer := day01.ProcessAnswerPart2("./part_1_sample.txt")

	assert.Equal(t, 31, answer)
}

func Test_ProcessAnswer_Actual_Part2(t *testing.T) {

	answer := day01.ProcessAnswerPart2("./part_1_input.txt")

	assert.Equal(t, 26674158, answer)
}
