package main

import "fmt"

// Problem: Fixed Size Sliding Window Template
// This is a template for solving fixed-size sliding window problems
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1) for basic problems, O(k) if we need to store window elements

// Example 1: Maximum sum of k consecutive elements
func maxSumSubarray(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}
	
	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	
	maxSum := windowSum
	
	// Slide the window
	for i := k; i < len(arr); i++ {
		// Remove the leftmost element and add the new element
		windowSum = windowSum - arr[i-k] + arr[i]
		maxSum = max(maxSum, windowSum)
	}
	
	return maxSum
}

// Example 2: Average of all subarrays of size k
func averageOfSubarrays(arr []int, k int) []float64 {
	if len(arr) < k {
		return []float64{}
	}
	
	var result []float64
	
	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	
	result = append(result, float64(windowSum)/float64(k))
	
	// Slide the window
	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i]
		result = append(result, float64(windowSum)/float64(k))
	}
	
	return result
}

// Example 3: Find maximum element in each window of size k
func maxElementInWindow(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}
	
	var result []int
	
	// For each window, find the maximum
	for i := 0; i <= len(arr)-k; i++ {
		maxVal := arr[i]
		for j := i; j < i+k; j++ {
			maxVal = max(maxVal, arr[j])
		}
		result = append(result, maxVal)
	}
	
	return result
}

// Example 4: Check if any subarray of size k has duplicate elements
func hasDuplicateInWindow(arr []int, k int) bool {
	if len(arr) < k {
		return false
	}
	
	for i := 0; i <= len(arr)-k; i++ {
		window := make(map[int]bool)
		for j := i; j < i+k; j++ {
			if window[arr[j]] {
				return true
			}
			window[arr[j]] = true
		}
	}
	
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("=== Fixed Size Sliding Window Examples ===")
	
	arr := []int{1, 4, 2, 9, 5, 10, 7}
	k := 3
	
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Window size: %d\n", k)
	
	// Example 1: Maximum sum
	fmt.Printf("\n1. Maximum sum of %d consecutive elements: %d\n", k, maxSumSubarray(arr, k))
	
	// Example 2: Averages
	fmt.Printf("\n2. Average of all subarrays of size %d: %v\n", k, averageOfSubarrays(arr, k))
	
	// Example 3: Maximum in each window
	fmt.Printf("\n3. Maximum element in each window: %v\n", maxElementInWindow(arr, k))
	
	// Example 4: Duplicate check
	fmt.Printf("\n4. Has duplicate in any window: %v\n", hasDuplicateInWindow(arr, k))
	
	// Test with different array
	fmt.Println("\n=== Test with array containing duplicates ===")
	arr2 := []int{1, 2, 3, 2, 4, 5}
	fmt.Printf("Array: %v\n", arr2)
	fmt.Printf("Has duplicate in window of size 3: %v\n", hasDuplicateInWindow(arr2, 3))
}