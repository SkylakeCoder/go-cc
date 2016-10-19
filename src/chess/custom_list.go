package chess

import (
	"container/list"
)

type MyList struct {
	*list.List
}

func NewMyList() *MyList {
	return &MyList { list.New() }
}

func (ml *MyList) PushFrontSlice(slice []ChessBoard) {
	for _, v := range slice {
		ml.PushFront(v)
	}
}

func (ml *MyList) PushBackSlice(slice []ChessBoard) {
	for _, v := range slice {
		ml.PushBack(v)
	}
}