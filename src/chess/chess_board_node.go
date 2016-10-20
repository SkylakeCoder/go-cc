package chess

import "log"

type NodeType byte
const (
	NODE_TYPE_NULL = iota
	NODE_TYPE_MAX
	NODE_TYPE_MIN
)

type ChessBoardNode struct {
	chessBoard ChessBoard
	parent *ChessBoardNode
	depth int
	childValues []int
	nodeType NodeType
}

func (cbn *ChessBoardNode) AddChildValue(v int) {
	cbn.childValues = append(cbn.childValues, v)
}

func (cbn *ChessBoardNode) GetScore() int {
	var score int
	if cbn.nodeType == NODE_TYPE_MAX {
		score = 0
		for _, v := range cbn.childValues {
			if v > score {
				score = v
			}
		}
	} else if cbn.nodeType == NODE_TYPE_MIN {
		score = 1000000
		for _, v := range cbn.childValues {
			if v < score {
				score = v
			}
		}
	} else {
		log.Fatalln("wrong node type...")
	}
	return score
}

