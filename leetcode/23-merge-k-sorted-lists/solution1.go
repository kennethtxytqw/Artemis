package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(list_a *ListNode, list_b *ListNode) *ListNode {
	if list_a == nil {
		return list_b
	}
	if list_b == nil {
		return list_a
	}
	root := &ListNode{}
	tail := root
	currA := list_a
	currB := list_b

	for {
		if currB == nil {
			tail.Next = currA
			return root.Next
		}
		if currA == nil {
			tail.Next = currB
			return root.Next
		}
		if currA.Val >= currB.Val {
			tail.Next = currB
			currB = currB.Next
		} else {
			tail.Next = currA
			currA = currA.Next
		}
		tail = tail.Next
	}
}

func solution1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	root := lists[0]

	for i := 1; i < len(lists); i++ {
		root = merge(root, lists[i])
	}
	return root
}
