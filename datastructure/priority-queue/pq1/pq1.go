package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Name string
	Expiry int
	Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// want to pop the item with lowest expiry, i.e. highest priority
	return pq[i].Expiry < pq[j].Expiry
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = j
	pq[j].Index = i
}

func main() {
	listItems := []*Item{
		{Name: "Carrot", Expiry: 30},
		{Name: "Potato", Expiry: 45},
		{Name: "Rice", Expiry: 100},
		{Name: "Spinach", Expiry: 5},
	}
	PriorityQueue := make(PriorityQueue, len(listItems))

	for i, item := range listItems {
		PriorityQueue[i] = item
		PriorityQueue[i].Index = i
	}

	heap.Init(&PriorityQueue)
	
	for PriorityQueue.Len() > 0 {
		item := heap.Pop(&PriorityQueue).(*Item)
		fmt.Printf("Name: %s Expiry: %d\n", item.Name, item.Expiry)
	}
}