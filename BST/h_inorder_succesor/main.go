package inordersuccesorwithparentnode

// Node represents a BST node with parent pointer
type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// NewNode creates a new node
func NewNode(key int) *Node {
	return &Node{
		Key:    key,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}

// FindInorderSuccessor finds the next inorder successor of target node
// Time: O(h), Space: O(1) where h = height of tree
func FindInorderSuccessor(target *Node) *Node {
	if target == nil {
		return nil
	}

	// Case 1: Target has RIGHT subtree
	// Successor is the leftmost (minimum) node in right subtree
	if target.Right != nil {
		return findMinInSubtree(target.Right)
	}

	// Case 2: Target has NO right subtree
	// Go up using parent pointers until we find a node that is
	// a LEFT child of its parent (or reach nil)
	current := target
	parent := target.Parent

	// Keep going up while current is the RIGHT child of its parent
	for parent != nil && current == parent.Right {
		current = parent
		parent = parent.Parent
	}

	// parent is either nil (no successor) or the successor
	return parent
}

// findMinInSubtree finds the minimum node in a subtree
func findMinInSubtree(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Keep going left until we can't
	for node.Left != nil {
		node = node.Left
	}
	return node
}
