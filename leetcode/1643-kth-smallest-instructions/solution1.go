package main

func solution1(destination []int, k int) string {
	h := destination[1]
	v := destination[0]
	memoFac := make(map[int]int64, h+v)
	var factorial func(int) int64
	factorial = func(n int) int64 {
		if n == 0 {
			return 1
		}
		if v, ok := memoFac[n]; ok {
			return v
		}
		memoFac[n] = int64(n) * factorial(n-1)
		return memoFac[n]
	}

	OcK := func(n int, k int) int {
		return int(factorial(n) / factorial(k) / factorial(n-k))
	}

	count := 0
	ans := ""
	for {
		for i := 0; i <= h; i++ {
			next := OcK(v+i-1, i)
			// fmt.Println(i, count, next, h)
			if count+next >= k {
				for numOfHLeftToPrint := 0; numOfHLeftToPrint < h-i; numOfHLeftToPrint++ {
					ans += "H"
				}
				ans += "V"
				h = i
				v--
				break
			}
			count += next
		}
		if v == 0 {
			for i := 0; i < h; i++ {
				ans += "H"
			}
			break
		} else if h == 0 {
			for i := 0; i < v; i++ {
				ans += "V"
			}
			break
		}
	}

	return ans
}
