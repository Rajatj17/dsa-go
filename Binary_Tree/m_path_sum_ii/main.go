package main

/*
Problem: Path Sum II

Description:
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum.

A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: []

Example 3:
Input: root = [1,2], targetSum = 0
Output: []

Constraints:
- The number of nodes in the tree is in the range [0, 5000]
- -1000 <= Node.val <= 1000
- -1000 <= targetSum <= 1000

LeetCode Link: https://leetcode.com/problems/path-sum-ii/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(N^2) in worst case where N is number of nodes
Space Complexity: O(N) for recursion stack and path storage

Related Problems:
- Path Sum
- Path Sum III
- Binary Tree Maximum Path Sum
- Sum Root to Leaf Numbers

Note: Backtracking in Tree:
1. Use DFS to traverse all root-to-leaf paths
2. Keep track of current path and current sum
3. When reaching a leaf, check if sum equals target
4. Use backtracking to explore all paths

Key insight: Use list to track current path and backtrack by removing last element when returning from recursion.
*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Approach 1: Backtracking with path tracking
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	
	dfs(root, targetSum, path, &result)
	return result
}

func dfs(node *TreeNode, targetSum int, path []int, result *[][]int) {
	if node == nil {
		return
	}
	
	// Add current node to path
	path = append(path, node.Val)
	
	// Check if it's a leaf node and sum equals target
	if node.Left == nil && node.Right == nil && targetSum == node.Val {
		// Make a copy of the path and add to result
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
	} else {
		// Continue DFS on left and right subtrees
		dfs(node.Left, targetSum-node.Val, path, result)
		dfs(node.Right, targetSum-node.Val, path, result)
	}
	
	// Backtrack: remove current node from path
	path = path[:len(path)-1]
}

// Approach 2: Optimized with pre-allocated slice
func pathSumOptimized(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	maxDepth := getMaxDepth(root)
	path := make([]int, 0, maxDepth) // Pre-allocate based on max depth
	
	dfsOptimized(root, targetSum, path, &result)
	return result
}

func dfsOptimized(node *TreeNode, targetSum int, path []int, result *[][]int) {
	if node == nil {
		return
	}
	
	// Add current node to path
	path = append(path, node.Val)
	
	// Check if it's a leaf node and sum equals target
	if node.Left == nil && node.Right == nil && targetSum == node.Val {
		// Make a copy of the path and add to result
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
	} else {
		// Continue DFS on left and right subtrees
		dfsOptimized(node.Left, targetSum-node.Val, path, result)
		dfsOptimized(node.Right, targetSum-node.Val, path, result)
	}
	
	// Backtrack: remove current node from path
	path = path[:len(path)-1]
}

// Approach 3: Without explicit backtracking (using slice copying)
func pathSumWithoutBacktrack(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	
	dfsWithoutBacktrack(root, targetSum, path, &result)
	return result
}

func dfsWithoutBacktrack(node *TreeNode, targetSum int, path []int, result *[][]int) {
	if node == nil {
		return
	}
	
	// Create new path with current node
	newPath := append(path, node.Val)
	
	// Check if it's a leaf node and sum equals target
	if node.Left == nil && node.Right == nil && targetSum == node.Val {
		// Add copy of path to result
		pathCopy := make([]int, len(newPath))
		copy(pathCopy, newPath)
		*result = append(*result, pathCopy)
	} else {
		// Continue DFS on left and right subtrees
		dfsWithoutBacktrack(node.Left, targetSum-node.Val, newPath, result)
		dfsWithoutBacktrack(node.Right, targetSum-node.Val, newPath, result)
	}
}

// Approach 4: Iterative using stack
func pathSumIterative(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	
	type stackItem struct {
		node *TreeNode
		path []int
		sum  int
	}
	
	stack := []stackItem{{root, []int{}, 0}}
	
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		node := item.node
		path := item.path
		sum := item.sum
		
		// Add current node to path and sum
		newPath := append(path, node.Val)
		newSum := sum + node.Val
		
		// Check if it's a leaf node and sum equals target
		if node.Left == nil && node.Right == nil && newSum == targetSum {
			pathCopy := make([]int, len(newPath))
			copy(pathCopy, newPath)
			result = append(result, pathCopy)
		} else {
			// Add children to stack
			if node.Right != nil {
				stack = append(stack, stackItem{node.Right, newPath, newSum})
			}
			if node.Left != nil {
				stack = append(stack, stackItem{node.Left, newPath, newSum})
			}
		}
	}
	
	return result
}

// Helper function to get maximum depth of tree
func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getMaxDepth(root.Left), getMaxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to create sample tree
func createSampleTree() *TreeNode {
	return &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}
}

// Helper function to create edge case trees
func createEdgeCaseTree1() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
}

func createEdgeCaseTree2() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2},
	}
}

func createSingleNodeTree() *TreeNode {
	return &TreeNode{Val: 1}
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

// Helper function to validate path sum
func validatePathSum(path []int, targetSum int) bool {
	sum := 0
	for _, val := range path {
		sum += val
	}
	return sum == targetSum
}

// Helper function to find all paths (not just sum-matching ones)
func findAllPaths(root *TreeNode) [][]int {
	var result [][]int
	var path []int
	
	findAllPathsHelper(root, path, &result)
	return result
}

func findAllPathsHelper(node *TreeNode, path []int, result *[][]int) {
	if node == nil {
		return
	}
	
	path = append(path, node.Val)
	
	if node.Left == nil && node.Right == nil {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
	} else {
		findAllPathsHelper(node.Left, path, result)
		findAllPathsHelper(node.Right, path, result)
	}
	
	path = path[:len(path)-1]
}

// Helper function to count total nodes and leaves
func getTreeStats(root *TreeNode) (nodes, leaves int) {
	if root == nil {
		return 0, 0
	}
	
	if root.Left == nil && root.Right == nil {
		return 1, 1
	}
	
	leftNodes, leftLeaves := getTreeStats(root.Left)
	rightNodes, rightLeaves := getTreeStats(root.Right)
	
	return 1 + leftNodes + rightNodes, leftLeaves + rightLeaves
}

func main() {
	fmt.Println("=== Binary Tree Path Sum II ===")
	
	// Test Case 1: Standard example
	fmt.Println("\n1. Standard Example:")
	tree1 := createSampleTree()
	printTree(tree1, "", true)
	
	targetSum := 22
	nodes, leaves := getTreeStats(tree1)
	fmt.Printf("\nTree stats: %d nodes, %d leaves\n", nodes, leaves)
	fmt.Printf("Target sum: %d\n", targetSum)
	
	fmt.Println("\nAll root-to-leaf paths:")
	allPaths := findAllPaths(tree1)
	for i, path := range allPaths {
		sum := 0
		for _, val := range path {
			sum += val
		}
		fmt.Printf("Path %d: %v (sum=%d)\n", i+1, path, sum)
	}
	
	fmt.Printf("\nApproach 1 - Backtracking (target=%d):\n", targetSum)
	result1 := pathSum(tree1, targetSum)
	for i, path := range result1 {
		fmt.Printf("Path %d: %v (valid=%v)\n", i+1, path, validatePathSum(path, targetSum))
	}
	
	fmt.Printf("\nApproach 2 - Optimized (target=%d):\n", targetSum)
	result2 := pathSumOptimized(tree1, targetSum)
	for i, path := range result2 {
		fmt.Printf("Path %d: %v (valid=%v)\n", i+1, path, validatePathSum(path, targetSum))
	}
	
	fmt.Printf("\nApproach 3 - Without Backtrack (target=%d):\n", targetSum)
	result3 := pathSumWithoutBacktrack(tree1, targetSum)
	for i, path := range result3 {
		fmt.Printf("Path %d: %v (valid=%v)\n", i+1, path, validatePathSum(path, targetSum))
	}
	
	fmt.Printf("\nApproach 4 - Iterative (target=%d):\n", targetSum)
	result4 := pathSumIterative(tree1, targetSum)
	for i, path := range result4 {
		fmt.Printf("Path %d: %v (valid=%v)\n", i+1, path, validatePathSum(path, targetSum))
	}
	
	// Test Case 2: Edge cases
	fmt.Println("\n2. Edge Cases:")
	
	// Empty tree
	fmt.Println("\nEmpty tree:")
	emptyResult := pathSum(nil, 5)
	fmt.Printf("Paths found: %d\n", len(emptyResult))
	
	// Single node (matches target)
	fmt.Println("\nSingle node (matches target):")
	singleTree := createSingleNodeTree()
	printTree(singleTree, "", true)
	singleResult := pathSum(singleTree, 1)
	fmt.Printf("Paths found: %d\n", len(singleResult))
	for i, path := range singleResult {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}
	
	// Single node (doesn't match target)
	fmt.Println("\nSingle node (doesn't match target):")
	singleResultNo := pathSum(singleTree, 5)
	fmt.Printf("Paths found: %d\n", len(singleResultNo))
	
	// No valid paths
	fmt.Println("\nNo valid paths:")
	tree2 := createEdgeCaseTree1()
	printTree(tree2, "", true)
	noPathResult := pathSum(tree2, 10)
	fmt.Printf("Paths found: %d\n", len(noPathResult))
	
	// Multiple valid paths
	fmt.Println("\nMultiple valid paths:")
	complexTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 3},
		},
	}
	printTree(complexTree, "", true)
	multipleResult := pathSum(complexTree, 6)
	fmt.Printf("Paths found: %d\n", len(multipleResult))
	for i, path := range multipleResult {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}
	
	// Test Case 3: Performance comparison
	fmt.Println("\n3. Performance Comparison:")
	fmt.Printf("Tree max depth: %d\n", getMaxDepth(tree1))
	fmt.Printf("Total paths found: %d\n", len(result1))
	fmt.Printf("Results match across approaches: %v\n", 
		len(result1) == len(result2) && len(result2) == len(result3) && len(result3) == len(result4))
}