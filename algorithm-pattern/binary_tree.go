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

// DFS 深度搜索 从上到下
func preorderTraversalDfsUpToDown(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

// 传入指针
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// 110 平衡二叉树    4 ms, faster than 95.22%
func isBalanced(root *TreeNode) bool {
	if maxDepthBalanced(root) == -1 {
		return false
	}
	return true
}

func maxDepthBalanced(root *TreeNode) int {
	// check
	if root == nil {
		return 0
	}
	left := maxDepthBalanced(root.Left)
	right := maxDepthBalanced(root.Right)

	// 高度差大于1 返回 -1
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// 124 二叉树中的最大路径和  16 ms, faster than 87.97%
// Binary Tree Maximum Path Sum

// ResultType 保存最大值
type ResultType struct {
	SinglePath int //单边最大值
	MaxPath    int // 最大值
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	// root为空
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 31),
		}
	}
	// 分治
	left := helper(root.Left)
	right := helper(root.Right)

	result := ResultType{}
	// 求单边最大值
	if left.SinglePath > right.SinglePath {
		result.SinglePath = maxInt(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = maxInt(right.SinglePath+root.Val, 0)
	}
	// 求两遍加根最大值
	maxPath := maxInt(left.MaxPath, right.MaxPath)
	result.MaxPath = maxInt(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 236.二叉树的最近公共祖先 4 ms, faster than 99.83%
// lowest-common-ancestor-of-a-binary-tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 是否为空
	if root == nil {
		return root
	}

	// 相等
	if root == p || root == q {
		return root
	}

	// 分治
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 左右两边都不为空,则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// 102. 二叉树的层序遍历  0 ms, faster than 100.00%
// Binary Tree Level Order Traversal
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

// 107. 二叉树的层序遍历 II 0 ms, faster than 100.00%
// 107. Binary Tree Level Order Traversal II
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	// 翻转结果
	reverse(result)
	return result
}
func reverse(nums [][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 103. 二叉树的锯齿形层序遍历 0 ms, faster than 100.00%
// Binary Tree Zigzag Level Order Traversal
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	toggle := false
	for len(queue) > 0 {
		list := make([]int, 0)
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		if toggle {
			reverseOne(list)
		}
		result = append(result, list)
		toggle = !toggle
	}
	return result
}

func reverseOne(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

// 98. 验证二叉搜索树  4 ms, faster than 96.22%
// 98. Validate Binary Search Tree
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)
	// check order
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}
