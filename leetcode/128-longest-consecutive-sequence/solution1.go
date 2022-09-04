package main

func longestConsecutive(nums []int) int {
	setOfNumbers := make(map[int]bool)

	for _, num := range nums {
		setOfNumbers[num] = false
	}

	longestChainLength := 0

	for num, _ := range setOfNumbers {
		if _, ok := setOfNumbers[num-1]; ok {
			continue
		}

		chainLength := 1
		for {
			_, ok := setOfNumbers[num+chainLength]
			if ok {
				setOfNumbers[num+chainLength] = true
				chainLength++
			} else {
				longestChainLength = max(chainLength, longestChainLength)
				break
			}
		}
	}
	return longestChainLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
