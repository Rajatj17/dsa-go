package subsetsofagivenstring

func generateSubsetsHelper(s string, index int, current string, result *[]string) {
	if index == len(s) {
		*result = append(*result, current)
		return
	}

	// Include the string
	generateSubsetsHelper(s, index+1, current+string(s[index]), result)

	// Exclude the string
	generateSubsetsHelper(s, index+1, current, result)

}

func GenerateSubsets(s string) []string {
	var result []string

	generateSubsetsHelper(s, 0, "", &result)

	return result
}
