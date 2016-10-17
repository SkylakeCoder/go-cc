package chess

type ChessMovementGenerator struct {

}

func NewChessMovementGenerator() *ChessMovementGenerator {
	return &ChessMovementGenerator {}
}

func (cmg *ChessMovementGenerator) GenerateMoves(chessBoard ChessBoard, color ChessColor) []ChessBoard {
	moves := []ChessBoard {}
	cmg.generateCarMoves(&moves, chessBoard, color)

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
	if color == COLOR_RED {
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
	} else {
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
}

