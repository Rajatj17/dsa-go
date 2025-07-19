package main

/*
Problem: Longest Consecutive Sequence

Description:
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Constraints:
- 0 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9

LeetCode Link: https://leetcode.com/problems/longest-consecutive-sequence/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(n)

Related Problems:
- Binary Tree Longest Consecutive Sequence
- Longest Consecutive Sequence II

Note: Hash set approach:
1. Put all numbers in a hash set for O(1) lookup
2. For each number, check if it's the start of a sequence (num-1 not in set)
3. If it's the start, count consecutive numbers
4. Keep track of maximum length found

Key insight: Only start counting from numbers that are the beginning of a sequence to avoid redundant work.
*/

func main() {
	// TODO: Implement the solution
}