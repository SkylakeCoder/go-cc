package chess

import (
	"log"
	"os"
	"strconv"
)

const (
	DEPTH = 3
)

type ChessMaster struct {
	chessBoard ChessBoard
}

func NewChessMaster() *ChessMaster {
	cm := &ChessMaster {}
	cm.InitChessBoard()
	return cm
}

func (cm *ChessMaster) InitChessBoard() {
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
	cm.chessBoard = [][]*Chess {}
	for row := 0; row < BOARD_ROW; row++ {
		cols := []*Chess {}
		for col := 0; col < BOARD_COL * 2; col+=2 {
			cols = append(cols, &Chess { Type: ChessType(initBoard[row][col + 1]), Color: ChessColor(initBoard[row][col]) })
		}
		cm.chessBoard = append(cm.chessBoard, cols)
	}
}

func (cm *ChessMaster) LoadChessBoard(value string) {
	if len(value) % 2 != 0 {
		log.Fatalln("error when LoadChessBoard...")
		os.Exit(1)
	}
	cm.chessBoard = [][]*Chess {}
	for row := 0; row < BOARD_ROW; row++ {
		cols := []*Chess {}
		for col := 0; col < BOARD_COL; col++ {
			cols = append(cols, &Chess { Type: CHESS_NULL, Color: COLOR_NULL })
		}
		cm.chessBoard = append(cm.chessBoard, cols)
	}
	idx := 0
	for i := 0; i < len(value); i+=2 {
		row := idx / BOARD_COL
		col := idx % BOARD_COL
		t, _ := strconv.Atoi(string(value[i]))
		c, _ := strconv.Atoi(string(value[i + 1]))
		cm.chessBoard[row][col].Type = ChessType(t)
		cm.chessBoard[row][col].Color = ChessColor(c)
		idx++
	}
}

func (cm *ChessMaster) Dump() {
	cm.chessBoard.dump()
}

func (cm *ChessMaster) Search(value string) string {
	return cm.search(value)
}

func (cm *ChessMaster) convertMoves(moves []ChessBoard, parentNode *ChessBoardNode, depth int, nodeType NodeType) []*ChessBoardNode {
	nodes := []*ChessBoardNode {}
	for _, v := range moves {
		node := &ChessBoardNode {
			chessBoard: v,
			parent: parentNode,
			depth: depth,
			childValues: []int {},
			nodeType: nodeType,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func (cm *ChessMaster) isAllWaitForEvalNode(nodes *MyList) bool {
	for e := nodes.Front(); e != nil; e = e.Next() {
		node, ok := e.Value.(*ChessBoardNode)
		if !ok {
			log.Fatalln("wrong type in MyList...")
		}
		if node.parent != nil {
			return false
		}
	}
	return true
}

func (cm *ChessMaster) search(value string) string {
	cm.LoadChessBoard(value)
	cm.Dump()
	currentColor := COLOR_BLACK
	mainQueue := NewMyList()
	waitForEvalQueue := NewMyList()
	evaluator := NewChessBoardEvaluator()
	generator := NewChessMovementGenerator()
	moves := generator.GenerateMoves(cm.chessBoard, currentColor)
	mainQueue.PushFrontSlice(cm.convertMoves(moves, nil, 1, NODE_TYPE_MIN))

	for mainQueue.Len() > 0 {
		node := mainQueue.PopFront()
		if node.depth < DEPTH {
			waitForEvalQueue.PushFront(node)
			nodeType := NODE_TYPE_NULL
			if currentColor == COLOR_BLACK {
				currentColor = COLOR_RED
				nodeType = NODE_TYPE_MAX
			} else {
				currentColor = COLOR_BLACK
				nodeType = NODE_TYPE_MIN
			}
			moves := generator.GenerateMoves(cm.chessBoard, currentColor)
			mainQueue.PushFrontSlice(cm.convertMoves(moves, node, node.depth + 1, nodeType))
		} else {
			v := evaluator.Eval(node.chessBoard)
			node.parent.AddChildValue(v)
		}
	}
	for waitForEvalQueue.Len() > 0 {
		if cm.isAllWaitForEvalNode(waitForEvalQueue) {
			break
		}
		node := waitForEvalQueue.PopFront()
		if node.parent == nil {
			waitForEvalQueue.PushBack(node)
		} else {
			node.parent.AddChildValue(node.GetScore())
		}
	}
	score := 0
	var targetNode *ChessBoardNode = nil
	for waitForEvalQueue.Len() > 0 {
		node := waitForEvalQueue.PopFront()
		nodeScore := node.GetScore()
		if nodeScore > score {
			score = nodeScore
			targetNode = node
		}
	}

	return targetNode.chessBoard.string()
}