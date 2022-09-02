package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	nodeCollection := []*TreeNode{}
	traverse(root, &nodeCollection)

	for i, node := range nodeCollection {
		node.Left = nil
		if i < len(nodeCollection)-1 {
			node.Right = nodeCollection[i+1]
		} else {
			node.Right = nil
		}
	}
}

func traverse(node *TreeNode, collection *[]*TreeNode) []*TreeNode {
	if node == nil {
		return *collection
	}

	*collection = append(*collection, node)
	traverse(node.Left, collection)
	traverse(node.Right, collection)
	return *collection
}
