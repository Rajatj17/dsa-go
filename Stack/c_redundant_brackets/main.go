package redundantbrackets

import basic "DSA/Stack/a_basic"

func CheckRedundants(input string) bool {
	stack := basic.NewStack[rune]()

	for _, ch := range input {
		if ch != ')' {
			stack.Push(ch)
		} else {
			operatorFound := false

			val, exists := stack.Top()
			for exists && val != '(' {
				if val == '+' || val == '-' || val == '*' || val == '/' {
					operatorFound = true
				}

				val, exists = stack.Pop()
			}
			stack.Pop() // Pop the opening bracket

			if !operatorFound {
				return true
			}
		}
	}

	return false
}
