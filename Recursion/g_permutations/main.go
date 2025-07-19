package main

/*
Problem: Permutations

Description:
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]

Constraints:
- 1 <= nums.length <= 6
- -10 <= nums[i] <= 10
- All the integers of nums are unique

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n! * n) where n is the length of nums
Space Complexity: O(n) for recursion stack

Related Problems:
- Permutations II
- Next Permutation
- Permutation Sequence
- Letter Combinations of a Phone Number

Note: Backtracking approach:
1. For each position, try all unused numbers
2. Add current number to path and recurse
3. Remove current number (backtrack) and try next
4. Base case: when path length equals nums length

Key insight: Use backtracking to generate all permutations by swapping elements.
*/

func main() {
	// TODO: Implement the solution
}