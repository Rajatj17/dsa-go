package main

/*
Problem: Best Time to Buy and Sell Stock with Cooldown

Description:
You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete as many transactions as you like 
(i.e., buy one and sell one share of the stock multiple times) with the following restrictions:

- After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).

Example 1:
Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]

Example 2:
Input: prices = [1]
Output: 0
Explanation: No transactions can be made, so the maximum profit is 0.

Constraints:
- 1 <= prices.length <= 5000
- 0 <= prices[i] <= 1000

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe

Time Complexity: O(n) where n is the number of days
Space Complexity: O(1) with space optimization

Related Problems:
- Best Time to Buy and Sell Stock
- Best Time to Buy and Sell Stock II
- Best Time to Buy and Sell Stock III
- Best Time to Buy and Sell Stock IV
- Best Time to Buy and Sell Stock with Transaction Fee

Note: This is a state machine DP problem.
We need to track different states:
- hold: currently holding a stock
- sold: just sold a stock (cooldown period)
- rest: not holding stock and not in cooldown

Transitions:
- hold[i] = max(hold[i-1], rest[i-1] - prices[i])
- sold[i] = hold[i-1] + prices[i]
- rest[i] = max(rest[i-1], sold[i-1])
*/

func main() {
	// TODO: Implement the solution
}