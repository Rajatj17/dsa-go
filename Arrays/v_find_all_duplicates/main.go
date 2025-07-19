package main

/*
Problem: Find All Duplicates in an Array

Description:
Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.

You must write an algorithm that runs in O(n) time and uses only constant extra space.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]

Example 2:
Input: nums = [1,1,2]
Output: [1]

Example 3:
Input: nums = [1]
Output: []

Constraints:
- n == nums.length
- 1 <= n <= 10^5
- 1 <= nums[i] <= n
- Each element in nums appears once or twice

LeetCode Link: https://leetcode.com/problems/find-all-duplicates-in-an-array/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(1) not counting output array

Related Problems:
- Find All Numbers Disappeared in an Array
- Find the Duplicate Number
- First Missing Positive

Note: Index marking approach:
1. For each number, mark its corresponding index as visited by negating
2. If index is already negative, the number is duplicate
3. Use abs(nums[i]) to get original value

Key insight: Use the constraint that numbers are in range [1,n] to use array indices for marking.
*/

func main() {
	// TODO: Implement the solution
}