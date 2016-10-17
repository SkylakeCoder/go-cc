package chess

type ChessMovementGenerator struct {

}

func NewChessMovementGenerator() *ChessMovementGenerator {
	return &ChessMovementGenerator {}
}

func (cmg *ChessMovementGenerator) GenerateMoves(chessBoard ChessBoard, color ChessColor) []ChessBoard {
	moves := []ChessBoard {}
	cmg.generateCarMoves(&moves, chessBoard, color)
	cmg.generateHorseMoves(&moves, chessBoard, color)
	return moves
}

func (cmg *ChessMovementGenerator) generateCarMove(outResult *[]ChessBoard, chessBoard ChessBoard, newRow, newCol, oldRow, oldCol int, selfColor ChessColor) bool {
	cType := chessBoard[newRow][newCol].Type
	cColor := chessBoard[newRow][newCol].Color
	if cColor == selfColor {
		return false
	} else if cType == CHESS_NULL || cColor != selfColor {
		newChessBoard := chessBoard.clone()
		newChess := newChessBoard[newRow][newCol]
		newChess.Type = CHESS_CAR
		newChess.Color = selfColor
		oldChess := newChessBoard[oldRow][oldCol]
		oldChess.Type = CHESS_NULL
		oldChess.Color = COLOR_NULL

		*outResult = append(*outResult, newChessBoard)

		if cType != CHESS_NULL && cColor != selfColor {
			return false
		}
	}
	return true
}

func (cmg *ChessMovementGenerator) generateCarMoves(outResult *[]ChessBoard, chessBoard ChessBoard, color ChessColor) {
	rowCols := chessBoard.findTargetChessPosition(CHESS_CAR, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		for forward := row - 1; forward >= 0; forward-- {
			if !cmg.generateCarMove(outResult, chessBoard, forward, col, row, col, color) {
				break
			}
		}
		for backward := row + 1; backward < BOARD_ROW; backward++ {
			if !cmg.generateCarMove(outResult, chessBoard, backward, col, row, col, color) {
				break
			}
		}
		for leftward := col - 1; leftward >= 0; leftward-- {
			if !cmg.generateCarMove(outResult, chessBoard, row, leftward, row, col, color) {
				break
			}
		}
		for rightward := col + 1; rightward < BOARD_COL; rightward++ {
			if !cmg.generateCarMove(outResult, chessBoard, row, rightward, row, col, color) {
				break
			}
		}
	}
}

func (cmg *ChessMovementGenerator) generateHorseMove(outResult *[]ChessBoard, chessBoard ChessBoard, newRow, newCol, oldRow, oldCol int, color ChessColor) {
	newChessBoard := chessBoard.clone()
	newChessBoard[newRow][newCol].Type = CHESS_HORSE
	newChessBoard[newRow][newCol].Color = color
	newChessBoard[oldRow][oldCol].Type = CHESS_NULL
	newChessBoard[oldRow][oldCol].Color = COLOR_NULL
	*outResult = append(*outResult, newChessBoard)
}

func (cmg *ChessMovementGenerator) generateHorseMoves(outResult *[]ChessBoard, chessBoard ChessBoard, color ChessColor) {
	rowCols := chessBoard.findTargetChessPosition(CHESS_HORSE, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		blockChess1 := chessBoard.visit(row - 1, col)
		blockChess2 := chessBoard.visit(row, col - 1)
		blockChess3 := chessBoard.visit(row + 1, col)
		blockChess4 := chessBoard.visit(row, col + 1)
		if blockChess1 != nil && blockChess1.Type == CHESS_NULL {
			cmg.generateHorseMove(outResult, chessBoard, row - 2, col - 1, row, col, color)
			cmg.generateHorseMove(outResult, chessBoard, row - 2, col + 1, row, col, color)
		}
		if blockChess2 != nil && blockChess2.Type == CHESS_NULL {
			cmg.generateHorseMove(outResult, chessBoard, row - 1, col - 2, row, col, color)
			cmg.generateHorseMove(outResult, chessBoard, row + 1, col - 2, row, col, color)
		}
		if blockChess3 != nil && blockChess3.Type == CHESS_NULL {
			cmg.generateHorseMove(outResult, chessBoard, row + 2, col - 1, row, col, color)
			cmg.generateHorseMove(outResult, chessBoard, row + 2, col + 1, row, col, color)
		}
		if blockChess4 != nil && blockChess4.Type == CHESS_NULL {
			cmg.generateHorseMove(outResult, chessBoard, row - 1, col + 2, row, col, color)
			cmg.generateHorseMove(outResult, chessBoard, row + 1, col + 2, row, col, color)
		}
	}
}

