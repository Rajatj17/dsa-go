package main

// ALLOCATE MINIMUM NUMBER OF PAGES:

// Given number of pages in n different books and m students.
// The books are arranged in ascending order of number of pages. Every student is assigned to read some consecutive books.
// The task is to assign books in such a way that the maximum number of pages assigned to a student is minimum.

// Example :

// Input : pages[] = {12, 34, 67, 90}
//         m = 2
// Output : 113
// Explanation:
// There are 2 number of students. Books can be distributed
// in following fashion :
//   1) [12] and [34, 67, 90]
//       Max number of pages is allocated to student
//       2 with 34 + 67 + 90 = 191 pages
//   2) [12, 34] and [67, 90]
//       Max number of pages is allocated to student
//       2 with 67 + 90 = 157 pages
//   3) [12, 34, 67] and [90]
//       Max number of pages is allocated to student
//       1 with 12 + 34 + 67 = 113 pages

// Of the 3 cases, Option 3 has the minimum pages = 113.

func FindMaxAndSum(pages []int) (max int, sum int) {
	max = 0
	sum = 0

	for _, v := range pages {
		if v > max {
			max = v
		}

		sum += v
	}

	return
}

func IsValid(pages []int, k int, mid int) bool {
	entity := 1
	sum := 0
	for _, v := range pages {
		sum += v

		if sum > mid {
			entity++
			sum = v
		}

		if entity < k {
			return false
		}
	}

	return true
}

func AllocatePages(pages []int, numberOfEntities int) int {
	start := 0                         // Or we can take maximum or minium dependending on question.
	start, end := FindMaxAndSum(pages) // Somone has to get maximum pages for sure, so we start with the maximum
	res := -1
	for {
		if start < end {
			break
		}

		mid := (start + end) / 2

		if IsValid(pages, numberOfEntities, mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return res
}
