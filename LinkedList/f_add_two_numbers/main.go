package main

/*
Problem: Add Two Numbers

Description:
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:
- The number of nodes in each linked list is in the range [1, 100]
- 0 <= Node.val <= 9
- It is guaranteed that the list represents a number that does not have leading zeros

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(max(m, n)) where m and n are lengths of the two lists
Space Complexity: O(max(m, n))

Related Problems:
- Multiply Strings
- Add Binary
- Sum of Two Integers
- Add Strings

Note: Elementary math with carry:
1. Process digits from least significant to most significant
2. Handle carry from previous addition
3. Continue until both lists are exhausted and no carry remains

Key insight: Since digits are in reverse order, we can process from left to right directly.
*/

func main() {
	// TODO: Implement the solution
}