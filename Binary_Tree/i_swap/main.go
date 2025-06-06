package swap

// The Sibling Swap problem asks:
// Can we transform one binary tree into another by swapping left and right children at any nodes?

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// canConvertBySwapping determines if tree1 can be converted to tree2 by swapping children
func CanConvertBySwapping(root1, root2 *TreeNode) bool {
	// Base case: both nodes are nil
	if root1 == nil && root2 == nil {
		return true
	}

	// Base case: one is nil, other is not
	if root1 == nil || root2 == nil {
		return false
	}

	// Values must match at current nodes
	if root1.Val != root2.Val {
		return false
	}

	// Two possibilities at each node:
	// 1. No swap: left→left, right→right
	// 2. Swap: left→right, right→left
	noSwap := CanConvertBySwapping(root1.Left, root2.Left) &&
		CanConvertBySwapping(root1.Right, root2.Right)

	withSwap := CanConvertBySwapping(root1.Left, root2.Right) &&
		CanConvertBySwapping(root1.Right, root2.Left)

	// Return true if either configuration works
	return noSwap || withSwap
}
