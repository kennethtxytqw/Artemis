package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	s     string
	words []string
	out   []int
}{

	{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
}

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution1(test.s, test.words), fmt.Sprintf("in: %v\n", test.out))
	}
}
