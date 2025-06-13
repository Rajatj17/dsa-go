package busylife

import "sort"

// Busy Life
// You are actually very busy person. You have a long list of activities. Your aim is to finish much as activities as possible.

// In the given figure, if you go to date with crush, you cannot participate in the coding contest and you can’t watch the movie.

// Also if you play DotA, you can’t study for the exam.

// If you study for the exam you can’t sleep peacefully.

// The maximum number of activities that you can do for this schedule is 3.

// Either you can

// watch movie, play DotA and sleep peacefully (or)
// date with crush, play DotA and sleep peacefully
// Given the start and finish times of activities,
// print the maximum number of activities.
// Input is already store in a vector of pairs.
// Each pair contains the start and end of the activity.

// Sample Input

// (7,9) (0,10) (4,5) (8,9) (4,10) (5,7)

// Output
// 3

// Activity represents an activity with start and end times
type Activity struct {
	Start int
	End   int
	Index int // Original index for tracking
}

// MaxActivities finds the maximum number of non-overlapping activities
func MaxActivities(activities []Activity) int {
	if len(activities) == 0 {
		return 0
	}

	// Sort activities by their end times (greedy choice)
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].End < activities[j].End
	})

	count := 1 // Always select the first activity (earliest end time)
	lastEndTime := activities[0].End

	// Greedily select activities that don't overlap
	for i := 1; i < len(activities); i++ {
		// If current activity starts after the last selected activity ends
		if activities[i].Start >= lastEndTime {
			count++
			lastEndTime = activities[i].End
		}
	}

	return count
}
