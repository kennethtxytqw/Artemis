package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	in  string
	out int
}{
	{"(()", 2},
	{")()())", 4},
	{"", 0},
	{"()(()", 2},
	{"(()(((()", 2},
	{"()()", 4},
	{"()(())", 6},
}

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution1(test.in), fmt.Sprintf("in: %s", test.in))
	}
}

func TestSolution2(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution2(test.in), fmt.Sprintf("in: %s", test.in))
	}
}

func TestSolution3(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution3(test.in), fmt.Sprintf("in: %s", test.in))
	}
}
