package main

func solution2(nums []int) int {
	total_len := len(nums)
	for i := 0; i < total_len; i++ {
		num := nums[i]
		for {
			if num <= 0 || num > total_len || nums[num-1] == num {
				break
			}
			temp := nums[num-1]
			nums[num-1] = num
			num = temp
		}
	}
	for i, val := range nums {
		if val != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}