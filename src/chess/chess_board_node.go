package chess

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
	board    chessBoard
	parent   *chessBoardNode
	move     move
	depth    int8
	value    int
	nodeType nodeType
	children []*chessBoardNode
	discard  bool
}

var _tempMoves = make([]move, 10)
func (cbn *chessBoardNode) getCurrentChessBoard() chessBoard {
	_tempMoves = _tempMoves[:0]
	parentNode := cbn
	var topNode *chessBoardNode
	for parentNode != nil {
		topNode = parentNode
		_tempMoves = append(_tempMoves, parentNode.move)
		parentNode = parentNode.parent
	}
	board := topNode.board.clone()
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

func (cbn *chessBoardNode) setValue(v int) {
	if cbn.nodeType == _NODE_TYPE_MAX {
		if v > cbn.value {
			cbn.value = v
		}
	} else {
		if v < cbn.value {
			cbn.value = v
		}
	}
	if cbn.parent != nil {
		brothers := cbn.parent.children
		if cbn.parent.nodeType == _NODE_TYPE_MIN {
			for _, v := range brothers {
				if v.value < cbn.value {
					v.discard = true
					break
				}
			}
		} else {
			for _, v := range brothers {
				if v.value > cbn.value {
					v.discard = true
					break
				}
			}
		}
	}
}

func (cbn *chessBoardNode) getValue() int {
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
