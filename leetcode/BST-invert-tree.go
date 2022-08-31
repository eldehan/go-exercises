// beginning with root, call invert on root
// in invert - temp = root.left, root.left = root.right, root.right = temp
// call func again on kids
// if func is called with nil as node, just return

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(node *TreeNode) {
	if node == nil {
		return
	}

	temp := node.Left
	node.Left = node.Right
	node.Right = temp

	invert(node.Left)
	invert(node.Right)
}
