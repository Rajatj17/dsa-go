package main

/*
Problem: Remove Nth Node From End of List

Description:
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]

Constraints:
- The number of nodes in the list is sz
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz

Follow up: Could you do this in one pass?

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(1)

Related Problems:
- Swapping Nodes in a Linked List
- Delete Node in a Linked List
- Remove Duplicates from Sorted List

Note: Two-pointer technique:
1. Use two pointers with n+1 gap between them
2. When the fast pointer reaches the end, slow pointer is at the node before the target
3. Remove the next node from slow pointer

Key insight: Maintain a gap of n+1 nodes between two pointers to find the node before the target.
*/

func main() {
	// TODO: Implement the solution
}