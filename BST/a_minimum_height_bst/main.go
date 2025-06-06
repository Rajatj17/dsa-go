package minimumheightbst

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

func SortedArrayToBST(arr []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2

	root := NewTreeNode(arr[mid])

	root.Left = SortedArrayToBST(arr, start, mid-1)
	root.Right = SortedArrayToBST(arr, mid+1, end)

	return root
}
