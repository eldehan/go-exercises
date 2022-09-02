package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder - NLR
// output - array of ints

// for each node, add its value to the subarr
// then, go to left node
// after left is done, go right

// Algorithm Preorder(tree)
//    1. Visit the root.
//    2. Traverse the left subtree, i.e., call Preorder(left-subtree)
//    3. Traverse the right subtree, i.e., call Preorder(right-subtree)

func preorderTraversal(root *TreeNode) []int {
	results := []int{}

	return traverse(root, &results)
}

func traverse(node *TreeNode, results *[]int) []int {
	if node == nil {
		return *results
	}

	*results = append(*results, node.Val)
	traverse(node.Left, results)
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
