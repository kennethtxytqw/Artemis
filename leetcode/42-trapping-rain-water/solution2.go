package main

func solution2(height []int) int {
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	ans := 0
	for left < right {
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		// What this block does is basically
		// 1. move right pointer until right pointer is larger than leftMax
		// or 2. move left pointer until left pointer is larger than rightMax
		if height[left] > height[right] {
			if rightMax > height[right] {
				// Since height[right] is less than height
				ans += rightMax - height[right]
			}
			right--
		} else {
			if leftMax > height[left] {
				ans += leftMax - height[left]
			}
			left++
		}
	}
	return ans
}
