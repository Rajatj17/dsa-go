package main

import "fmt"

// Problem: Generate all possible subsets (power set) of a given array
// Time Complexity: O(2^n * n) where n is the number of elements
// Space Complexity: O(2^n * n) for storing all subsets

// Approach 1: Backtracking (DFS)
func subsets(nums []int) [][]int {
	var result [][]int
	var current []int
	
	backtrack(nums, 0, current, &result)
	return result
}

func backtrack(nums []int, start int, current []int, result *[][]int) {
	// Add current subset to result (make a copy)
	subset := make([]int, len(current))
	copy(subset, current)
	*result = append(*result, subset)
	
	// Try adding each remaining element
	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])     // Include nums[i]
		backtrack(nums, i+1, current, result) // Recurse
		current = current[:len(current)-1]     // Backtrack (exclude nums[i])
	}
}

// Approach 2: Iterative (Build subsets progressively)
func subsetsIterative(nums []int) [][]int {
	result := [][]int{{}} // Start with empty subset
	
	for _, num := range nums {
		var newSubsets [][]int
		// For each existing subset, create a new subset by adding current number
		for _, subset := range result {
			newSubset := make([]int, len(subset)+1)
			copy(newSubset, subset)
			newSubset[len(subset)] = num
			newSubsets = append(newSubsets, newSubset)
		}
		result = append(result, newSubsets...)
	}
	
	return result
}

// Approach 3: Bit Manipulation
func subsetsBitManipulation(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0, 1<<n) // 2^n subsets
	
	// Generate all possible combinations using bit masks
	for mask := 0; mask < (1 << n); mask++ {
		var subset []int
		for i := 0; i < n; i++ {
			// Check if i-th bit is set in mask
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}
	
	return result
}

func main() {
	nums := []int{1, 2, 3}
	
	fmt.Println("=== Backtracking Approach ===")
	result1 := subsets(nums)
	for i, subset := range result1 {
		fmt.Printf("Subset %d: %v\n", i, subset)
	}
	
	fmt.Println("\n=== Iterative Approach ===")
	result2 := subsetsIterative(nums)
	for i, subset := range result2 {
		fmt.Printf("Subset %d: %v\n", i, subset)
	}
	
	fmt.Println("\n=== Bit Manipulation Approach ===")
	result3 := subsetsBitManipulation(nums)
	for i, subset := range result3 {
		fmt.Printf("Subset %d: %v\n", i, subset)
	}
	
	fmt.Printf("\nTotal subsets: %d (Expected: 2^%d = %d)\n", 
		len(result1), len(nums), 1<<len(nums))
}