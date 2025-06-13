package minswaps

import "slices"

// Given an array arr[] of distinct elements, find the minimum number of swaps required to sort the array.

// Examples:

// Input: arr[] = [2, 8, 5, 4]
// Output: 1
// Explanation: Swap 8 with 4 to get the sorted array.

// Input: arr[] = [10, 19, 6, 3, 5]
// Output: 2
// Explanation: Swap 10 with 3 and 19 with 5 to get the sorted array.

// Input: arr[] = [1, 3, 4, 5, 6]
// Output: 0
// Explanation: Input array is already sorted.

func MinCountSwaps(arr []int) int {
	temp := []int{}
	copy(temp, arr)

	slices.Sort(temp)

	lookupSet := make(map[int]int)
	for key, value := range temp {
		lookupSet[value] = key
	}

	visited := []bool{}

	swaps := 0
	for key, value := range arr {
		if visited[key] || lookupSet[value] == key {
			continue
		}

		cycle := 0
		position := key
		for {
			if visited[position] {
				break
			}

			visited[position] = true
			position = lookupSet[arr[position]]
			cycle++
		}

		if cycle > 0 {
			swaps += (cycle - 1)
		}
	}

	return swaps
}
