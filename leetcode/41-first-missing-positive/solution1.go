package main

func solution1(nums []int) int {
	total_len := len(nums)
	arr := make([]bool, total_len)
	for _, num := range nums {
		if num <= total_len && num > 0 {
			arr[num-1] = true
		}
	}
	for i, isPresent := range arr {
		if !isPresent {
			return i + 1
		}
	}
	return len(arr) + 1
}
