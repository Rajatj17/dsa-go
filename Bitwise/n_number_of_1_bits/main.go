package main

/*
Problem: Number of 1 Bits

Description:
Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

Example 1:
Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.

Example 2:
Input: n = 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.

Example 3:
Input: n = 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.

Constraints:
- The input must be a binary string of length 32

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(log n) or O(1) since we're dealing with 32-bit integers
Space Complexity: O(1)

Related Problems:
- Reverse Bits
- Power of Two
- Counting Bits
- Binary Watch

Note: Multiple approaches:
1. Loop and check each bit: n & 1, then n >>= 1
2. Brian Kernighan's algorithm: n & (n-1) removes the rightmost set bit
3. Built-in functions: __builtin_popcount() in C++

Key insight: n & (n-1) always removes the rightmost set bit from n.
*/

func main() {
	// TODO: Implement the solution
}