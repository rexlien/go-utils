package priority_queue_test

import (
	utils "github.com/rexlien/go-utils/xln-utils/container"
	"testing"
)

type item struct {
	value int
}

func (i *item) Less(j utils.Comparable) bool {
	return i.value < j.(*item).value
}

func TestPriority(t *testing.T) {

	pq := utils.NewPriorityQueue()
	pq.Enqueue(&item{value: 1})
	pq.Enqueue(&item{value: 0})

	front := pq.Dequeue()
	if front.(*item).value != 0 {
		t.FailNow()
	}

	front = pq.Dequeue()
	if front.(*item).value != 1 {
		t.FailNow()
	}



}