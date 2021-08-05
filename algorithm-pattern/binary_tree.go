package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg
	// [3,9,20,null,null,15,7]
	tree := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	_ = maxDepth(&tree)
	result := make([]int, 0)
	resultRe := preorderTraversal(&tree, &result)
	resultNoRe := preorderTraversalNoRecursive(&tree)
	fmt.Println(resultRe)
	fmt.Println(resultNoRe)
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

// 前序递归
func preorderTraversal(root *TreeNode, result *[]int) []int {
	if root == nil {
		return nil
	}
	*result = append(*result, root.Val)
	preorderTraversal(root.Left, result)
	preorderTraversal(root.Right, result)
	return *result
}

// 前序非递归
func preorderTraversalNoRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	return result
}
