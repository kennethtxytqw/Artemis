package main

func solution2(s string) int {
	if len(s) == 0 {
		return 0
	}
	memo := make([]int, len(s))
	maxlen := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				memo[i] = 2
				if i > 2 {
					memo[i] += memo[i-2]
				}
			} else {
				lookback := i - memo[i-1] - 1
				if lookback >= 0 && s[lookback] == '(' {
					memo[i] = memo[i-1] + 2
					if lookback-1 >= 0 {
						memo[i] += memo[lookback-1]
					}
				}
			}
		} else {
			memo[i] = 0
		}
		if memo[i] > maxlen {
			maxlen = memo[i]
		}
	}
	return maxlen
}
