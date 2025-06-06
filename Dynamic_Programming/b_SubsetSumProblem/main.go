package main

/*
*

	Question:
	 	Arr[] = [1,2, 3, 4, 6]
		Check whether the subset with sum = 4 exists in the arry

*
*/
func SubsetSum(arr []int, sum int) {
	n := len(arr) + 1
	w := sum + 1

	t := [][]bool{{}}

	for i := 0; i < n; i++ {
		for j := 0; j < w; j++ {

			if arr[i-1] <= j {
				//      = exclude  || include
				t[i][j] = t[i-1][j] || t[i-1][j-arr[i-1]]
			} else {

				t[i][j] = t[i-1][j]
			}
		}
	}
}
