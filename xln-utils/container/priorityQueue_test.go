package container

import (
	"github.com/rexlien/go-utils/xln-utils/common"
	//"github.com/rexlien/go-utils/xln-utils/container"
	"testing"
)

type item struct {
	value int
}

func (i *item) Less(j common.Comparable) bool {
	return i.value < j.(*item).value
}

func TestPriority(t *testing.T) {


	pq := NewPriorityQueue()
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