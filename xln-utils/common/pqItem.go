package common

type Comparable interface {
	Less(j Comparable) bool
}


type PqItem struct {
	index int
	Value Comparable

}

func (item *PqItem) SetIndex(index int) {
	item.index = index
}

func (item *PqItem) Index() int {
	return item.index
}
