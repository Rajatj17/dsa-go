package kwaymerge

import (
	"container/heap"
)

type Element struct {
	Value    int // The actual value
	ArrayIdx int // Which array this element came from
	ElemIdx  int // Position in that array
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinHeap for k-way merge using priority queue
type MinHeap []Element

func KWayMerge(arr [][]int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Add first element from each non-empty array
	for i, arr := range arr {
		if len(arr) > 0 {
			heap.Push(minHeap, Element{
				Value:    arr[0],
				ArrayIdx: i,
				ElemIdx:  0,
			})
		}
	}

	result := []int{}

	// Process elements until heap is empty
	for minHeap.Len() > 0 {
		// Get minimum element
		minElem := heap.Pop(minHeap).(Element)
		result = append(result, minElem.Value)

		// Add next element from the same array if exists
		nextIdx := minElem.ElemIdx + 1
		if nextIdx < len(arr[minElem.ArrayIdx]) {
			heap.Push(minHeap, Element{
				Value:    arr[minElem.ArrayIdx][nextIdx],
				ArrayIdx: minElem.ArrayIdx,
				ElemIdx:  nextIdx,
			})
		}
	}

	return result
}
