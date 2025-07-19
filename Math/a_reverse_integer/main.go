package main

/*
Problem: Reverse Integer

Description:
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21

Constraints:
- -2^31 <= x <= 2^31 - 1

LeetCode Link: https://leetcode.com/problems/reverse-integer/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(log x)
Space Complexity: O(1)

Related Problems:
- Palindrome Number
- String to Integer (atoi)
- Reverse String

Note: Mathematical approach:
1. Handle negative numbers by tracking sign
2. Reverse digits using modulo and division
3. Check for overflow before updating result
4. Return 0 if overflow occurs

Key insight: Check for overflow before multiplying by 10 and adding next digit.
*/

func main() {
	// TODO: Implement the solution
}