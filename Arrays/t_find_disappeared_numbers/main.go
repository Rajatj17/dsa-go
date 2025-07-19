package main

/*
Problem: Find All Numbers Disappeared in an Array

Description:
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Example 2:
Input: nums = [1,1]
Output: [2]

Constraints:
- n == nums.length
- 1 <= n <= 10^5
- 1 <= nums[i] <= n

Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

LeetCode Link: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(1) not counting output array

Related Problems:
- Find All Duplicates in an Array
- Missing Number
- First Missing Positive

Note: Multiple approaches:
1. Hash set: Mark seen numbers and find missing
2. Index marking: Mark visited indices by negating values
3. Cyclic sort: Place each number at correct position

Key insight: Use array indices to mark presence - negate value at index (num-1) to mark num as present.
*/

func main() {
	// TODO: Implement the solution
}