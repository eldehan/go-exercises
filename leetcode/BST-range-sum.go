package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	return helper(root, low, high, &sum)
}

func helper(node *TreeNode, low, high int, sum *int) int {
	if node == nil {
		return 0
	}

	if node.Val <= high && node.Val >= low {
		*sum += node.Val
	}

	helper(node.Left, low, high, sum)
	helper(node.Right, low, high, sum)

	return *sum
}
