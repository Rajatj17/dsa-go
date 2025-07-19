package main

import "fmt"

// Problem: Partition string into palindromic substrings
// Given a string s, partition s such that every substring is a palindrome
// Return all possible palindrome partitioning of s
// Time Complexity: O(N * 2^N) where N is length of string
// Space Complexity: O(N) for recursion stack + O(N^2) for memoization

// Approach 1: Backtracking with basic palindrome check
func partition(s string) [][]string {
	var result [][]string
	var path []string
	
	backtrack(s, 0, path, &result)
	return result
}

func backtrack(s string, start int, path []string, result *[][]string) {
	if start == len(s) {
		// Make a copy of the path and add to result
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}
	
	for end := start; end < len(s); end++ {
		substring := s[start:end+1]
		if isPalindrome(substring) {
			path = append(path, substring)
			backtrack(s, end+1, path, result)
			path = path[:len(path)-1] // backtrack
		}
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Approach 2: Optimized with memoization for palindrome checks
func partitionOptimized(s string) [][]string {
	n := len(s)
	// Create memoization table for palindrome checks
	memo := make([][]bool, n)
	for i := range memo {
		memo[i] = make([]bool, n)
	}
	
	// Pre-compute palindrome checks
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			memo[i][j] = isPalindromeRange(s, i, j)
		}
	}
	
	var result [][]string
	var path []string
	
	backtrackOptimized(s, 0, path, &result, memo)
	return result
}

func backtrackOptimized(s string, start int, path []string, result *[][]string, memo [][]bool) {
	if start == len(s) {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}
	
	for end := start; end < len(s); end++ {
		if memo[start][end] {
			path = append(path, s[start:end+1])
			backtrackOptimized(s, end+1, path, result, memo)
			path = path[:len(path)-1]
		}
	}
}

func isPalindromeRange(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Approach 3: Dynamic Programming to build palindrome table
func partitionDP(s string) [][]string {
	n := len(s)
	
	// Build palindrome table using DP
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	
	// Every single character is a palindrome
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	
	// Check for palindromes of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	
	// Check for palindromes of length 3 and more
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}
	
	var result [][]string
	var path []string
	
	backtrackDP(s, 0, path, &result, dp)
	return result
}

func backtrackDP(s string, start int, path []string, result *[][]string, dp [][]bool) {
	if start == len(s) {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}
	
	for end := start; end < len(s); end++ {
		if dp[start][end] {
			path = append(path, s[start:end+1])
			backtrackDP(s, end+1, path, result, dp)
			path = path[:len(path)-1]
		}
	}
}

// Helper function to count total palindromic substrings
func countPalindromicSubstrings(s string) int {
	n := len(s)
	count := 0
	
	// Check all possible substrings
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isPalindromeRange(s, i, j) {
				count++
			}
		}
	}
	
	return count
}

// Helper function to find the minimum cuts needed
func minCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	
	// Build palindrome table
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}
	
	// Find minimum cuts
	cuts := make([]int, n)
	for i := 0; i < n; i++ {
		cuts[i] = i // Maximum cuts needed
		if dp[0][i] {
			cuts[i] = 0
		} else {
			for j := 0; j < i; j++ {
				if dp[j+1][i] {
					cuts[i] = min(cuts[i], cuts[j]+1)
				}
			}
		}
	}
	
	return cuts[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := []string{"aab", "racecar", "abcba", "aba"}
	
	for _, s := range testCases {
		fmt.Printf("=== String: '%s' ===\n", s)
		
		fmt.Println("Backtracking Approach:")
		result1 := partition(s)
		for i, partition := range result1 {
			fmt.Printf("  %d: %v\n", i+1, partition)
		}
		
		fmt.Println("Optimized Approach:")
		result2 := partitionOptimized(s)
		for i, partition := range result2 {
			fmt.Printf("  %d: %v\n", i+1, partition)
		}
		
		fmt.Println("DP Approach:")
		result3 := partitionDP(s)
		for i, partition := range result3 {
			fmt.Printf("  %d: %v\n", i+1, partition)
		}
		
		fmt.Printf("Total partitions: %d\n", len(result1))
		fmt.Printf("Total palindromic substrings: %d\n", countPalindromicSubstrings(s))
		fmt.Printf("Minimum cuts needed: %d\n", minCut(s))
		fmt.Println()
	}
}