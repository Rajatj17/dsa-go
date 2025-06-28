package closesinbst

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func FindClosestInBST(node *TreeNode, target int) {
	if node == nil {
		return 
	}
}