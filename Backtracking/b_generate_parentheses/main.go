package main

import "fmt"

// Problem: Generate all combinations of well-formed parentheses
// Given n pairs of parentheses, write a function to generate all combinations
// Time Complexity: O(4^n / sqrt(n)) - Catalan number
// Space Complexity: O(4^n / sqrt(n)) for storing all valid combinations

// Approach 1: Backtracking with counters
func generateParenthesis(n int) []string {
	var result []string
	backtrack("", 0, 0, n, &result)
	return result
}

func backtrack(current string, open, close, max int, result *[]string) {
	// Base case: if we've used all pairs
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}
	
	// Add opening parenthesis if we haven't used all of them
	if open < max {
		backtrack(current+"(", open+1, close, max, result)
	}
	
	// Add closing parenthesis if it won't make string invalid
	if close < open {
		backtrack(current+")", open, close+1, max, result)
	}
}

// Approach 2: Backtracking with string builder for efficiency
func generateParenthesisOptimized(n int) []string {
	var result []string
	var current []rune
	
	backtrackOptimized(current, 0, 0, n, &result)
	return result
}

func backtrackOptimized(current []rune, open, close, max int, result *[]string) {
	if len(current) == max*2 {
		*result = append(*result, string(current))
		return
	}
	
	if open < max {
		current = append(current, '(')
		backtrackOptimized(current, open+1, close, max, result)
		current = current[:len(current)-1] // backtrack
	}
	
	if close < open {
		current = append(current, ')')
		backtrackOptimized(current, open, close+1, max, result)
		current = current[:len(current)-1] // backtrack
	}
}

// Approach 3: Dynamic Programming (Bottom-up)
func generateParenthesisDP(n int) []string {
	if n == 0 {
		return []string{""}
	}
	
	// dp[i] contains all valid parentheses combinations for i pairs
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// For each way to split i pairs into j and i-1-j
			for _, left := range dp[j] {
				for _, right := range dp[i-1-j] {
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}
		}
	}
	
	return dp[n]
}

// Helper function to validate parentheses (for testing)
func isValidParentheses(s string) bool {
	count := 0
	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

func main() {
	n := 3
	
	fmt.Printf("=== Backtracking Approach (n=%d) ===\n", n)
	result1 := generateParenthesis(n)
	for i, combo := range result1 {
		fmt.Printf("%d: %s (valid: %v)\n", i+1, combo, isValidParentheses(combo))
	}
	
	fmt.Printf("\n=== Optimized Backtracking (n=%d) ===\n", n)
	result2 := generateParenthesisOptimized(n)
	for i, combo := range result2 {
		fmt.Printf("%d: %s\n", i+1, combo)
	}
	
	fmt.Printf("\n=== Dynamic Programming (n=%d) ===\n", n)
	result3 := generateParenthesisDP(n)
	for i, combo := range result3 {
		fmt.Printf("%d: %s\n", i+1, combo)
	}
	
	fmt.Printf("\nTotal combinations: %d\n", len(result1))
	fmt.Printf("Expected (Catalan number C_%d): %d\n", n, catalan(n))
}

// Calculate nth Catalan number
func catalan(n int) int {
	if n <= 1 {
		return 1
	}
	
	result := 0
	for i := 0; i < n; i++ {
		result += catalan(i) * catalan(n-1-i)
	}
	return result
}