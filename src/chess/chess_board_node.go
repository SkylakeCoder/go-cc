package chess

import (
	"fmt"
)

type nodeType byte
const (
	_NODE_TYPE_NULL nodeType = iota
	_NODE_TYPE_MAX
	_NODE_TYPE_MIN
)

type move struct {
	oldRow, oldCol int8
	newRow, newCol int8
	_type chessType
	color chessColor
}

type chessBoardNode struct {
	parent   *chessBoardNode
	move     move
	depth    int8
	value    int16
	valueNodeForDebug *chessBoardNode
	nodeType nodeType
	children *chessBoardNodeList
	discard  bool
	setValueCount uint16
	next *chessBoardNode
}

var _chessBoardNodeList *chessBoardNodeList = newChessBoardNodeList()
const _POOL_INCREASE_NUM = 50000000
var _getNodeNum = 0
var _returnNodeNum = 0

func getChessBoardNode() *chessBoardNode {
	_getNodeNum++
	if _chessBoardNodeList.len() <= 0 {
		fmt.Println("realloc node......node num:", _getNodeNum)
		for i := 0; i < _POOL_INCREASE_NUM; i++ {
			newNode := &chessBoardNode {
				next: nil,
			}
			_chessBoardNodeList.pushBack(newNode)
		}
	}
	node := _chessBoardNodeList.popFront()
	return node
}

func returnChessBoardNode(node *chessBoardNode) {
	_returnNodeNum++
	if node.children != nil {
		//!--node.children.clear()
	}
	_chessBoardNodeList.pushBack(node)
}

func clearChessBoardNodeCounter() {
	_getNodeNum, _returnNodeNum = 0, 0
}

var _tempMoves = make([]move, 10)
func (cbn *chessBoardNode) getCurrentChessBoard() chessBoard {
	_tempMoves = _tempMoves[:0]
	parentNode := cbn
	for parentNode != nil {
		_tempMoves = append(_tempMoves, parentNode.move)
		parentNode = parentNode.parent
	}
	board := currentBoard.clone()
	for i := len(_tempMoves) - 1; i >= 0; i-- {
		move := _tempMoves[i]
		newRow, newCol := move.newRow, move.newCol
		oldRow, oldCol := move.oldRow, move.oldCol
		board[newRow][newCol]._type = move._type
		board[newRow][newCol].color = move.color
		board[oldRow][oldCol]._type = _CHESS_NULL
		board[oldRow][oldCol].color = _COLOR_NULL
	}
	return board
}

func (cbn *chessBoardNode) setNodeType(nodeType nodeType) {
	cbn.nodeType = nodeType
	if nodeType == _NODE_TYPE_MAX {
		cbn.value = _MIN_VALUE
	} else {
		cbn.value = _MAX_VALUE
	}
}

func (cbn *chessBoardNode) setValue(v int16, nodeForDebug *chessBoardNode) {
	cbn.setValueCount++
	if cbn.nodeType == _NODE_TYPE_MAX {
		if v > cbn.value {
			cbn.value = v
			cbn.valueNodeForDebug = nodeForDebug
		}
	} else {
		if v < cbn.value {
			cbn.value = v
			cbn.valueNodeForDebug = nodeForDebug
		}
	}
	if cbn.parent != nil {
		brothers := cbn.parent.children
		if cbn.parent.nodeType == _NODE_TYPE_MIN {
			//for _, v := range brothers {
			for e := brothers.front(); e != nil; e = e.next {
				if e.value != _MIN_VALUE && e.value < cbn.value {
					cbn.discard = true
					break
				}
			}
		} else {
			//for _, v := range brothers {
			for e := brothers.front(); e != nil; e = e.next {
				if e.value != _MAX_VALUE && e.value > cbn.value {
					cbn.discard = true
					break
				}
			}
		}
		if int64(cbn.setValueCount) >= cbn.children.len() {
			cbn.parent.setValue(cbn.value, cbn)
		} else if cbn.isDiscard() {
			if cbn.parent.nodeType == _NODE_TYPE_MAX {
				cbn.parent.setValue(_MIN_VALUE, cbn)
			} else {
				cbn.parent.setValue(_MAX_VALUE, cbn)
			}
		}
	}
}

func (cbn *chessBoardNode) getValue() int16 {
	return cbn.value
}

func (cbn *chessBoardNode) isDiscard() bool {
	temp := cbn
	for temp != nil {
		if temp.discard {
			return true
		}
		temp = temp.parent
	}
	return false
}

func (cbn *chessBoardNode) nextNode() *chessBoardNode {
	return cbn.next
}
