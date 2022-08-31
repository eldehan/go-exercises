// if one side has children and the other doesn't
// the other side's children cannot have any more children (return false)
// helper func takes left and right
// if left and right are

package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return -1 != balancedHelper(root)
}

func balancedHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftLength := balancedHelper(node.Left)
	rightLength := balancedHelper(node.Right)

	if leftLength == -1 || rightLength == -1 || math.Abs(float64(leftLength-rightLength)) > 1 {
		return -1
	}
	return 1 + int(math.Max(float64(leftLength), float64(rightLength)))
}
