package main

/*
Problem: Palindrome Number

Description:
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:
- -2^31 <= x <= 2^31 - 1

Follow up: Could you solve it without converting the integer to a string?

LeetCode Link: https://leetcode.com/problems/palindrome-number/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(log x)
Space Complexity: O(1)

Related Problems:
- Reverse Integer
- Valid Palindrome
- Palindrome Linked List

Note: Mathematical approach without string conversion:
1. Negative numbers are not palindromes
2. Numbers ending in 0 (except 0 itself) are not palindromes
3. Reverse half the number and compare with remaining half
4. Handle odd-digit numbers by dividing middle digit

Key insight: Only reverse half the digits to avoid overflow and improve efficiency.
*/

func main() {
	// TODO: Implement the solution
}