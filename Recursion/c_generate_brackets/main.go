package generatebrackets

import "fmt"

func generateBracketsHelper(bracketString string, openBrackets int, closeBrakcet int, n int) {
	if len(bracketString) == n*2 {
		fmt.Println(bracketString)
		return
	}

	if openBrackets < n {
		generateBracketsHelper(bracketString+"(", openBrackets+1, closeBrakcet, n)
	}

	if closeBrakcet < openBrackets {
		generateBracketsHelper(bracketString+")", openBrackets, closeBrakcet+1, n)
	}

}
func GenerateBrackers(n int) {

}

// // Generate brackers Rules
// ((()))
// ()()()
// (())()
