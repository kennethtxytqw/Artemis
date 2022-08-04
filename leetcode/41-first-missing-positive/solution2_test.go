package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution2(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.out, solution2(test.in))
	}
}
