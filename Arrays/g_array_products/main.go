package arrayproducts

// Array Products

// Implement a function that takes in a vector of integers,
// and returns a vector of the same length,
// where each element in the output array is equal to the product of every other number in the input array.
// Solve this problem without using division.

// In other words, the value at output[i] is equal to the
// product of every number in the input array other than input[i].
// You can assume that answer can be stored inside int datatype and no-overflow will occur due to products.

// Sample Input

// Both inputs and outputs are vectors.

// {1,2,3,4,5}

// Sample Output

// {120, 60, 40, 30, 24}

// ArrayProducts calculates product of all elements except self without division
func ArrayProducts(nums []int) []int {
	// 		1,   2,    6,  24, 1
	//     1,  120, ,60, 20, 5
	// 120 * 1, 60 * 1,

	n := len(nums)
	result := make([]int, n)

	// Step 1: Calculate left products
	// result[i] will contain product of all elements to the left of i
	result[0] = 1 // No elements to the left of index 0
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Step 2: Calculate right products and multiply with left products
	// Use a variable to track running product from right
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct *= nums[i]
	}

	return result
}

// ArrayProductsVerbose - Alternative approach using separate left and right slices
func ArrayProductsVerbose(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	result := make([]int, n)

	// Calculate left products
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// Calculate right products
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	// Multiply left and right products
	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}
