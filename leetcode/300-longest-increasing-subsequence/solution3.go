package main

func solution3(nums []int) (int, []int) {
	seq := make([]int, len(nums))
	seqIndex := make([]int, len(nums))
	predecessor := make([]int, len(nums))

	seq[0] = nums[0]
	seqIndex[0] = 0
	predecessor[0] = -1
	curr_index := 0
	for i := 1; i < len(nums); i++ {
		if seq[curr_index] == nums[i] {
			continue
		} else if seq[curr_index] < nums[i] {
			curr_index++
			seq[curr_index] = nums[i]
			seqIndex[curr_index] = i
			predecessor[i] = seqIndex[curr_index-1]
		} else {
			pos := bs(seq[0:curr_index+1], nums[i])
			seq[pos] = nums[i]
			seqIndex[pos] = i
			if pos > 1 {
				predecessor[i] = seqIndex[pos-1]
			} else {
				predecessor[i] = -1
			}
		}
	}

	ans := make([]int, curr_index+1)
	traced := seqIndex[curr_index]
	for i := curr_index; i > 0; i-- {
		ans[i] = nums[traced]
		traced = predecessor[traced]
	}
	return len(ans), ans
}
