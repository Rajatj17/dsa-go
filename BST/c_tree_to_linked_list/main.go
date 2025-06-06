package treetolinkedlist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LinkedList struct {
	Head *TreeNode
	Tail *TreeNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
	}
}

func TreeToLinkedList(root *TreeNode) *LinkedList {
	l := NewLinkedList()

	if root == nil {
		return l
	}

	if root.Left == nil && root.Right == nil {
		l.Head = root
		l.Tail = root
		return l
	} else if root.Left != nil && root.Right == nil {
		leftLL := TreeToLinkedList(root.Left)
		leftLL.Tail.Right = root

		l.Head = leftLL.Head
		l.Tail = root
	} else if root.Left == nil && root.Right != nil {
		rightLL := TreeToLinkedList(root.Right)

		root.Right = rightLL.Head

		l.Head = root
		l.Tail = rightLL.Tail
	} else {
		leftLL := TreeToLinkedList(root.Left)
		rightLL := TreeToLinkedList(root.Right)

		leftLL.Tail.Right = root
		root.Right = rightLL.Head

		l.Head = root
		l.Tail = rightLL.Tail
	}

	return l
}

// Method 2: Flatten to right-skewed tree (preorder)
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// Recursively flatten left and right subtrees
	Flatten(root.Left)
	Flatten(root.Right)

	// Store right subtree
	rightSubtree := root.Right

	// Move left subtree to right
	root.Right = root.Left
	root.Left = nil

	// Find the rightmost node of moved subtree
	current := root
	for current.Right != nil {
		current = current.Right
	}

	// Attach original right subtree
	current.Right = rightSubtree
}
