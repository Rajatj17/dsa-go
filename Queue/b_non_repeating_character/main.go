package nonrepeatingcharacter

import "container/list"

// First Non-Repeating Character
// Given a stream of characters (lowercase alphabets),
// find the first non-repeating character from stream.
// You need to tell the first non-repeating character in O(1) time at each index.
// If for a current index there is no such character return '0' for that particular index.

// Input Format:
// A String S of length N passed as a parameter to the given function.
// Output Format:
// Return a vector of characters of length N where V[i] character represents first non-repeating character from S[0] to S[i].

// Constraints:
// 1<=N<=100

// Expected time complexity:
// O(N) where N is the total number of input characters in one testcase. (Use the property of queue)
// Sample Testcase :

// Input:

// aabcbcd

// Output:

// a 0 b b c 0 d

func NonRepeatingCharacter(input string) []rune {
	result := []rune{}
	deque := list.New()
	runeMap := map[rune]int{}
	for _, val := range input {
		runeMap[val]++
		if runeMap[val] == 1 {
			deque.PushBack(val)
		}

		// if deque.Len() < 0
		// top := deque.Front()
		// if top
	}

	return result
}
