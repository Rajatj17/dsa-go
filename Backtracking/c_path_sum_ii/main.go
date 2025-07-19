package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Problem: Find all root-to-leaf paths where each path's sum equals targetSum
// Time Complexity: O(N * H) where N is number of nodes, H is height (for copying paths)
// Space Complexity: O(H) for recursion stack, O(N * H) for storing all paths

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
	path := make([]int, 0, getMaxDepth(root)) // Pre-allocate based on max depth
	
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

func main() {
	// Example tree:
	//       5
	//      / \
	//     4   8
	//    /   / \
	//   11  13  4
	//  / \    / \
	// 7   2  5   1
	
	root := &TreeNode{
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
	
	fmt.Println("Tree structure:")
	printTree(root, "", true)
	
	targetSum := 22
	
	fmt.Printf("\n=== Backtracking Approach (target=%d) ===\n", targetSum)
	result1 := pathSum(root, targetSum)
	for i, path := range result1 {
		sum := 0
		for _, val := range path {
			sum += val
		}
		fmt.Printf("Path %d: %v (sum=%d)\n", i+1, path, sum)
	}
	
	fmt.Printf("\n=== Optimized Approach (target=%d) ===\n", targetSum)
	result2 := pathSumOptimized(root, targetSum)
	for i, path := range result2 {
		sum := 0
		for _, val := range path {
			sum += val
		}
		fmt.Printf("Path %d: %v (sum=%d)\n", i+1, path, sum)
	}
	
	fmt.Printf("\n=== Without Backtrack (target=%d) ===\n", targetSum)
	result3 := pathSumWithoutBacktrack(root, targetSum)
	for i, path := range result3 {
		sum := 0
		for _, val := range path {
			sum += val
		}
		fmt.Printf("Path %d: %v (sum=%d)\n", i+1, path, sum)
	}
	
	fmt.Printf("\nTotal paths found: %d\n", len(result1))
	fmt.Printf("Tree max depth: %d\n", getMaxDepth(root))
}