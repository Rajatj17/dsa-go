package main

// You are playing a game with your k friends. You have an array of N coins, at each index i you have a value of a[i]
// Your task is to divide the coins among the group of k friends by doing consecutive segments of the array.
// Each friend will get a total sum of the coins in that subarray. Since all your friends are greedy and they will pick
// the largest k - 1 segments and you will get the smallest segment.
// Find out the maximum value you can make by making an optimal partition

// Note: The coin array may not be sorted

// Input
// K = 3
// coins = {1, 2, 3, 4}
// Output = 3

// Explanation: {1 + 2}= 3, {3}, {4}
