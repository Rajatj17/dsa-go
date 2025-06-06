package leftviewtree

// Example Tree:
//        1      <- Level 0: First node = 1
//       / \
//      2   3    <- Level 1: First node = 2
//     / \   \
//    4   5   6  <- Level 2: First node = 4
//       /
//      8        <- Level 3: First node = 8
// Left View Result: [1, 2, 4, 8]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method 1: Using BFS (Level Order Traversal)
func LeftViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		// Process all nodes at current level
		for i := range levelSize {
			node := queue[0]
			queue = queue[1:]

			// First node of each level is part of left view
			if i == 0 {
				result = append(result, node.Val)
			}

			// Add children for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

// Method 2: Using DFS (Preorder Traversal)
func LeftViewDFS(root *TreeNode) []int {
	result := []int{}
	leftViewDFSHelper(root, 0, &result)
	return result
}

func leftViewDFSHelper(node *TreeNode, level int, result *[]int) {
	if node == nil {
		return
	}

	// If this is the first node at this level, add to result
	if level == len(*result) {
		*result = append(*result, node.Val)
	}

	// Traverse left first, then right
	leftViewDFSHelper(node.Left, level+1, result)
	leftViewDFSHelper(node.Right, level+1, result)
}

// Method 3: Using map to track levels
func LeftViewWithMap(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	levelMap := make(map[int]int) // level -> first node value
	maxLevel := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		// If this level hasn't been seen before, record this node
		if _, exists := levelMap[level]; !exists {
			levelMap[level] = node.Val
		}

		if level > maxLevel {
			maxLevel = level
		}

		// Visit left first, then right
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	// Build result array
	result := make([]int, maxLevel+1)
	for level := 0; level <= maxLevel; level++ {
		result[level] = levelMap[level]
	}

	return result
}
