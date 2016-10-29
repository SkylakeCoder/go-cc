package chess

import "log"

type chessBoardNodeList struct {
	head *chessBoardNode
	tail *chessBoardNode
	num int64
}

func newChessBoardNodeList() *chessBoardNodeList {
	return &chessBoardNodeList{
		head: &chessBoardNode{},
	}
}

func (cbnl *chessBoardNodeList) pushBack(node *chessBoardNode) {
	if cbnl.tail == nil {
		cbnl.head.next = node
	} else {
		cbnl.tail.next = node
	}
	cbnl.tail = node
	node.next = nil
	cbnl.num++
}

func (cbnl *chessBoardNodeList) pushFront(node *chessBoardNode) {
	node.next = cbnl.head.next
	cbnl.head.next = node
	if cbnl.tail == nil {
		cbnl.tail = node
	}
	cbnl.num++
}

func (cbnl *chessBoardNodeList) pushFrontSlice(slice []*chessBoardNode) {
	for i := len(slice) - 1; i >= 0; i-- {
		cbnl.pushFront(slice[i])
	}
}

func (cbnl *chessBoardNodeList) pushFrontList(list *chessBoardNodeList) {
	if list == nil || list.tail == nil {
		log.Fatalln("pushFrontList failed...")
	}
	list.tail.next = cbnl.head.next
	cbnl.head.next = list.head.next
	cbnl.num += list.num
	if cbnl.tail == nil {
		cbnl.tail = list.tail
	}
}

func (cbnl *chessBoardNodeList) popFront() *chessBoardNode {
	var node *chessBoardNode
	if cbnl.head.next != nil {
		node = cbnl.head.next
		cbnl.head.next = cbnl.head.next.next
		node.next = nil
		if node == cbnl.tail {
			cbnl.tail = nil
		}
		cbnl.num--
	}
	return node
}

func (cbnl *chessBoardNodeList) len() int64 {
	return cbnl.num
}

func (cbnl *chessBoardNodeList) front() *chessBoardNode {
	return cbnl.head.next
}

func (cbnl *chessBoardNodeList) clear() {
	cbnl.head.next = nil
	cbnl.tail = nil
	cbnl.num = 0
}