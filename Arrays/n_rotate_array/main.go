package main

/*
Problem: Rotate Array

Description:
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Constraints:
- 1 <= nums.length <= 10^5
- -2^31 <= nums[i] <= 2^31 - 1
- 0 <= k <= 10^5

Follow up:
- Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
- Could you do it in-place with O(1) extra space?

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(1) for reverse approach

Related Problems:
- Rotate List
- Reverse Words in a String
- Reverse String

Note: Multiple approaches:
1. Extra array: O(n) space
2. Cyclic replacements: O(1) space, complex
3. Reverse approach: O(1) space, elegant
   - Reverse entire array
   - Reverse first k elements
   - Reverse remaining elements

Key insight: k = k % n to handle cases where k > n
*/

func main() {
	// TODO: Implement the solution
}