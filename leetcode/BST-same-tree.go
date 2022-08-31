// call recursive func to compare
// each call, check if treeNode1.Val == treeNode2.Val && compareFunc(treeNode1.left, treeNode2.left) and compareFunc(treeNode1.right, treeNode2.right)

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
