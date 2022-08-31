// given the root of a binary tree, check whether it is a MIRROR of itself (symmetric around center)

// if at any point left or right are null (but NOT both) return false

// if left val == right val (current args of recursive func) AND
// if left's right and right's left are mirror
//      (recursive call with left.right and right.left passed in) AND
// if left's left and right's right are mirror
//      (recursive call with left.left and right.right passed in)

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	left, right := root.Left, root.Right

	return symmetricHelper(left, right)
}

func symmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && symmetricHelper(left.Right, right.Left) && symmetricHelper(left.Left, right.Right)
}
