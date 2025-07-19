package main

/*
Problem: Validate Binary Search Tree

Description:
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.

Constraints:
- The number of nodes in the tree is in the range [1, 10^4]
- -2^31 <= Node.val <= 2^31 - 1

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(h) where h is the height of the tree

Related Problems:
- Binary Tree Inorder Traversal
- Find Mode in Binary Search Tree
- Two Sum IV - Input is a BST

Note: Multiple approaches:
1. Inorder traversal: Should be sorted for a valid BST
2. Recursive with bounds: Pass min/max bounds to each recursive call
3. Iterative with stack: Similar to recursive but using explicit stack

Key insight: For each node, all nodes in left subtree must be < node.val, and all nodes in right subtree must be > node.val.
*/

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Approach 1: Recursive with bounds (Min/Max approach)
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	
	if node.Val <= min || node.Val >= max {
		return false
	}
	
	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}

// Approach 2: Inorder traversal (should be sorted for valid BST)
func isValidBSTInorder(root *TreeNode) bool {
	var inorder []int
	inorderTraversal(root, &inorder)
	
	for i := 1; i < len(inorder); i++ {
		if inorder[i] <= inorder[i-1] {
			return false
		}
	}
	return true
}

func inorderTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	
	inorderTraversal(node.Left, result)
	*result = append(*result, node.Val)
	inorderTraversal(node.Right, result)
}

// Approach 3: Optimized inorder with previous value tracking
func isValidBSTOptimized(root *TreeNode) bool {
	var prev *int
	return inorderValidate(root, &prev)
}

func inorderValidate(node *TreeNode, prev **int) bool {
	if node == nil {
		return true
	}
	
	if !inorderValidate(node.Left, prev) {
		return false
	}
	
	if *prev != nil && node.Val <= **prev {
		return false
	}
	*prev = &node.Val
	
	return inorderValidate(node.Right, prev)
}

// Approach 4: Iterative with stack
func isValidBSTIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}
	
	stack := []*TreeNode{}
	var prev *int
	current := root
	
	for current != nil || len(stack) > 0 {
		// Go to leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		
		// Pop and process
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if prev != nil && current.Val <= *prev {
			return false
		}
		prev = &current.Val
		
		current = current.Right
	}
	
	return true
}

// Helper function to create a sample BST
func createValidBST() *TreeNode {
	return &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{Val: 7},
			Right: &TreeNode{
				Val: 9,
				Right: &TreeNode{Val: 10},
			},
		},
	}
}

// Helper function to create an invalid BST
func createInvalidBST() *TreeNode {
	return &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{Val: 6}, // Invalid: 6 > 5 in left subtree
		},
		Right: &TreeNode{
			Val: 4, // Invalid: 4 < 5 in right subtree
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
	}
}

// Helper function to print tree structure
func printTree(root *TreeNode, prefix string, isLast bool) {
	if root == nil {
		return
	}
	
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	
	fmt.Printf("%s%s%d\n", prefix, connector, root.Val)
	
	nextPrefix := prefix
	if isLast {
		nextPrefix += "    "
	} else {
		nextPrefix += "│   "
	}
	
	if root.Left != nil || root.Right != nil {
		if root.Right != nil {
			printTree(root.Right, nextPrefix, root.Left == nil)
		}
		if root.Left != nil {
			printTree(root.Left, nextPrefix, true)
		}
	}
}

// Helper function to get inorder traversal for display
func getInorderTraversal(root *TreeNode) []int {
	var result []int
	inorderTraversal(root, &result)
	return result
}

func main() {
	fmt.Println("=== Binary Search Tree Validation ===")
	
	// Test Case 1: Valid BST
	fmt.Println("\n1. Valid BST:")
	validBST := createValidBST()
	printTree(validBST, "", true)
	
	fmt.Printf("\nInorder traversal: %v\n", getInorderTraversal(validBST))
	fmt.Printf("Approach 1 (Recursive with bounds): %v\n", isValidBST(validBST))
	fmt.Printf("Approach 2 (Inorder array): %v\n", isValidBSTInorder(validBST))
	fmt.Printf("Approach 3 (Optimized inorder): %v\n", isValidBSTOptimized(validBST))
	fmt.Printf("Approach 4 (Iterative): %v\n", isValidBSTIterative(validBST))
	
	// Test Case 2: Invalid BST
	fmt.Println("\n2. Invalid BST:")
	invalidBST := createInvalidBST()
	printTree(invalidBST, "", true)
	
	fmt.Printf("\nInorder traversal: %v\n", getInorderTraversal(invalidBST))
	fmt.Printf("Approach 1 (Recursive with bounds): %v\n", isValidBST(invalidBST))
	fmt.Printf("Approach 2 (Inorder array): %v\n", isValidBSTInorder(invalidBST))
	fmt.Printf("Approach 3 (Optimized inorder): %v\n", isValidBSTOptimized(invalidBST))
	fmt.Printf("Approach 4 (Iterative): %v\n", isValidBSTIterative(invalidBST))
	
	// Test Case 3: Edge cases
	fmt.Println("\n3. Edge Cases:")
	
	// Empty tree
	fmt.Printf("Empty tree: %v\n", isValidBST(nil))
	
	// Single node
	singleNode := &TreeNode{Val: 1}
	fmt.Printf("Single node: %v\n", isValidBST(singleNode))
	
	// Two nodes (valid)
	twoNodesValid := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 1},
	}
	fmt.Printf("Two nodes (valid): %v\n", isValidBST(twoNodesValid))
	
	// Two nodes (invalid)
	twoNodesInvalid := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2},
	}
	fmt.Printf("Two nodes (invalid): %v\n", isValidBST(twoNodesInvalid))
	
	// Common mistake: [2,1,3] vs [5,1,4,null,null,3,6]
	fmt.Println("\n4. Common Test Cases:")
	
	// Valid: [2,1,3]
	validCase := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	fmt.Printf("[2,1,3]: %v\n", isValidBST(validCase))
	
	// Invalid: [5,1,4,null,null,3,6]
	invalidCase := &TreeNode{
		Val: 5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}
	fmt.Printf("[5,1,4,null,null,3,6]: %v\n", isValidBST(invalidCase))
}