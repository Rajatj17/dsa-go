package basic_heap

import (
	"container/heap"
	"fmt"
)

// Min Heap
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func (h *IntMinHeap) Top() int {
	if len(*h) == 0 {
		panic("heap is empty")
	}
	return (*h)[0]
}

func MinHeap() {
	minHeap := &IntMinHeap{2, 1, 5, 3}
	heap.Init(minHeap)
	fmt.Printf("Initial Heap %v\n", *minHeap)

	heap.Push(minHeap, 0)
	fmt.Printf("Initial Heap %v\n", *minHeap)
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func (h *IntMaxHeap) Top() int {
	if len(*h) == 0 {
		panic("heap is empty")
	}
	return (*h)[0]
}

func MaxHeap() {

}

type PriorityQueueItem struct {
	Value    interface{}
	Priority int
	Index    int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)

	item := x.(*PriorityQueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Update(item *PriorityQueueItem, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority

	heap.Fix(pq, item.Index)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil
	item.Index = -1 // For safety

	*pq = old[0 : n-1]

	return item
}
