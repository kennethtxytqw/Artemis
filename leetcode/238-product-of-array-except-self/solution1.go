package main

import "fmt"

func solution1(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1

	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	store := nums[len(nums)-1]
	nums[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		store, nums[i] = nums[i], store*nums[i+1]
	}
	fmt.Println(nums, ans)
	for i := 0; i < len(ans); i++ {
		ans[i] = ans[i] * nums[i]
	}
	return ans
}
