package longestband

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
