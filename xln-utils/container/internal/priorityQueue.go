package internal

import priority_queue "github.com/rexlien/go-utils/xln-utils/container"



//type PriorityQueue []*priorityqueue.Item
type PriorityQueue []*priority_queue.Item
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
	item := x.(*priority_queue.Item)
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

func NewQueue() PriorityQueue {

	return make(PriorityQueue, 0)

}