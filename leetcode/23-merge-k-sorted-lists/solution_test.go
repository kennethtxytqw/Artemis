package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in  []*ListNode
	out *ListNode
}{
	{
		in: []*ListNode{
			constructListNode([]int{1, 4, 5}),
			constructListNode([]int{1, 3, 4}),
			constructListNode([]int{2, 6})},
		out: constructListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
	},
}

func constructListNode(arr []int) *ListNode {
	var last *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		last = &ListNode{
			arr[i],
			last,
		}
	}
	return last
}

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution1(test.in), fmt.Sprintf("in: %v\n", test.out))
	}
}
