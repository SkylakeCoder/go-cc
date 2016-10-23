package chess

import "log"

type nodeType byte
const (
	_NODE_TYPE_NULL nodeType = iota
	_NODE_TYPE_MAX
	_NODE_TYPE_MIN
)

type chessBoardNode struct {
	chessBoard chessBoard
	parent     *chessBoardNode
	depth      int
	value      int
	nodeType   nodeType
}

func (cbn *chessBoardNode) setNodeType(nodeType nodeType) {
	cbn.nodeType = nodeType
	if nodeType == _NODE_TYPE_MAX {
		cbn.value = _MIN_VALUE
	} else {
		cbn.value = _MAX_VALUE
	}
}

func (cbn *chessBoardNode) setValue(v int) {
	if cbn.nodeType == _NODE_TYPE_MAX {
		if v > cbn.value {
			cbn.value = v
		}
	} else if cbn.nodeType == _NODE_TYPE_MIN {
		if v < cbn.value {
			cbn.value = v
		}
	} else {
		log.Fatalln("wrong node type...")
	}
}

func (cbn *chessBoardNode) getValue() int {
	return cbn.value
}

