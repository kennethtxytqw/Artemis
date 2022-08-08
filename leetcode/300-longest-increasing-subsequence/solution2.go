package main

func bs(seq []int)

func solution2(nums []int) int {
	seq := make([]int, len(nums))

	seq[0] = nums[0]
	curr_index := 0
	for i := 1; i < len(nums); i++ {
		if seq[curr_index] < nums[i] {
			curr_index++
			seq[curr_index] = nums[i]
		} else {

		}
	}

	max_len := 0
	for _, v := range max_lens {
		if max_len < v {
			max_len = v
		}
	}
	return max_len
}
