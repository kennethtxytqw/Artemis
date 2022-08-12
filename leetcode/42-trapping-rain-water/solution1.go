package main

func solution1(height []int) int {
	water := 0

	for {
		isTop := true
		left := -1

		for i := 0; i < len(height); i++ {
			if height[i] == 0 {
				continue
			}
			if left == -1 {
				left = i
			} else {
				water += i - left - 1
				left = i
				isTop = false
			}
			height[i]--
		}
		if isTop {
			break
		}
	}

	return water
}
