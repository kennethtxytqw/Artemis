package main

import (
	"math"
)

func solution1(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	total_left_size := (len(nums1) + len(nums2) + 1) / 2.0

	low := -1
	high := len(nums1) - 1

	for {
		l1_index := (high + low) / 2
		l1, l2 := math.MinInt64, math.MinInt64
		r1, r2 := math.MaxInt64, math.MaxInt64
		if l1_index >= 0 {
			l1 = nums1[l1_index]
		}
		if len(nums1) > l1_index+1 {
			r1 = nums1[l1_index+1]
		}
		l2_index := total_left_size - l1_index - 2

		if l2_index >= 0 && l2_index < len(nums2) {
			l2 = nums2[l2_index]
		}
		if len(nums2) > l2_index+1 && l2_index+1 >= 0 {
			r2 = nums2[l2_index+1]
		}
		if l1 > r2 {
			high = l1_index - 1
			continue
		} else if l2 > r1 {
			low = l1_index + 1
			continue
		} else {
			res := 0.0
			if l1 > l2 {
				res = float64(l1)
			} else {
				res = float64(l2)
			}
			if (len(nums1)+len(nums2))%2 != 0 {
				return res
			} else {
				if r1 < r2 {
					res += float64(r1)
				} else {
					res += float64(r2)
				}
				return res / 2.0
			}
		}
	}
}
