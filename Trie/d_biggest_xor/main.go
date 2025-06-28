package biggestxor

// Maximum XOR Pair
// Given an vector consisting of N integers, the task is to find the maximum Bitwise XOR from all the possible pairs in the given array in O(N) time.

// Sample Input

// {25, 10, 2, 8, 5, 3}

// Sample Output

// 28

// Create a trie with left and right node

// left - 0
// right - 1

type Node struct {
	left  *Node
	right *Node
}

func MaxXorHelper(value int) {
	currentAns := 0
	var temp *Node
	for i := 31; i >= 1; i-- {
		bit := (value >> i) & 1

		if bit == 0 {
			if temp.right != nil {
				temp = temp.right
				currentAns += (1 << i)
			}
		}
	}
}
