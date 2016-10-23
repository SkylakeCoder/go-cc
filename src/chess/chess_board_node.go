package chess

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
	children []*chessBoardNode
	discard bool
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
