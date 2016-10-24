package chess

type chessMovementGenerator struct {

}

func newChessMovementGenerator() *chessMovementGenerator {
	return &chessMovementGenerator{}
}

var _tempMoveResult = make([]move, 100)
func (cmg *chessMovementGenerator) generateMoves(board chessBoard, color chessColor) []move {
	_tempMoveResult = _tempMoveResult[:0]
	cmg.generateCarMoves(&_tempMoveResult, board, color)
	cmg.generateHorseMoves(&_tempMoveResult, board, color)
	cmg.generateCannonMoves(&_tempMoveResult, board, color)
	cmg.generateElephantMoves(&_tempMoveResult, board, color)
	cmg.generateGuardMoves(&_tempMoveResult, board, color)
	cmg.generateKingMoves(&_tempMoveResult, board, color)
	cmg.generatePawnMoves(&_tempMoveResult, board, color)

	return _tempMoveResult
}

func (cmg *chessMovementGenerator) generateCarMove(outResult *[]move, board chessBoard, newRow, newCol, oldRow, oldCol int8, selfColor chessColor) bool {
	cType := board[newRow][newCol]._type
	cColor := board[newRow][newCol].color
	if cColor == selfColor {
		return false
	} else if cType == _CHESS_NULL || cColor != selfColor {
		*outResult = append(*outResult, move {
			newRow: newRow, newCol: newCol,
			oldRow: oldRow, oldCol: oldCol,
			_type: _CHESS_CAR,
			color: selfColor,
		})
		if cType != _CHESS_NULL && cColor != selfColor {
			return false
		}
	}
	return true
}

func (cmg *chessMovementGenerator) generateCarMoves(outResult *[]move, board chessBoard, color chessColor) {
	rowCols := board.findTargetChessPosition(_CHESS_CAR, color)
	for i := 0; i < len(rowCols); i+=2 {
		row := rowCols[i]
		col := rowCols[i + 1]
		for forward := row - 1; forward >= 0; forward-- {
			if !cmg.generateCarMove(outResult, board, forward, col, row, col, color) {
				break
			}
		}
		for backward := row + 1; backward < int8(_BOARD_ROW); backward++ {
			if !cmg.generateCarMove(outResult, board, backward, col, row, col, color) {
				break
			}
		}
		for leftward := col - 1; leftward >= 0; leftward-- {
			if !cmg.generateCarMove(outResult, board, row, leftward, row, col, color) {
				break
			}
		}
		for rightward := col + 1; rightward < int8(_BOARD_COL); rightward++ {
			if !cmg.generateCarMove(outResult, board, row, rightward, row, col, color) {
				break
			}
		}
	}
}

func (cmg *chessMovementGenerator) isRowColValid(row, col int8) bool {
	if row < 0 || row >= int8(_BOARD_ROW) {
		return false
	}
	if col < 0 || col >= int8(_BOARD_COL) {
		return false
	}
	return true
}

func (cmg *chessMovementGenerator) generateHorseMove(outResult *[]move, board chessBoard, newRow, newCol, oldRow, oldCol int8, color chessColor) {
	if !cmg.isRowColValid(newRow, newCol) {
		return
	}
	if board[newRow][newCol]._type != _CHESS_NULL &&
		board[newRow][newCol].color == color {
		return
	}
	*outResult = append(*outResult, move {
		newRow: newRow, newCol: newCol,
		oldRow: oldRow, oldCol: oldCol,
		_type: _CHESS_HORSE,
		color: color,
	})
}

func (cmg *chessMovementGenerator) generateHorseMoves(outResult *[]move, board chessBoard, color chessColor) {
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

func (cmg *chessMovementGenerator) generateCannonMove(outResult *[]move, board chessBoard, newRow, newCol, oldRow, oldCol int8, color chessColor) {
	*outResult = append(*outResult, move {
		newRow: newRow, newCol: newCol,
		oldRow: oldRow, oldCol: oldCol,
		_type: _CHESS_CANNON,
		color: color,
	})
}

func (cmg *chessMovementGenerator) generateCannonMoves(outResult *[]move, board chessBoard, color chessColor) {
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
		for backward := row + 1; backward < int8(_BOARD_ROW); backward++ {
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
		for rightward := col + 1; rightward < int8(_BOARD_COL); rightward++ {
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

func (cmg *chessMovementGenerator) generateElephantMove(outResult *[]move, board chessBoard, newRow, newCol, oldRow, oldCol int8, color chessColor) {
	if color == _COLOR_RED {
		if newRow < 5 {
			return
		}
	} else {
		if newRow > 4 {
			return
		}
	}
	*outResult = append(*outResult, move {
		newRow: newRow, newCol: newCol,
		oldRow: oldRow, oldCol: oldCol,
		_type: _CHESS_ELEPHANT,
		color: color,
	})
}

func (cmg *chessMovementGenerator) generateElephantMoves(outResult *[]move, board chessBoard, color chessColor) {
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

func (cmg *chessMovementGenerator) generateGuardMove(outResult *[]move, board chessBoard, newRow, newCol, oldRow, oldCol int8, color chessColor) {
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
	*outResult = append(*outResult, move {
		newRow: newRow, newCol: newCol,
		oldRow: oldRow, oldCol: oldCol,
		_type: _CHESS_GUARD,
		color: color,
	})
}

func (cmg *chessMovementGenerator) generateGuardMoves(outResult *[]move, board chessBoard, color chessColor) {
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

func (cmg *chessMovementGenerator) generateKingMove(outResult *[]move, board chessBoard, newRow, newCol, oldRow, oldCol int8, color chessColor) {
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
	*outResult = append(*outResult, move {
		newRow: newRow, newCol: newCol,
		oldRow: oldRow, oldCol: oldCol,
		_type: _CHESS_KING,
		color: color,
	})
}

func (cmg *chessMovementGenerator) generateKingMoves(outResult *[]move, board chessBoard, color chessColor) {
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

func (cmg *chessMovementGenerator) generatePawnMove(outResult *[]move, board chessBoard, newRow, newCol, oldRow, oldCol int8, color chessColor) {
	*outResult = append(*outResult, move {
		newRow: newRow, newCol: newCol,
		oldRow: oldRow, oldCol: oldCol,
		_type: _CHESS_PAWN,
		color: color,
	})
}

func (cmg *chessMovementGenerator) generatePawnMoves(outResult *[]move, board chessBoard, color chessColor) {
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