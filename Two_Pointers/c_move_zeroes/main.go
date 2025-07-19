package main

/*
Problem: Move Zeroes

Description:
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]

Constraints:
- 1 <= nums.length <= 10^4
- -2^31 <= nums[i] <= 2^31 - 1

Follow up: Could you minimize the total number of operations done?

LeetCode Link: https://leetcode.com/problems/move-zeroes/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(1)

Related Problems:
- Remove Element
- Remove Duplicates from Sorted Array
- Apply Operations to an Array

Note: Two-pointer approach:
1. Use two pointers: one for non-zero elements, one for iteration
2. When non-zero element is found, move it to the non-zero position
3. Fill remaining positions with zeros
4. Maintain relative order of non-zero elements

Key insight: Use two pointers to segregate zero and non-zero elements while maintaining order.
*/

func main() {
	// TODO: Implement the solution
}