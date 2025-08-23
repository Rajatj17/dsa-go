package longestband

/*
Problem: Longest Consecutive Sequence (Longest Band)

Description:
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
The algorithm must run in O(n) time complexity.

Example 1:
Input: nums = [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
Output: 9
Explanation: The longest consecutive sequence is [0, 1, 2, 3, 4, 5, 6, 7, 8] with length 9.

Example 3:
Input: nums = [9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6]
Output: 7
Explanation: The longest consecutive sequence is [-1, 0, 1, 3, 4, 5, 6, 7] with length 7.

Constraints:
- 0 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9

Algorithm:
1. Store all elements in a hash set for O(1) lookup
2. For each number, check if it's the start of a sequence (no number-1 exists)
3. If it's the start, count consecutive numbers from that point
4. Keep track of the maximum length found

Time Complexity: O(n)
Space Complexity: O(n)

Companies: Google, Facebook, Amazon, Microsoft, Bloomberg
*/

func LargestBand(arr []int) {
	lookupSet := make(map[int]struct{})

	for _, v := range arr {
		lookupSet[v] = struct{}{}
	}

	largestBand := 0
	for _, element := range arr {
		parent := element - 1
		if _, exists := lookupSet[parent]; exists {
			continue
		}

		bandLength := 0
		child := element + 1

		for {
			if _, exists := lookupSet[child]; exists {
				bandLength++
			} else {
				break
			}
		}

		if bandLength > largestBand {
			largestBand = bandLength
		}
	}
}
