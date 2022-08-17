package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution1(root *TreeNode) int {
	globalMax := math.MinInt64
	getMaxPathSum(root, &globalMax)
	return globalMax
}

// Chose to pass two values in instead of having a global variable
func getMaxPathSum(root *TreeNode, globalMax *int) int {
	if root == nil {
		return 0
	}
	left := max(getMaxPathSum(root.Left, globalMax), 0)
	right := max(getMaxPathSum(root.Right, globalMax), 0)
	*globalMax = max(*globalMax, root.Val+left+right)
	return root.Val + max(left, right)
}
