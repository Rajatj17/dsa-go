package balancedparanthesis

import basic "DSA/Stack/a_basic"

func IsBalanced(input string) bool {
	stack := basic.NewStack[rune]()

	for _, char := range input {
		switch char {
		case '(':
		case '{':
		case '[':
			stack.Push(char)
		case '}':
			if val, exists := stack.Top(); exists && val == '}' {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	return false
}
