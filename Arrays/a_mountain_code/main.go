package mountaincode

//  Longest Mountain in Array

// You may recall that an array arr is a mountain array if and only if:

// arr.length >= 3
// There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
// Given an integer array arr, return the length of the longest subarray,
// which is a mountain. Return 0 if there is no mountain subarray.

// Example 1:

// Input: arr = [2,1,4,7,3,2,5]
// Output: 5
// Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
// Example 2:

// Input: arr = [2,2,2]
// Output: 0
// Explanation: There is no mountain.

// Constraints:

// 1 <= arr.length <= 104
// 0 <= arr[i] <= 104

// Follow up:

// Can you solve it using only one pass?
// Can you solve it in O(1) space?

func HighestMountain(arr []int) int {
	n := len(arr)

	largest := 0
	peakStart := 0
	peak := false
	for i := 1; i <= n-2; i++ {
		// Find peak
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			peak = true
		}

		// Find peak end
		if peak == true && arr[i] < arr[i-1] && arr[i] < arr[i+1] {
			maxPeak := i - peakStart
			if maxPeak > largest {
				largest = maxPeak
			}
			// Mark peak as false
			peak = false
			// Mark last peak distance as this one
			peakStart = i
		}
	}

	return largest
}

// Alternative cleaner approach
func HighestMountainV2(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	maxLen := 0

	for i := 1; i < n-1; i++ {
		// Check if current element is a peak
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			// Expand left to find start of mountain
			left := i
			for left > 0 && arr[left-1] < arr[left] {
				left--
			}

			// Expand right to find end of mountain
			right := i
			for right < n-1 && arr[right] > arr[right+1] {
				right++
			}

			// Calculate mountain length
			length := right - left + 1
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}
