package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Could do recursion but it is just wasteful of resources
func findLCA(root, p, q *TreeNode) *TreeNode {
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
}

func solution1(root, p, q *TreeNode) *TreeNode {
	return findLCA(root, p, q)
}
