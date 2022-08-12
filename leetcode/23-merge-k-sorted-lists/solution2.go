package main

func solution2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	step := 1
	for step < len(lists) {
		for i := 0; i < len(lists)-step; i += 2 * step {
			lists[i] = merge(lists[i], lists[i+step])
		}
		step *= 2
	}
	return lists[0]
}
