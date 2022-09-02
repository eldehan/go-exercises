package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Algorithm Inorder(tree)
//    1. Traverse the left subtree, i.e., call Inorder(left-subtree)
//    2. Visit the root.
//    3. Traverse the right subtree, i.e., call Inorder(right-subtree)

func inorderTraversal(root *TreeNode) []int {
	results := []int{}

	return traverse(root, &results)
}

func traverse(node *TreeNode, results *[]int) []int {
	if node == nil {
		return *results
	}

	traverse(node.Left, results)
	*results = append(*results, node.Val)
	traverse(node.Right, results)
	return *results
}

// without pointers...

// func preorderTraversal(root *TreeNode) (values []int) {
//   if root == nil {
//     return
//   }
//   values = append(values, root.Val)
//   values = append(values, preorderTraversal(root.Left)...)
//   values = append(values, preorderTraversal(root.Right)...)
//   return
// }
