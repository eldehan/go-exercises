package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Algorithm Postorder(tree)
//    1. Traverse the left subtree, i.e., call Postorder(left-subtree)
//    2. Traverse the right subtree, i.e., call Postorder(right-subtree)
//    3. Visit the root.

func postorderTraversal(root *TreeNode) []int {
	results := []int{}

	return traverse(root, &results)
}

func traverse(node *TreeNode, results *[]int) []int {
	if node == nil {
		return *results
	}

	traverse(node.Left, results)
	traverse(node.Right, results)
	*results = append(*results, node.Val)
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
