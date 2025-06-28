package main

func knapsackProblem(price []int, weight []int, n int, maxWeight int) int {
	t := [][]int{{}}

	for i := range n + 1 {
		for j := range maxWeight + 1 {

			if weight[i-1] <= j {

				t[i][j] = max(
					price[i-1]+t[i-1][j-weight[i-1]],
					t[i-1][j],
				)
			} else {

				t[i][j] = t[i-1][j]
			}

		}
	}

	return t[n][maxWeight]
}

func main() {
	var price []int
	var weight []int
	var maxWeight int

	knapsackProblem(price, weight, len(price), maxWeight)
}
