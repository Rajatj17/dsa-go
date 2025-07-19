package main

/*
Problem: Missing Number

Description:
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Example 1:
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

Example 2:
Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.

Example 3:
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

Constraints:
- n == nums.length
- 1 <= n <= 10^4
- 0 <= nums[i] <= n
- All the numbers of nums are unique

LeetCode Link: https://leetcode.com/problems/missing-number/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(1)

Related Problems:
- Find the Duplicate Number
- Find All Numbers Disappeared in an Array
- First Missing Positive

Note: Multiple approaches:
1. Aggregate and subtract: Sum of [0,n] - Sum of array
2. XOR: XOR all numbers from 0 to n with all array elements
3. Cyclic sort: Place each number at its correct index
4. Hash set: Mark seen numbers and find missing

Key insight: Expected sum = n*(n+1)/2, actual sum = sum of array, missing = expected - actual
*/

func main() {
	// TODO: Implement the solution
}