package main

import "math"
import "fmt"

func solution1(s string) int {
	fmt.Printf("test: %v\n", s)
	l := '('
	left := []float64{0}
	completed := []float64{0}
	// (()(((()
	for _, v := range s {
		if v != l {
			if left[len(left)-1] == 0 {
				left = append(left, 0)
				left = append(left, 0)
				completed = append(completed, 0)
				completed = append(completed, 0)
			} else if left[len(left)-1] > 0 {
				completed[len(completed)-1] += 1
				left[len(left)-1] -= 1
				if left[len(left)-1] == 0 {
					if len(left) > 1 {
						left = left[:len(left)-1]
					}
					if len(completed) > 1 {
						temp := completed[len(completed)-1]
						completed = completed[:len(completed)-1]
						completed[len(completed)-1] += temp
					}
				}
			}
		} else {
			if left[len(left)-1] >= 0 && completed[len(completed)-1] > 0 {
				left = append(left, 1)
				completed = append(completed, 0)
			} else if left[len(left)-1] >= 0 && completed[len(completed)-1] == 0 {
				left[len(left)-1] += 1
			}
		}

		fmt.Printf("left: %v\n", left)
		fmt.Printf("completed: %v\n", completed)
	}
	max_len := 0.0
	for _, v := range completed {
		max_len = math.Max(v, max_len)
	}
	return int(max_len) * 2
}
