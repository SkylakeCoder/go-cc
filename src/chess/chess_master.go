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
	cm.chessBoard = [][]*Chess {}
	for row := 0; row < BOARD_ROW; row++ {
		cols := []*Chess {}
		for col := 0; col < BOARD_COL; col++ {
			cols = append(cols, &Chess { Type: CHESS_NULL, Color: COLOR_NULL })
		}
		cm.chessBoard = append(cm.chessBoard, cols)
	}
	// Black.
	cm.chessBoard[0][0] = &Chess {Type:CHESS_CAR, Color:COLOR_BLACK}
	cm.chessBoard[0][1] = &Chess {Type:CHESS_HORSE, Color:COLOR_BLACK}
	cm.chessBoard[0][2] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_BLACK}
	cm.chessBoard[0][3] = &Chess {Type:CHESS_GUARD, Color:COLOR_BLACK}
	cm.chessBoard[0][4] = &Chess {Type:CHESS_KING, Color:COLOR_BLACK}
	cm.chessBoard[0][5] = &Chess {Type:CHESS_GUARD, Color:COLOR_BLACK}
	cm.chessBoard[0][6] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_BLACK}
	cm.chessBoard[0][7] = &Chess {Type:CHESS_HORSE, Color:COLOR_BLACK}
	cm.chessBoard[0][8] = &Chess {Type:CHESS_CAR, Color:COLOR_BLACK}

	cm.chessBoard[2][1] = &Chess {Type:CHESS_CANNON, Color:COLOR_BLACK}
	cm.chessBoard[2][7] = &Chess {Type:CHESS_CANNON, Color:COLOR_BLACK}

	cm.chessBoard[3][0] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cm.chessBoard[3][2] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cm.chessBoard[3][4] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cm.chessBoard[3][6] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cm.chessBoard[3][8] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}

	// Red.
	cm.chessBoard[9][0] = &Chess {Type:CHESS_CAR, Color:COLOR_RED}
	cm.chessBoard[9][1] = &Chess {Type:CHESS_HORSE, Color:COLOR_RED}
	cm.chessBoard[9][2] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_RED}
	cm.chessBoard[9][3] = &Chess {Type:CHESS_GUARD, Color:COLOR_RED}
	cm.chessBoard[9][4] = &Chess {Type:CHESS_KING, Color:COLOR_RED}
	cm.chessBoard[9][5] = &Chess {Type:CHESS_GUARD, Color:COLOR_RED}
	cm.chessBoard[9][6] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_RED}
	cm.chessBoard[9][7] = &Chess {Type:CHESS_HORSE, Color:COLOR_RED}
	cm.chessBoard[9][8] = &Chess {Type:CHESS_CAR, Color:COLOR_RED}

	cm.chessBoard[7][1] = &Chess {Type:CHESS_CANNON, Color:COLOR_RED}
	cm.chessBoard[7][7] = &Chess {Type:CHESS_CANNON, Color:COLOR_RED}

	cm.chessBoard[6][0] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cm.chessBoard[6][2] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cm.chessBoard[6][4] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cm.chessBoard[6][6] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cm.chessBoard[6][8] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
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