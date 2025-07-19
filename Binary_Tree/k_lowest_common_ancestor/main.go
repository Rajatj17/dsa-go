package main

/*
Problem: Lowest Common Ancestor of a Binary Tree

Description:
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself)."

Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.

Example 2:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

Example 3:
Input: root = [1,2], p = 1, q = 2
Output: 1

Constraints:
- The number of nodes in the tree is in the range [2, 10^5]
- -10^9 <= Node.val <= 10^9
- All Node.val are unique
- p != q
- p and q will exist in the tree

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(h) where h is the height of the tree

Related Problems:
- Lowest Common Ancestor of a Binary Search Tree
- Smallest Common Region
- Lowest Common Ancestor of Deepest Leaves

Note: Recursive approach:
1. If current node is null, return null
2. If current node is p or q, return current node
3. Recursively find LCA in left and right subtrees
4. If both left and right return non-null, current node is LCA
5. Otherwise, return the non-null subtree result

Key insight: The LCA is the first node where p and q are found in different subtrees.
*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Approach 1: Recursive DFS (Most common approach)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	
	if left != nil && right != nil {
		return root
	}
	
	if left != nil {
		return left
	}
	return right
}

// Approach 2: Using parent pointers (path tracking)
func lowestCommonAncestorWithPath(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	// Find paths from root to both nodes
	var pathP []int
	var pathQ []int
	
	if !findPath(root, p.Val, &pathP) || !findPath(root, q.Val, &pathQ) {
		return nil
	}
	
	// Find LCA from paths
	var lca int
	i := 0
	for i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i] {
		lca = pathP[i]
		i++
	}
	
	// Find node with LCA value
	return findNodeByValue(root, lca)
}

func findPath(root *TreeNode, target int, path *[]int) bool {
	if root == nil {
		return false
	}
	
	*path = append(*path, root.Val)
	
	if root.Val == target {
		return true
	}
	
	if findPath(root.Left, target, path) || findPath(root.Right, target, path) {
		return true
	}
	
	*path = (*path)[:len(*path)-1]
	return false
}

func findNodeByValue(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	
	if root.Val == val {
		return root
	}
	
	left := findNodeByValue(root.Left, val)
	if left != nil {
		return left
	}
	
	return findNodeByValue(root.Right, val)
}

// Approach 3: Iterative using stack and parent tracking
func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	// Map to store parent of each node
	parent := make(map[*TreeNode]*TreeNode)
	parent[root] = nil
	
	// Stack for DFS
	stack := []*TreeNode{root}
	
	// Continue until we find both p and q
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if node.Left != nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}
		
		if node.Right != nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
		
		// Check if we found both nodes
		if _, foundP := parent[p]; foundP {
			if _, foundQ := parent[q]; foundQ {
				break
			}
		}
	}
	
	// Get all ancestors of p
	ancestors := make(map[*TreeNode]bool)
	current := p
	for current != nil {
		ancestors[current] = true
		current = parent[current]
	}
	
	// Find first common ancestor of q
	current = q
	for current != nil {
		if ancestors[current] {
			return current
		}
		current = parent[current]
	}
	
	return nil
}

// Approach 4: Using post-order traversal with result tracking
func lowestCommonAncestorPostOrder(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode
	postOrder(root, p, q, &result)
	return result
}

func postOrder(node, p, q *TreeNode, result **TreeNode) bool {
	if node == nil {
		return false
	}
	
	left := postOrder(node.Left, p, q, result)
	right := postOrder(node.Right, p, q, result)
	
	current := node == p || node == q
	
	if (left && right) || (current && (left || right)) {
		if *result == nil {
			*result = node
		}
	}
	
	return current || left || right
}

// Helper function to create sample tree
func createSampleTree() *TreeNode {
	return &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
}

// Helper function to find node by value
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	
	if root.Val == val {
		return root
	}
	
	left := findNode(root.Left, val)
	if left != nil {
		return left
	}
	
	return findNode(root.Right, val)
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

// Helper function to get path from root to node
func getPathString(root *TreeNode, target int) string {
	var path []int
	if findPath(root, target, &path) {
		result := ""
		for i, val := range path {
			if i > 0 {
				result += " -> "
			}
			result += fmt.Sprintf("%d", val)
		}
		return result
	}
	return "Not found"
}

func main() {
	fmt.Println("=== Lowest Common Ancestor of Binary Tree ===")
	
	// Create sample tree
	root := createSampleTree()
	
	fmt.Println("\nTree structure:")
	printTree(root, "", true)
	
	// Test cases
	testCases := []struct {
		pVal, qVal int
		desc       string
	}{
		{5, 1, "nodes in different subtrees"},
		{5, 4, "one node is ancestor of another"},
		{6, 4, "nodes in same subtree"},
		{0, 8, "leaf nodes in different subtrees"},
		{7, 4, "nodes in same subtree (deeper)"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("\n%d. Test case: %s\n", i+1, tc.desc)
		
		p := findNode(root, tc.pVal)
		q := findNode(root, tc.qVal)
		
		if p == nil || q == nil {
			fmt.Printf("One or both nodes not found\n")
			continue
		}
		
		fmt.Printf("Finding LCA of %d and %d:\n", tc.pVal, tc.qVal)
		fmt.Printf("Path to %d: %s\n", tc.pVal, getPathString(root, tc.pVal))
		fmt.Printf("Path to %d: %s\n", tc.qVal, getPathString(root, tc.qVal))
		
		// Test all approaches
		lca1 := lowestCommonAncestor(root, p, q)
		lca2 := lowestCommonAncestorWithPath(root, p, q)
		lca3 := lowestCommonAncestorIterative(root, p, q)
		lca4 := lowestCommonAncestorPostOrder(root, p, q)
		
		fmt.Printf("Approach 1 (Recursive): %d\n", lca1.Val)
		fmt.Printf("Approach 2 (Path tracking): %d\n", lca2.Val)
		fmt.Printf("Approach 3 (Iterative): %d\n", lca3.Val)
		fmt.Printf("Approach 4 (Post-order): %d\n", lca4.Val)
	}
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Root and leaf
	root_node := findNode(root, 3)
	leaf_node := findNode(root, 6)
	lca := lowestCommonAncestor(root, root_node, leaf_node)
	fmt.Printf("LCA of root(3) and leaf(6): %d\n", lca.Val)
	
	// Same node
	same_node := findNode(root, 5)
	lca_same := lowestCommonAncestor(root, same_node, same_node)
	fmt.Printf("LCA of same node(5): %d\n", lca_same.Val)
	
	// Adjacent nodes
	parent_node := findNode(root, 2)
	child_node := findNode(root, 7)
	lca_adjacent := lowestCommonAncestor(root, parent_node, child_node)
	fmt.Printf("LCA of parent(2) and child(7): %d\n", lca_adjacent.Val)
}