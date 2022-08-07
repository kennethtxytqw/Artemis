package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in1 []int
	in2 []int
	out float64
}{

	{[]int{1, 3}, []int{2}, 2},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{1, 4}, []int{2, 3}, 2.5},
	{[]int{}, []int{1}, 1},
	{[]int{}, []int{2, 3}, 2.5},
	{[]int{2}, []int{1}, 1.5},
	{[]int{1}, []int{2, 3, 4}, 2.5},
	{[]int{4}, []int{1, 2, 3, 5}, 3},
	{[]int{}, []int{1, 2, 3, 4, 5}, 3},
}

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution1(test.in1, test.in2), fmt.Sprintf("in1: %v\n in2: %v", test.in1, test.in2))
	}
}
