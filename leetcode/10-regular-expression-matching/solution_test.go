package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	s        string
	p        string
	expected bool
}{
	{"mississippi", "mis*is*p*.", false},
	{"aaa", "a.a", true},
	{"ab", ".*c", false},
	{"aasdfasdfasdfasdfas",
		"aasdf.*asdf.*asdf.*asdf.*s", true},
}

func TestSolution1(t *testing.T) {
	for _, test := range tests {
		assert := assert.New(t)
		assert.Equal(test.expected, solution1(test.s, test.p), fmt.Sprintf("s: %s, p: %s", test.s, test.p))
	}
}
