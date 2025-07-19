package main

/*
Problem: Valid Parentheses

Description:
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
- 1 <= s.length <= 10^4
- s consists of parentheses only '()[]{}'

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(n)

Related Problems:
- Generate Parentheses
- Longest Valid Parentheses
- Remove Invalid Parentheses
- Check if a Parentheses String Can Be Valid

Note: Classic stack problem:
1. Push opening brackets onto stack
2. For closing brackets, check if top of stack matches
3. At the end, stack should be empty

Key insight: Use a stack to keep track of opening brackets and match them with closing brackets.
*/

func main() {
	// TODO: Implement the solution
}