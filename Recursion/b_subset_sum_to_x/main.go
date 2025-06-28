package subsetsumtox

// let's say we are given a set of non-negative integers and values sum,
// then we are to find if there is a subset that sums to a given , OK, maybe
// we are given this array and we are given this number that is 16 we want to find out, a subset whose sum is sixteen.
// OK, all we need to just tell whether such a subset exists or not, this question can be asked in various forms.
// Maybe it could be like the check whether there is a subset or not.
// Maybe true or false, or it could be count how many such subjects are possible.
// Maybe we also had the number one then 10 and six.
// Would be one subset, we would have like 15 and one this could be another subset.

// Input
// arr = [10,12, 15, 6, 19, 20]
// sum = 6

// Output
// Yes

func subsetSumHenlper(arr []int, arrLength int, index int, sum int) int {
	if index == arrLength {
		if sum == 0 {
			return 1
		}
		return 0
	}

	// inc
	inc := subsetSumHenlper(arr, arrLength, index+1, sum-arr[index])
	// exc
	exc := subsetSumHenlper(arr, arrLength, index, sum)

	return inc + exc
}

func SubsetSum(arr []int, sum int) {
	subsetSumHenlper(arr, len(arr), 0, sum)
}
