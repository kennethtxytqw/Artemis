package main

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for {
		if left >= right {
			break
		}
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}
