package main

func solution2(destination []int, k int) string {
	v := destination[0]
	h := destination[1]
	memo := make([][]int, h+1)
	// Count number of paths to destination from each coordinate
	// This is actually doing nCr but without the factorial, thus bypassing the size limitation
	for i := 0; i <= h; i++ {
		memo[i] = make([]int, v+1)
		memo[i][v] = 1
	}
	for i := 0; i <= v; i++ {
		memo[h][i] = 1
	}

	for i := h - 1; i >= 0; i-- {
		for j := v - 1; j >= 0; j-- {
			memo[i][j] = memo[i+1][j] + memo[i][j+1]
		}
	}

	ans := ""
	src := []int{0, 0}
	for {
		if memo[src[0]+1][src[1]] >= k {
			ans += "H"
			src[0]++
		} else {
			ans += "V"
			k -= memo[src[0]+1][src[1]]
			src[1]++
		}
		if src[0] == h || src[1] == v {
			break
		}
	}

	for i := src[0]; i < h; i++ {
		ans += "H"
	}
	for i := src[1]; i < v; i++ {
		ans += "V"
	}
	return ans
}
