package chess

type chessBoardNodeList struct {
	head *chessBoardNode
	tail *chessBoardNode
	num uint64
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

func (cbnl *chessBoardNodeList) len() uint64 {
	return cbnl.num
}

func (cbnl *chessBoardNodeList) front() *chessBoardNode {
	return cbnl.head.next
}