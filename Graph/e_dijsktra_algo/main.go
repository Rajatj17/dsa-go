package dijsktraalgo

import (
	graph "DSA/Graph/a_basic"
	basic_heap "DSA/Heap/a__Basic_Code"
	"container/heap"
	"math"
)

/*
Graph Representation for Dijkstra's Algorithm:

     A ----5---- B
     |          /\
     |8        /  \
     |        /    \3
     |       /4   	\
     |      /        \
     |     /          \
     |    /            C
     |   /            /	|
     |  /            / 	|
     | /            /  	|
     |/             /		|1
     D    				6/		|
      \           /     |
       \7        /      |
        \       /       |
         \     /        |
          \   /         |
           \ /          |
            E ----------+

Vertices: A, B, C, D, E
Edges with weights:
- A to B: 5
- A to D: 8
- B to C: 3
- B to D: 4
- B to E: 2
- C to E: 6
- C to F: 1  (where F is an additional vertex)
- D to E: 7
- D to B: 10

*/

func Dijkstra(source int, destination int) int {
	g := graph.NewGraph(5, false)

	// Store distance array
	// I'm given a source & lets say starting distance is 0
	distance := make([]int, g.Vertices)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[source] = 0

	// Its kind of an greedy algorithm, so always pick the node which has shortest distance untill now
	pq := &basic_heap.PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &basic_heap.PriorityQueueItem{
		Value:    "A",
		Priority: 0,
	})

	visited := make([]bool, g.Vertices)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*basic_heap.PriorityQueueItem)

		currentNode := current.Value.(int)

		if visited[currentNode] {
			continue
		}

		visited[currentNode] = true

		// I need to traverse all the nbrs of the source & update their distance value from the source.
		for _, nbr := range g.AdjList[currentNode] {

			// I will now go to next neighbour, find nbrs of the tht node & see if the node is visited & if the distance from
			// this nbr node to a new node is less than what is currently we have, I will update its value to that distance value
			if visited[nbr.To] {
				newDist := nbr.Weight + distance[currentNode]
				if newDist < distance[nbr.To] {
					distance[nbr.To] = newDist
					heap.Push(pq, &basic_heap.PriorityQueueItem{
						Value:    nbr.To,
						Priority: nbr.Weight,
					})
				}
			}
		}
	}

	return distance[destination]
}
