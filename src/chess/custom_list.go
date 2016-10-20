package chess

import (
	"container/list"
	"log"
)

type MyList struct {
	*list.List
}

func NewMyList() *MyList {
	return &MyList { list.New() }
}

func (ml *MyList) PushFrontSlice(slice []*ChessBoardNode) {
	for _, v := range slice {
		ml.PushFront(v)
	}
}

func (ml *MyList) PushBackSlice(slice []*ChessBoardNode) {
	for _, v := range slice {
		ml.PushBack(v)
	}
}

func (ml *MyList) PopFront() *ChessBoardNode {
	v := ml.Front()
	ml.Remove(v)
	value, ok := v.Value.(*ChessBoardNode)
	if !ok {
		log.Fatalln("wrong type in MyList...")
	}
	return value
}