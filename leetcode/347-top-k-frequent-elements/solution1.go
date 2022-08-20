package main

const null = -(100000000)

func topKFrequent(nums []int, k int) []int {
	M := make(map[int]int)

	for _, n := range nums {
		M[n]++
	}

	ans := make([]int, k)
	for i := 0; i < len(ans); i++ {
		ans[i] = null
	}
	for element, count := range M {
		for i := 0; i < len(ans); i++ {
			if ans[i] == null || count > M[ans[i]] {
				for j := len(ans) - 1; j > i; j-- {
					ans[j] = ans[j-1]
				}
				ans[i] = element
				break
			}
		}
	}
	return ans
}
