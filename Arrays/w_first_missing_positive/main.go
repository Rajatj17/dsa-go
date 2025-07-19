package main

/*
Problem: First Missing Positive

Description:
Given an unsorted integer array nums, return the smallest missing positive integer.

You must implement an algorithm that runs in O(n) time and uses constant extra space.

Example 1:
Input: nums = [1,2,0]
Output: 3

Example 2:
Input: nums = [3,4,-1,1]
Output: 2

Example 3:
Input: nums = [7,8,9,11,12]
Output: 1

Constraints:
- 1 <= nums.length <= 10^5
- -2^31 <= nums[i] <= 2^31 - 1

LeetCode Link: https://leetcode.com/problems/first-missing-positive/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(1)

Related Problems:
- Missing Number
- Find All Numbers Disappeared in an Array
- Kth Missing Positive Number

Note: Cyclic sort approach:
1. Place each positive number at its correct position (num at index num-1)
2. After sorting, first index without correct number gives the answer
3. Handle edge cases: numbers <= 0, numbers > n

Key insight: The answer must be in range [1, n+1] where n is array length.
*/

func main() {
	// TODO: Implement the solution
}