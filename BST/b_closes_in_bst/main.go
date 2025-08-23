package closesinbst

/*
Problem: Find Closest Value in BST

Description:
Given the root of a binary search tree and a target value, return the value in the BST 
that is closest to the target. If there are multiple values with the same minimum difference, 
return any of them.

Example 1:
Input: root = [4,2,5,1,3], target = 3.714286
Output: 4
Explanation: The closest value is 4.

Example 2:
Input: root = [1], target = 4.428571
Output: 1
Explanation: There's only one value in the tree, so return 1.

Example 3:
Input: root = [4,2,5,1,3], target = 2.1
Output: 2
Explanation: 2 is closer to 2.1 than 1 or 3.

Constraints:
- The number of nodes in the tree is in the range [1, 10^4]
- 0 <= Node.val <= 10^9
- -10^9 <= target <= 10^9

Algorithm:
Use BST property to efficiently navigate:
- If target < current.val, go left (but current might still be closest)
- If target > current.val, go right (but current might still be closest)
- Keep track of the closest value seen so far

Time Complexity: O(log n) average, O(n) worst case
Space Complexity: O(1) iterative, O(log n) recursive

Companies: Google, Facebook, Amazon, Microsoft
*/

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