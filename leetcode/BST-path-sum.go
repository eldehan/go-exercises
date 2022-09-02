package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	found := false

	if root == nil {
		return found
	} else {
		traverse(root, targetSum, 0, &found)
	}

	return found
}

func traverse(node *TreeNode, targetSum int, total int, found *bool) {
	total += node.Val

	if node.Left == nil && node.Right == nil {
		if total == targetSum {
			*found = true
		}
	}

	if node.Left != nil {
		traverse(node.Left, targetSum, total, found)
	}
	if node.Right != nil {
		traverse(node.Right, targetSum, total, found)
	}
}
