package main

/*
Problem: K Closest Numbers to X

Description:
Given an unsorted array and two numbers x and k, find k closest values to x.
Return the k elements from the array that have the smallest absolute difference from x.

Example 1:
Input: arr = [10, 2, 14, 4, 7, 6], x = 5, k = 3
Output: [4, 6, 7]
Explanation: The differences are: |10-5|=5, |2-5|=3, |14-5|=9, |4-5|=1, |7-5|=2, |6-5|=1
The 3 closest numbers are 4, 6, and 7.

Example 2:
Input: arr = [1, 3, 4, 6, 8, 9], x = 5, k = 2
Output: [4, 6]
Explanation: 4 and 6 have the smallest differences from 5.

Example 3:
Input: arr = [5, 6, 7, 8, 9], x = 7, k = 3
Output: [6, 7, 8]
Explanation: Elements with differences: |5-7|=2, |6-7|=1, |7-7|=0, |8-7|=1, |9-7|=2

Constraints:
- 1 <= arr.length <= 10^4
- 1 <= k <= arr.length
- -10^4 <= arr[i], x <= 10^4

Algorithm Options:
1. Max Heap: Use max heap of size k, keep k smallest differences
2. Sort: Sort by absolute difference, take first k elements
3. Quick Select: Partition around kth closest element

Time Complexity: O(n log k) using heap, O(n log n) using sort
Space Complexity: O(k)

Companies: Amazon, Google, Microsoft, Facebook
*/
