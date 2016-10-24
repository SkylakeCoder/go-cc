package chess

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	_DEPTH = 4
)

type chessMaster struct {
	chessBoard chessBoard
}

func newChessMaster() *chessMaster {
	cm := &chessMaster{}
	cm.initChessBoard()
	return cm
}

func (cm *chessMaster) initChessBoard() {
	initBoard := [][]byte {
		{ 2, 1, 2, 2, 2, 4, 2, 5, 2, 6, 2, 5, 2, 4, 2, 2, 2, 1 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 3, 0, 0 },
		{ 2, 7, 0, 0, 2, 7, 0, 0, 2, 7, 0, 0, 2, 7, 0, 0, 2, 7 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 1, 7, 0, 0, 1, 7, 0, 0, 1, 7, 0, 0, 1, 7, 0, 0, 1, 7 },
		{ 0, 0, 1, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 1, 1, 1, 2, 1, 4, 1, 5, 1, 6, 1, 5, 1, 4, 1, 2, 1, 1 },
	}
	cm.chessBoard = [][]*chess{}
	for row := 0; row < _BOARD_ROW; row++ {
		cols := []*chess{}
		for col := 0; col < _BOARD_COL * 2; col+=2 {
			cols = append(cols, &chess{ _type: chessType(initBoard[row][col + 1]), color: chessColor(initBoard[row][col]) })
		}
		cm.chessBoard = append(cm.chessBoard, cols)
	}
}

func (cm *chessMaster) loadChessBoard(value string) {
	if len(value) % 2 != 0 {
		log.Fatalln("error when LoadChessBoard...")
		os.Exit(1)
	}
	cm.chessBoard = [][]*chess{}
	for row := 0; row < _BOARD_ROW; row++ {
		cols := []*chess{}
		for col := 0; col < _BOARD_COL; col++ {
			cols = append(cols, &chess{ _type: _CHESS_NULL, color: _COLOR_NULL })
		}
		cm.chessBoard = append(cm.chessBoard, cols)
	}
	idx := 0
	for i := 0; i < len(value); i+=2 {
		row := idx / _BOARD_COL
		col := idx % _BOARD_COL
		t, _ := strconv.Atoi(string(value[i]))
		c, _ := strconv.Atoi(string(value[i + 1]))
		cm.chessBoard[row][col]._type = chessType(t)
		cm.chessBoard[row][col].color = chessColor(c)
		idx++
	}
}

func (cm *chessMaster) dump() {
	cm.chessBoard.dump()
}

func (cm *chessMaster) convertMoves(moves []move, parentNode *chessBoardNode, board chessBoard, depth int8, nodeType nodeType) []*chessBoardNode {
	nodes := make([]*chessBoardNode, 100)
	nodes = nodes[:0]
	for _, v := range moves {
		node := &chessBoardNode{
			board: board,
			move: v,
			parent: parentNode,
			depth: depth,
		}
		node.setNodeType(nodeType)
		node.discard = false
		nodes = append(nodes, node)
	}
	if parentNode != nil {
		parentNode.children = nodes
	}
	return nodes
}

func (cm *chessMaster) isAllWaitForEvalNode(nodes *myList) bool {
	for e := nodes.Front(); e != nil; e = e.Next() {
		node, ok := e.Value.(*chessBoardNode)
		if !ok {
			log.Fatalln("wrong type in MyList...")
		}
		if node.parent != nil {
			return false
		}
	}
	return true
}

func (cm *chessMaster) search(value string) string {
	st := time.Now()
	cm.loadChessBoard(value)
	cm.dump()
	mainQueue := newMyList()
	waitForEvalQueue := newMyList()
	evaluator := newChessBoardEvaluator()
	generator := newChessMovementGenerator()
	moves := make([]move, 100)
	moves = moves[:0]
	moves = generator.generateMoves(cm.chessBoard, _COLOR_BLACK)
	mainQueue.pushFrontSlice(cm.convertMoves(moves, nil, cm.chessBoard, 1, _NODE_TYPE_MIN))

	for mainQueue.Len() > 0 {
		node := mainQueue.popFront()
		if node.depth < _DEPTH {
			if node.isDiscard() {
				continue
			}
			waitForEvalQueue.PushFront(node)
			nodeType := _NODE_TYPE_NULL
			color := _COLOR_NULL
			if node.nodeType == _NODE_TYPE_MIN {
				nodeType = _NODE_TYPE_MAX
				color = _COLOR_RED
			} else {
				nodeType = _NODE_TYPE_MIN
				color = _COLOR_BLACK
			}
			moves = moves[:0]
			moves = generator.generateMoves(node.getCurrentChessBoard(), color)
			mainQueue.pushFrontSlice(cm.convertMoves(moves, node, nil, node.depth + 1, nodeType))
		} else {
			v := evaluator.eval(node.getCurrentChessBoard())
			node.parent.setValue(v)
		}
	}
	for waitForEvalQueue.Len() > 0 {
		if cm.isAllWaitForEvalNode(waitForEvalQueue) {
			break
		}
		node := waitForEvalQueue.popFront()
		if node.isDiscard() {
			continue
		}
		if node.parent == nil {
			waitForEvalQueue.PushBack(node)
		} else {
			node.parent.setValue(node.getValue())
		}
	}
	score := _MIN_VALUE
	var targetNode *chessBoardNode = nil
	for waitForEvalQueue.Len() > 0 {
		node := waitForEvalQueue.popFront()
		nodeScore := node.getValue()
		if nodeScore > score {
			score = nodeScore
			targetNode = node
		}
	}

	if targetNode == nil {
		log.Fatalln("search targetNode == nil...")
	}
	log.Printf("time cost: %f s", time.Since(st).Seconds())
	return targetNode.getCurrentChessBoard().string()
}