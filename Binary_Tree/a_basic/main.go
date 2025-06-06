package basic

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

type BST struct {
	Root *TreeNode
	Size int
}

func NewBST() *BST {
	return &BST{
		Root: nil,
		Size: 0,
	}
}

func (bst *BST) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		bst.Size++
		return NewTreeNode(val)
	}

	if val < node.Val {
		node.Left = bst.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = bst.insertHelper(node.Right, val)
	}

	// if val == node.Val, skip insert (no duplicates)

	return node
}

func (bst *BST) Insert(val int) {
	bst.Root = bst.insertHelper(bst.Root, val)
}

func (bst *BST) searchHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if val == node.Val {
		return node
	} else if val < node.Val {
		return bst.searchHelper(node.Left, val)
	} else {
		return bst.searchHelper(node.Right, val)
	}
}

func (bst *BST) Search(val int) bool {
	node := bst.searchHelper(bst.Root, val)

	return node != nil
}

func (bst *BST) Delete(val int) {
	bst.Root = bst.deleteHelper(bst.Root, val)
}

func (bst *BST) deleteHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if val < node.Val {
		node.Left = bst.deleteHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = bst.deleteHelper(node.Right, val)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}

		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// Case3: 2 children
		// Find inorder successor (smallest in right subtree)
		minNode := bst.findMin(node.Right)
		// Update current nodes value to that minNode
		node.Val = minNode.Val
		// Delete that node from the right subtree now
		node.Right = bst.deleteHelper(node.Right, minNode.Val)
	}

	return node
}

func (bst *BST) findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (bst *BST) inorderHelper(node *TreeNode, result *[]int) {
	if node != nil {
		bst.inorderHelper(node.Left, result)
		*result = append(*result, node.Val)
		bst.inorderHelper(node.Right, result)
	}
}

func (bst *BST) InOrderTraversal() []int {
	var result []int
	bst.inorderHelper(bst.Root, &result)

	return result
}

func (bst *BST) heightHelper(node *TreeNode) int {
	if node == nil {
		return -1
	}

	leftHeight := bst.heightHelper(node.Left)
	rightHeight := bst.heightHelper(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func (bst *BST) Height() int {
	return bst.heightHelper(bst.Root)
}
