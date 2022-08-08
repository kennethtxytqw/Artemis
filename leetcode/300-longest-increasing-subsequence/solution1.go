package main

import "fmt"

const shouldPrint = false

func print(str string) {
	if shouldPrint {
		fmt.Println(str)
	}
}

func solution1(nums []int) int {
	max_lens := make([]int, len(nums))
	smallest_number := make([]int, len(nums))

	max_lens[len(nums)-1] = 1
	smallest_number[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		max_len := 0
		for j := i + 1; j < len(nums)-max_len; j++ {
			if nums[i] < nums[j] && max_lens[j] > max_len {
				max_len = max_lens[j]
			}
		}
		max_lens[i] = 1 + max_len
	}

	max_len := 0
	for _, v := range max_lens {
		if max_len < v {
			max_len = v
		}
	}
	return max_len
}
