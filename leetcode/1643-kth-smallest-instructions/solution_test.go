package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	coordinate []int
	k          int
	out        string
}{
	{[]int{2, 3}, 3, "HHVVH"},
	{[]int{2, 3}, 2, "HHVHV"},
	{[]int{2, 2}, 4, "VHHV"},
	{[]int{15, 15}, 155117520, "VVVVVVVVVVVVVVVHHHHHHHHHHHHHHH"},
}

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	for i, test := range tests {
		if i >= 3 {
			return
		}
		assert.Equal(test.out, solution1(test.coordinate, test.k), fmt.Sprintf("in: %v %d", test.coordinate, test.k))
	}
}
func TestSolution2(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution2(test.coordinate, test.k), fmt.Sprintf("in: %v %d", test.coordinate, test.k))
	}
}
