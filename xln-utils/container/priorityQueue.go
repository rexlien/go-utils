package container

import (
	"container/heap"
	common "github.com/rexlien/go-utils/xln-utils/common"
	internal "github.com/rexlien/go-utils/xln-utils/internal"
)
/*
type Comparable interface {
	Less(j Comparable) bool
}


// An Item is something we manage in a priority queue.
type Item struct {
	index int
	Value Comparable
}

func (item *Item) SetIndex(index int) {
	item.index = index
}

func (item *Item) Index() int {
	return item.index
}
*/
//type TestType int

type PriorityQueue struct {

	internal *internal.PriorityQueue

}


// A PriorityQueue implements heap.Interface and holds Items.

/*
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Value.Less(pq[j].Value) //Priority() > pq[j].Priority()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].SetIndex(i)
	pq[j].SetIndex(j)
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	//item.index = n
	item.SetIndex(n)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil    // avoid memory leak
	item.SetIndex(-1) // for safety
	*pq = old[0 : n-1]
	return item
}
*/
// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *common.PqItem) {
	heap.Fix(pq.internal, item.Index())
}

func (pq *PriorityQueue) PopItem() *common.PqItem {
	return heap.Pop(pq.internal).(*common.PqItem)

}

func (pq* PriorityQueue) Enqueue(comparable common.Comparable) {
	heap.Push(pq.internal, &common.PqItem{Value: comparable})
}

func (pq *PriorityQueue) Dequeue() common.Comparable {
	return heap.Pop(pq.internal).(*common.PqItem).Value
}

func (pq *PriorityQueue) Len() int {
	return pq.internal.Len()
}


func NewPriorityQueue() *PriorityQueue {
	newPq := internal.NewQueue()
	heap.Init(&newPq)
	return &PriorityQueue{internal: &newPq}
}
