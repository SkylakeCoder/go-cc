package chess

import (
	"container/list"
	"log"
)

type myList struct {
	*list.List
}

func newMyList() *myList {
	return &myList{ list.New() }
}

func (ml *myList) pushFrontSlice(slice []*chessBoardNode) {
	for _, v := range slice {
		ml.PushFront(v)
	}
}

func (ml *myList) pushBackSlice(slice []*chessBoardNode) {
	for _, v := range slice {
		ml.PushBack(v)
	}
}

func (ml *myList) popFront() *chessBoardNode {
	v := ml.Front()
	ml.Remove(v)
	value, ok := v.Value.(*chessBoardNode)
	if !ok {
		log.Fatalln("wrong type in MyList...")
	}
	return value
}