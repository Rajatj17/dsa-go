package runningmedian

import (
	basic_heap "DSA/Heap/a__Basic_Code"
	"container/heap"
)

// MedianFinder implements running median using two heaps
type MedianFinder struct {
	maxHeap *basic_heap.IntMaxHeap // Stores smaller half (left side)
	minHeap *basic_heap.IntMinHeap // Stores larger half (right side)
}

// NewMedianFinder creates a new MedianFinder
func NewMedianFinder() *MedianFinder {
	maxHeap := &basic_heap.IntMaxHeap{}
	minHeap := &basic_heap.IntMinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	return &MedianFinder{
		maxHeap: maxHeap,
		minHeap: minHeap,
	}
}

// balance ensures heap sizes differ by at most 1
func (mf *MedianFinder) balance() {
	maxSize := mf.maxHeap.Len()
	minSize := mf.minHeap.Len()

	// If maxHeap has more than 1 extra element
	if maxSize > minSize+1 {
		val := heap.Pop(mf.maxHeap).(int)
		heap.Push(mf.minHeap, val)
	} else if minSize > maxSize+1 {
		// If minHeap has more than 1 extra element
		val := heap.Pop(mf.minHeap).(int)
		heap.Push(mf.maxHeap, val)
	}
}

func (mf *MedianFinder) AddNum(num int) {
	// Add to appropriate heap
	if mf.maxHeap.Len() == 0 || num <= mf.maxHeap.Top() {
		heap.Push(mf.maxHeap, num)
	} else {
		heap.Push(mf.minHeap, num)
	}

	// Balance heaps
	mf.balance()
}

// FindMedian returns the current median
// Time: O(1), Space: O(1)
func (mf *MedianFinder) FindMedian() float64 {
	maxSize := mf.maxHeap.Len()
	minSize := mf.minHeap.Len()

	if maxSize == 0 && minSize == 0 {
		panic("no elements added yet")
	}

	// If total elements is odd
	if maxSize > minSize {
		return float64(mf.maxHeap.Top())
	} else if minSize > maxSize {
		return float64(mf.minHeap.Top())
	} else {
		// Even number of elements - average of two middle elements
		return float64(mf.maxHeap.Top()+mf.minHeap.Top()) / 2.0
	}
}
