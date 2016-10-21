package chess

import "log"

type NodeType byte
const (
	NODE_TYPE_NULL NodeType = iota
	NODE_TYPE_MAX
	NODE_TYPE_MIN
)

type ChessBoardNode struct {
	chessBoard ChessBoard
	parent *ChessBoardNode
	depth int
	value int
	value int
	nodeType NodeType
}

func (cbn *ChessBoardNode) SetNodeType(nodeType NodeType) {
	cbn.nodeType = nodeType
	if nodeType == NODE_TYPE_MAX {
		cbn.value = -1000000
	} else {
		cbn.value = 1000000
	}
}

func (cbn *ChessBoardNode) SetValue(v int) {
	if cbn.nodeType == NODE_TYPE_MAX {
		if v > cbn.value {
			cbn.value = v
		}
	} else if cbn.nodeType == NODE_TYPE_MIN {
		if v < cbn.value {
			cbn.value = v
		}
	} else {
		log.Fatalln("wrong node type...")
	}
}

func (cbn *ChessBoardNode) GetValue() int {
	return cbn.value
}

