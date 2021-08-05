package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	preorderTraversal(&tree)
}

// 前序递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 104 4ms 90.80%
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	// 判断左右谁大,返回最大的
	if left > right {
		return left + 1
	}
	return right + 1
}
