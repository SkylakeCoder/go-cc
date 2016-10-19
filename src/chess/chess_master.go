package chess

import (
	"log"
	"os"
	"fmt"
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
		cm.chessBoard[row][col].Type = ChessType(value[i])
		cm.chessBoard[row][col].Color = ChessColor(value[i + 1])
		idx++
	}
}

func (cm *ChessMaster) Dump() {
	cm.chessBoard.dump()
}

func (cm *ChessMaster) TestSomething() {
	cm.Dump()
	evaluator := NewChessBoardEvaluator()
	generator := NewChessMovementGenerator()
	moves := generator.GenerateMoves(cm.chessBoard, COLOR_RED)
	for _, board := range moves {
		board.dump()
		fmt.Printf("eval-value=%d\n", evaluator.Eval(board))
	}
}