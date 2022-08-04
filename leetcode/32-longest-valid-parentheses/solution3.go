package main

func solution3(s string) int {
	left := 0
	right := 0
	maxlen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left += 1
		} else {
			right += 1
		}
		if left == right {
			if maxlen < right*2 {
				maxlen = right * 2
			}
		} else if right > left {
			left = 0
			right = 0
		}
	}
	left = 0
	right = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if maxlen < right*2 {
				maxlen = right * 2
			}
		} else if right < left {
			left = 0
			right = 0
		}
	}
	return maxlen
}
