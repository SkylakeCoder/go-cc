package chess

type chessMovementGenerator struct {

}

func newChessMovementGenerator() *chessMovementGenerator {
	return &chessMovementGenerator{}
}

func (cmg *chessMovementGenerator) generateMoves(board chessBoard, color chessColor) []chessBoard {
	moves := []chessBoard{}
	cmg.generateCarMoves(&moves, board, color)
	cmg.generateHorseMoves(&moves, board, color)
	cmg.generateCannonMoves(&moves, board, color)
	cmg.generateElephantMoves(&moves, board, color)
	cmg.generateGuardMoves(&moves, board, color)
	cmg.generateKingMoves(&moves, board, color)
	cmg.generatePawnMoves(&moves, board, color)

	return moves
}

func (cmg *chessMovementGenerator) generateCarMove(outResult *[]chessBoard, board chessBoard, newRow, newCol, oldRow, oldCol int, selfColor chessColor) bool {
	cType := board[newRow][newCol]._type
	cColor := board[newRow][newCol].color
	if cColor == selfColor {
		return false
	} else if cType == _CHESS_NULL || cColor != selfColor {
		newChessBoard := board.clone()
		newChess := newChessBoard[newRow][newCol]
		newChess._type = _CHESS_CAR
		newChess.color = selfColor
		oldChess := newChessBoard[oldRow][oldCol]
		oldChess._type = _CHESS_NULL
		oldChess.color = _COLOR_NULL

		*outResult = append(*outResult, newChessBoard)

		if cType != _CHESS_NULL && cColor != selfColor {
			return false
		}
	}
	return true
}

func (cmg *chessMovementGenerator) generateCarMoves(outResult *[]chessBoard, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_CAR, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		for forward := row - 1; forward >= 0; forward-- {
			if !cmg.generateCarMove(outResult, board, forward, col, row, col, color) {
				break
			}
		}
		for backward := row + 1; backward < _BOARD_ROW; backward++ {
			if !cmg.generateCarMove(outResult, board, backward, col, row, col, color) {
				break
			}
		}
		for leftward := col - 1; leftward >= 0; leftward-- {
			if !cmg.generateCarMove(outResult, board, row, leftward, row, col, color) {
				break
			}
		}
		for rightward := col + 1; rightward < _BOARD_COL; rightward++ {
			if !cmg.generateCarMove(outResult, board, row, rightward, row, col, color) {
				break
			}
		}
	}
}

func (cmg *chessMovementGenerator) isRowColValid(row, col int) bool {
	if row < 0 || row >= _BOARD_ROW {
		return false
	}
	if col < 0 || col >= _BOARD_COL {
		return false
	}
	return true
}

func (cmg *chessMovementGenerator) generateHorseMove(outResult *[]chessBoard, board chessBoard, newRow, newCol, oldRow, oldCol int, color chessColor) {
	if !cmg.isRowColValid(newRow, newCol) {
		return
	}
	if board[newRow][newCol]._type != _CHESS_NULL &&
		board[newRow][newCol].color == color {
		return
	}
	newChessBoard := board.clone()
	newChessBoard[newRow][newCol]._type = _CHESS_HORSE
	newChessBoard[newRow][newCol].color = color
	newChessBoard[oldRow][oldCol]._type = _CHESS_NULL
	newChessBoard[oldRow][oldCol].color = _COLOR_NULL
	*outResult = append(*outResult, newChessBoard)
}

func (cmg *chessMovementGenerator) generateHorseMoves(outResult *[]chessBoard, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_HORSE, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		blockChess1 := board.visit(row - 1, col)
		blockChess2 := board.visit(row, col - 1)
		blockChess3 := board.visit(row + 1, col)
		blockChess4 := board.visit(row, col + 1)
		if blockChess1 != nil && blockChess1._type == _CHESS_NULL {
			cmg.generateHorseMove(outResult, board, row - 2, col - 1, row, col, color)
			cmg.generateHorseMove(outResult, board, row - 2, col + 1, row, col, color)
		}
		if blockChess2 != nil && blockChess2._type == _CHESS_NULL {
			cmg.generateHorseMove(outResult, board, row - 1, col - 2, row, col, color)
			cmg.generateHorseMove(outResult, board, row + 1, col - 2, row, col, color)
		}
		if blockChess3 != nil && blockChess3._type == _CHESS_NULL {
			cmg.generateHorseMove(outResult, board, row + 2, col - 1, row, col, color)
			cmg.generateHorseMove(outResult, board, row + 2, col + 1, row, col, color)
		}
		if blockChess4 != nil && blockChess4._type == _CHESS_NULL {
			cmg.generateHorseMove(outResult, board, row - 1, col + 2, row, col, color)
			cmg.generateHorseMove(outResult, board, row + 1, col + 2, row, col, color)
		}
	}
}

func (cmg *chessMovementGenerator) generateCannonMove(outResult *[]chessBoard, board chessBoard, newRow, newCol, oldRow, oldCol int, color chessColor) {
	newChessBoard := board.clone()
	newChessBoard[newRow][newCol]._type = _CHESS_CANNON
	newChessBoard[newRow][newCol].color = color
	newChessBoard[oldRow][oldCol]._type = _CHESS_NULL
	newChessBoard[oldRow][oldCol].color = _COLOR_NULL
	*outResult = append(*outResult, newChessBoard)
}

func (cmg *chessMovementGenerator) generateCannonMoves(outResult *[]chessBoard, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_CANNON, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		hasMeetAChess := false
		for forward := row - 1; forward >= 0; forward-- {
			chess := board[forward][col]
			if chess._type == _CHESS_NULL {
				if !hasMeetAChess {
					cmg.generateCannonMove(outResult, board, forward, col, row, col, color)
				}
			} else {
				if (hasMeetAChess) {
					if chess.color != color {
						cmg.generateCannonMove(outResult, board, forward, col, row, col, color)
					}
					break
				} else {
					hasMeetAChess = true
				}
			}
		}
		hasMeetAChess = false
		for backward := row + 1; backward < _BOARD_ROW; backward++ {
			chess := board[backward][col]
			if chess._type == _CHESS_NULL {
				if !hasMeetAChess {
					cmg.generateCannonMove(outResult, board, backward, col, row, col, color)
				}
			} else {
				if (hasMeetAChess) {
					if chess.color != color {
						cmg.generateCannonMove(outResult, board, backward, col, row, col, color)
					}
					break
				} else {
					hasMeetAChess = true
				}
			}
		}
		hasMeetAChess = false
		for leftward := col - 1; leftward >= 0; leftward-- {
			chess := board[row][leftward]
			if chess._type == _CHESS_NULL {
				if !hasMeetAChess {
					cmg.generateCannonMove(outResult, board, row, leftward, row, col, color)
				}
			} else {
				if (hasMeetAChess) {
					if chess.color != color {
						cmg.generateCannonMove(outResult, board, row, leftward, row, col, color)
					}
					break
				} else {
					hasMeetAChess = true
				}
			}
		}
		hasMeetAChess = false
		for rightward := col + 1; rightward < _BOARD_COL; rightward++ {
			chess := board[row][rightward]
			if chess._type == _CHESS_NULL {
				if !hasMeetAChess {
					cmg.generateCannonMove(outResult, board, row, rightward, row, col, color)
				}
			} else {
				if (hasMeetAChess) {
					if chess.color != color {
						cmg.generateCannonMove(outResult, board, row, rightward, row, col, color)
					}
					break
				} else {
					hasMeetAChess = true
				}
			}
		}
	}
}

func (cmg *chessMovementGenerator) generateElephantMove(outResult *[]chessBoard, board chessBoard, newRow, newCol, oldRow, oldCol int, color chessColor) {
	if color == _COLOR_RED {
		if newRow < 5 {
			return
		}
	} else {
		if newRow > 4 {
			return
		}
	}
	newChessBoard := board.clone()
	newChessBoard[newRow][newCol]._type = _CHESS_ELEPHANT
	newChessBoard[newRow][newCol].color = color
	newChessBoard[oldRow][oldCol]._type = _CHESS_NULL
	newChessBoard[oldRow][oldCol].color = _COLOR_NULL
	*outResult = append(*outResult, newChessBoard)
}

func (cmg *chessMovementGenerator) generateElephantMoves(outResult *[]chessBoard, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_ELEPHANT, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		blockChess1 := board.visit(row - 1, col - 1)
		blockChess2 := board.visit(row + 1, col - 1)
		blockChess3 := board.visit(row + 1, col + 1)
		blockChess4 := board.visit(row - 1, col + 1)
		if blockChess1 != nil && blockChess1._type == _CHESS_NULL {
			cmg.generateElephantMove(outResult, board, row - 2, col - 2, row, col, color)
		}
		if blockChess2 != nil && blockChess2._type == _CHESS_NULL {
			cmg.generateElephantMove(outResult, board, row + 2, col - 2, row, col, color)
		}
		if blockChess3 != nil && blockChess3._type == _CHESS_NULL {
			cmg.generateElephantMove(outResult, board, row + 2, col + 2, row, col, color)
		}
		if blockChess4 != nil && blockChess4._type == _CHESS_NULL {
			cmg.generateElephantMove(outResult, board, row - 2, col + 2, row, col, color)
		}
	}
}

func (cmg *chessMovementGenerator) generateGuardMove(outResult *[]chessBoard, board chessBoard, newRow, newCol, oldRow, oldCol int, color chessColor) {
	if newCol < 3 || newCol > 5 {
		return
	}
	if color == _COLOR_RED {
		if newRow < 7 {
			return
		}
	} else {
		if newRow > 2 {
			return
		}
	}
	newChessBoard := board.clone()
	newChessBoard[newRow][newCol]._type = _CHESS_GUARD
	newChessBoard[newRow][newCol].color = color
	newChessBoard[oldRow][oldCol]._type = _CHESS_NULL
	newChessBoard[oldRow][oldCol].color = _COLOR_NULL
	*outResult = append(*outResult, newChessBoard)
}

func (cmg *chessMovementGenerator) generateGuardMoves(outResult *[]chessBoard, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_GUARD, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		if c, valid := board.getChessColor(row - 1, col - 1); valid && c != color {
			cmg.generateGuardMove(outResult, board, row - 1, col - 1, row, col, color)
		}
		if c, valid := board.getChessColor(row + 1, col - 1); valid && c != color {
			cmg.generateGuardMove(outResult, board, row + 1, col - 1, row, col, color)
		}
		if c, valid := board.getChessColor(row + 1, col + 1); valid && c != color {
			cmg.generateGuardMove(outResult, board, row + 1, col + 1, row, col, color)
		}
		if c, valid := board.getChessColor(row - 1, col + 1); valid && c != color {
			cmg.generateGuardMove(outResult, board, row - 1, col + 1, row, col, color)
		}
	}
}

func (cmg *chessMovementGenerator) generateKingMove(outResult *[]chessBoard, board chessBoard, newRow, newCol, oldRow, oldCol int, color chessColor) {
	if newCol < 3 || newCol > 5 {
		return
	}
	if color == _COLOR_RED {
		if newRow < 7 {
			return
		}
	} else {
		if newRow > 2 {
			return
		}
	}
	newChessBoard := board.clone()
	newChessBoard[newRow][newCol]._type = _CHESS_KING
	newChessBoard[newRow][newCol].color = color
	newChessBoard[oldRow][oldCol]._type = _CHESS_NULL
	newChessBoard[oldRow][oldCol].color = _COLOR_NULL
	*outResult = append(*outResult, newChessBoard)
}

func (cmg *chessMovementGenerator) generateKingMoves(outResult *[]chessBoard, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_KING, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		if c, valid := board.getChessColor(row - 1, col); valid && c != color {
			cmg.generateKingMove(outResult, board, row - 1, col, row, col, color)
		}
		if c, valid := board.getChessColor(row, col - 1); valid && c != color {
			cmg.generateKingMove(outResult, board, row, col - 1, row, col, color)
		}
		if c, valid := board.getChessColor(row + 1, col); valid && c != color {
			cmg.generateKingMove(outResult, board, row + 1, col, row, col, color)
		}
		if c, valid := board.getChessColor(row, col + 1); valid && c != color {
			cmg.generateKingMove(outResult, board, row, col + 1, row, col, color)
		}
	}
}

func (cmg *chessMovementGenerator) generatePawnMove(outResult *[]chessBoard, board chessBoard, newRow, newCol, oldRow, oldCol int, color chessColor) {
	newChessBoard := board.clone()
	newChessBoard[newRow][newCol]._type = _CHESS_PAWN
	newChessBoard[newRow][newCol].color = color
	newChessBoard[oldRow][oldCol]._type = _CHESS_NULL
	newChessBoard[oldRow][oldCol].color = _COLOR_NULL
	*outResult = append(*outResult, newChessBoard)
}

func (cmg *chessMovementGenerator) generatePawnMoves(outResult *[]chessBoard, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_PAWN, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		if color == _COLOR_RED {
			if row >= 5 {
				if c, valid := board.getChessColor(row - 1, col); valid && c != color {
					cmg.generatePawnMove(outResult, board, row - 1, col, row, col, color)
				}
			} else {
				if c, valid := board.getChessColor(row - 1, col); valid && c != color {
					cmg.generatePawnMove(outResult, board, row - 1, col, row, col, color)
				}
				if c, valid := board.getChessColor(row, col - 1); valid && c != color {
					cmg.generatePawnMove(outResult, board, row, col - 1, row, col, color)
				}
				if c, valid := board.getChessColor(row, col + 1); valid && c != color {
					cmg.generatePawnMove(outResult, board, row, col + 1, row, col, color)
				}
			}
		} else {
			if row <= 4 {
				if c, valid := board.getChessColor(row + 1, col); valid && c != color {
					cmg.generatePawnMove(outResult, board, row + 1, col, row, col, color)
				}
			} else {
				if c, valid := board.getChessColor(row + 1, col); valid && c != color {
					cmg.generatePawnMove(outResult, board, row + 1, col, row, col, color)
				}
				if c, valid := board.getChessColor(row, col - 1); valid && c != color {
					cmg.generatePawnMove(outResult, board, row, col - 1, row, col, color)
				}
				if c, valid := board.getChessColor(row, col + 1); valid && c != color {
					cmg.generatePawnMove(outResult, board, row, col + 1, row, col, color)
				}
			}
		}
	}
}