package main

/*
Problem: Serialize and Deserialize Binary Tree

Description:
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Example 1:
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

Example 2:
Input: root = []
Output: []

Constraints:
- The number of nodes in the tree is in the range [0, 10^4]
- -1000 <= Node.val <= 1000

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n) for both serialize and deserialize
Space Complexity: O(n)

Related Problems:
- Encode and Decode Strings
- Serialize and Deserialize BST
- Find Duplicate Subtrees

Note: Multiple approaches:
1. Preorder traversal: Root -> Left -> Right (easier to reconstruct)
2. Level order traversal: BFS approach
3. Postorder traversal: Left -> Right -> Root

Key insight: Preorder traversal with null markers allows easy reconstruction.
During deserialization, we can build the tree in the same order as serialization.
*/

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec encapsulates the serialization and deserialization logic
type Codec struct{}

// Approach 1: Preorder traversal with null markers
func (c *Codec) serialize(root *TreeNode) string {
	var result []string
	c.serializeHelper(root, &result)
	return strings.Join(result, ",")
}

func (c *Codec) serializeHelper(node *TreeNode, result *[]string) {
	if node == nil {
		*result = append(*result, "null")
		return
	}
	
	*result = append(*result, strconv.Itoa(node.Val))
	c.serializeHelper(node.Left, result)
	c.serializeHelper(node.Right, result)
}

func (c *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	index := 0
	return c.deserializeHelper(values, &index)
}

func (c *Codec) deserializeHelper(values []string, index *int) *TreeNode {
	if *index >= len(values) || values[*index] == "null" {
		*index++
		return nil
	}
	
	val, _ := strconv.Atoi(values[*index])
	*index++
	
	node := &TreeNode{Val: val}
	node.Left = c.deserializeHelper(values, index)
	node.Right = c.deserializeHelper(values, index)
	
	return node
}

// Approach 2: Level-order traversal (BFS)
func (c *Codec) serializeBFS(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	
	var result []string
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		
		if node == nil {
			result = append(result, "null")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	
	return strings.Join(result, ",")
}

func (c *Codec) deserializeBFS(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	
	values := strings.Split(data, ",")
	if len(values) == 0 {
		return nil
	}
	
	rootVal, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	
	i := 1
	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]
		
		// Left child
		if i < len(values) && values[i] != "null" {
			leftVal, _ := strconv.Atoi(values[i])
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}
		i++
		
		// Right child
		if i < len(values) && values[i] != "null" {
			rightVal, _ := strconv.Atoi(values[i])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
		i++
	}
	
	return root
}

// Approach 3: Postorder traversal
func (c *Codec) serializePostOrder(root *TreeNode) string {
	var result []string
	c.serializePostOrderHelper(root, &result)
	return strings.Join(result, ",")
}

func (c *Codec) serializePostOrderHelper(node *TreeNode, result *[]string) {
	if node == nil {
		*result = append(*result, "null")
		return
	}
	
	c.serializePostOrderHelper(node.Left, result)
	c.serializePostOrderHelper(node.Right, result)
	*result = append(*result, strconv.Itoa(node.Val))
}

func (c *Codec) deserializePostOrder(data string) *TreeNode {
	values := strings.Split(data, ",")
	index := len(values) - 1
	return c.deserializePostOrderHelper(values, &index)
}

func (c *Codec) deserializePostOrderHelper(values []string, index *int) *TreeNode {
	if *index < 0 || values[*index] == "null" {
		*index--
		return nil
	}
	
	val, _ := strconv.Atoi(values[*index])
	*index--
	
	node := &TreeNode{Val: val}
	node.Right = c.deserializePostOrderHelper(values, index)
	node.Left = c.deserializePostOrderHelper(values, index)
	
	return node
}

// Approach 4: Compact serialization (without null markers where possible)
func (c *Codec) serializeCompact(root *TreeNode) string {
	if root == nil {
		return ""
	}
	
	var result []string
	c.serializeCompactHelper(root, &result)
	return strings.Join(result, ",")
}

func (c *Codec) serializeCompactHelper(node *TreeNode, result *[]string) {
	if node == nil {
		return
	}
	
	*result = append(*result, strconv.Itoa(node.Val))
	
	if node.Left != nil || node.Right != nil {
		*result = append(*result, "(")
		c.serializeCompactHelper(node.Left, result)
		*result = append(*result, ",")
		c.serializeCompactHelper(node.Right, result)
		*result = append(*result, ")")
	}
}

// Helper functions for testing
func createSampleTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
}

func createComplexTree() *TreeNode {
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

func compareTrees(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	
	return a.Val == b.Val && compareTrees(a.Left, b.Left) && compareTrees(a.Right, b.Right)
}

func getInorderTraversal(root *TreeNode) []int {
	var result []int
	inorderHelper(root, &result)
	return result
}

func inorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	
	inorderHelper(node.Left, result)
	*result = append(*result, node.Val)
	inorderHelper(node.Right, result)
}

func main() {
	fmt.Println("=== Binary Tree Serialization and Deserialization ===")
	
	codec := &Codec{}
	
	// Test Case 1: Simple tree
	fmt.Println("\n1. Simple Tree:")
	tree1 := createSampleTree()
	printTree(tree1, "", true)
	
	fmt.Println("\nApproach 1 - Preorder:")
	serialized1 := codec.serialize(tree1)
	fmt.Printf("Serialized: %s\n", serialized1)
	deserialized1 := codec.deserialize(serialized1)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(tree1, deserialized1))
	fmt.Printf("Original inorder: %v\n", getInorderTraversal(tree1))
	fmt.Printf("Deserialized inorder: %v\n", getInorderTraversal(deserialized1))
	
	fmt.Println("\nApproach 2 - Level-order (BFS):")
	serializedBFS1 := codec.serializeBFS(tree1)
	fmt.Printf("Serialized: %s\n", serializedBFS1)
	deserializedBFS1 := codec.deserializeBFS(serializedBFS1)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(tree1, deserializedBFS1))
	
	fmt.Println("\nApproach 3 - Postorder:")
	serializedPost1 := codec.serializePostOrder(tree1)
	fmt.Printf("Serialized: %s\n", serializedPost1)
	deserializedPost1 := codec.deserializePostOrder(serializedPost1)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(tree1, deserializedPost1))
	
	// Test Case 2: Complex tree
	fmt.Println("\n2. Complex Tree:")
	tree2 := createComplexTree()
	printTree(tree2, "", true)
	
	fmt.Println("\nApproach 1 - Preorder:")
	serialized2 := codec.serialize(tree2)
	fmt.Printf("Serialized: %s\n", serialized2)
	deserialized2 := codec.deserialize(serialized2)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(tree2, deserialized2))
	
	fmt.Println("\nApproach 2 - Level-order (BFS):")
	serializedBFS2 := codec.serializeBFS(tree2)
	fmt.Printf("Serialized: %s\n", serializedBFS2)
	deserializedBFS2 := codec.deserializeBFS(serializedBFS2)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(tree2, deserializedBFS2))
	
	// Test Case 3: Edge cases
	fmt.Println("\n3. Edge Cases:")
	
	// Empty tree
	fmt.Println("\nEmpty tree:")
	emptyTree := (*TreeNode)(nil)
	serializedEmpty := codec.serialize(emptyTree)
	fmt.Printf("Serialized: %s\n", serializedEmpty)
	deserializedEmpty := codec.deserialize(serializedEmpty)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(emptyTree, deserializedEmpty))
	
	// Single node
	fmt.Println("\nSingle node:")
	singleNode := &TreeNode{Val: 42}
	serializedSingle := codec.serialize(singleNode)
	fmt.Printf("Serialized: %s\n", serializedSingle)
	deserializedSingle := codec.deserialize(serializedSingle)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(singleNode, deserializedSingle))
	
	// Left-skewed tree
	fmt.Println("\nLeft-skewed tree:")
	leftSkewed := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 4},
			},
		},
	}
	serializedLeft := codec.serialize(leftSkewed)
	fmt.Printf("Serialized: %s\n", serializedLeft)
	deserializedLeft := codec.deserialize(serializedLeft)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(leftSkewed, deserializedLeft))
	
	// Right-skewed tree
	fmt.Println("\nRight-skewed tree:")
	rightSkewed := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{Val: 4},
			},
		},
	}
	serializedRight := codec.serialize(rightSkewed)
	fmt.Printf("Serialized: %s\n", serializedRight)
	deserializedRight := codec.deserialize(serializedRight)
	fmt.Printf("Deserialized matches original: %v\n", compareTrees(rightSkewed, deserializedRight))
	
	// Demonstrate compact serialization
	fmt.Println("\n4. Compact Serialization (Advanced):")
	compactSerialized := codec.serializeCompact(tree1)
	fmt.Printf("Compact serialized: %s\n", compactSerialized)
	
	// Performance comparison
	fmt.Println("\n5. Serialization Length Comparison:")
	fmt.Printf("Preorder: %d characters\n", len(serialized2))
	fmt.Printf("Level-order: %d characters\n", len(serializedBFS2))
	fmt.Printf("Postorder: %d characters\n", len(serializedPost1))
	fmt.Printf("Compact: %d characters\n", len(compactSerialized))
}