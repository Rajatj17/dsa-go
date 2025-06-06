package isbst

import (
	"math"
)

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Approach 1: Range-based validation (Most Efficient)
// Time: O(n), Space: O(h) where h is height
func IsValidBSTRange(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}

	// Current node must be within the valid range
	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}

	// Recursively validate left and right subtrees with updated ranges
	return validate(node.Left, minVal, node.Val) &&
		validate(node.Right, node.Val, maxVal)
}

// Approach 2: Inorder Traversal (Clean and Intuitive)
// Time: O(n), Space: O(h)
func IsValidBSTInorder(root *TreeNode) bool {
	var prev *int
	return InorderCheck(root, &prev)
}

func InorderCheck(node *TreeNode, prev **int) bool {
	if node == nil {
		return true
	}

	// Check left subtree
	if !InorderCheck(node.Left, prev) {
		return false
	}

	// Check current node
	if *prev != nil && node.Val <= **prev {
		return false
	}
	*prev = &node.Val

	// Check right subtree
	return InorderCheck(node.Right, prev)
}
