package main

func EqualSumPartition(arr []int) bool {
	n := len(arr) + 1
	sum := 0
	for i := 0; i <= n; i++ {
		sum += arr[i]
	}

	if sum%2 != 0 {
		return false
	}

	// Call SubsetSum(arr, sum/2)

	return true
}
