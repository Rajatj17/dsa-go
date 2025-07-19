package main

/*
Problem: Climbing Stairs

Description:
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Example 3:
Input: n = 4
Output: 5
Explanation: There are five ways to climb to the top.
1. 1+1+1+1
2. 1+1+2
3. 1+2+1
4. 2+1+1
5. 2+2

Constraints:
- 1 <= n <= 45

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber, LinkedIn

Time Complexity: O(n)
Space Complexity: O(1) with space optimization

Related Problems:
- Climbing Stairs with Variable Steps
- Min Cost Climbing Stairs
- House Robber
- Fibonacci Number

Note: This is essentially a Fibonacci sequence problem in disguise.
dp[i] = dp[i-1] + dp[i-2] represents the number of ways to reach step i.
*/

func main() {
	// TODO: Implement the solution
}