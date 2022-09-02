package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return traverse(root, root.Val)
}

// if a node's children are null, then return the current built up root-to-leaf path
func traverse(root *TreeNode, current int) int {
	if root.Left == nil && root.Right == nil {
		return current
	}

	total := 0
	if root.Left != nil {
		total += traverse(root.Left, current*10+root.Left.Val)
	}
	if root.Right != nil {
		total += traverse(root.Right, current*10+root.Right.Val)
	}

	return total
}
