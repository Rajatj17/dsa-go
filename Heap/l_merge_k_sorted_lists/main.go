package main

/*
Problem: Merge k Sorted Lists

Description:
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

Example 1:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Example 2:
Input: lists = []
Output: []

Example 3:
Input: lists = [[]]
Output: []

Constraints:
- k == lists.length
- 0 <= k <= 10^4
- 0 <= lists[i].length <= 500
- -10^4 <= lists[i][j] <= 10^4
- lists[i] is sorted in ascending order
- The sum of lists[i].length will not exceed 10^4

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(N log k) where N is total number of nodes and k is number of lists
Space Complexity: O(k) for the heap

Related Problems:
- Merge Two Sorted Lists
- Ugly Number II
- Smallest Range Covering Elements from K Lists

Note: Multiple approaches:
1. Brute force: Collect all values, sort, and create new list - O(N log N)
2. Compare one by one: Compare k lists repeatedly - O(kN)
3. Divide and conquer: Merge lists pairwise - O(N log k)
4. Priority queue/heap: Use min heap to always get the smallest element - O(N log k)

Key insight: Use a min heap to efficiently find the smallest element among k lists.
*/

func main() {
	// TODO: Implement the solution
}