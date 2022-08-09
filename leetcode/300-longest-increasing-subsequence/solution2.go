package main

func bs(seq []int, num int) int {
	low := 0
	high := len(seq) - 1
	for {
		if high == low {
			return high
		}
		mid := (high + low) / 2
		if seq[mid] == num {
			return mid
		} else if seq[mid] < num {
			low = mid + 1
		} else {
			high = mid
		}
	}
}

func solution2(nums []int) int {
	seq := make([]int, len(nums))

	seq[0] = nums[0]
	curr_index := 0
	for i := 1; i < len(nums); i++ {
		if seq[curr_index] == nums[i] {
			continue
		} else if seq[curr_index] < nums[i] {
			curr_index++
			seq[curr_index] = nums[i]
		} else {
			pos := bs(seq[0:curr_index+1], nums[i])
			seq[pos] = nums[i]
		}
	}
	return curr_index + 1
}
