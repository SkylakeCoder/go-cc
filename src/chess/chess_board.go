package chess

import (
	"fmt"
)

type chessBoard [][]*chess

var _chessPositionResult = make([]int8, 10)
var _cloneBoard chessBoard

func init() {
	_cloneBoard = chessBoard{}
	for row := 0; row < _BOARD_ROW; row++ {
		cols := []*chess {}
		for col := 0; col < _BOARD_COL; col++ {
			cols = append(cols, &chess {
				_type: _CHESS_NULL,
				color: _COLOR_NULL,
			})
		}
		_cloneBoard = append(_cloneBoard, cols)
	}
}

func (cb chessBoard) findTargetChessPosition(t chessType, c chessColor) []int8 {
	_chessPositionResult = _chessPositionResult[:0]
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			if cb[row][col]._type == t && cb[row][col].color == c {
				_chessPositionResult = append(_chessPositionResult, int8(row), int8(col))
			}
		}
	}
	return _chessPositionResult
}

func (cb chessBoard) clone() chessBoard {
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			oldChess := cb[row][col]
			_cloneBoard[row][col]._type = oldChess._type
			_cloneBoard[row][col].color = oldChess.color
		}
	}
	return _cloneBoard
}

func (cb chessBoard) visit(row, col int8) *chess {
	if row < 0 || row >= int8(_BOARD_ROW) {
		return nil
	}
	if col < 0 || col >= int8(_BOARD_COL) {
		return nil
	}
	return cb[row][col]
}

func (cb chessBoard) getChessColor(row, col int8) (chessColor, bool) {
	if row < 0 || row >= int8(_BOARD_ROW) {
		return _COLOR_NULL, false
	}
	if col < 0 || col >= int8(_BOARD_COL) {
		return _COLOR_NULL, false
	}
	return cb[row][col].color, true
}

func (cb chessBoard) string() string {
	result := ""
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			result += cb[row][col].string()
		}
	}
	return result
}

func (cb chessBoard) dumpString() string {
	result := ""
	redChessNames := []string {
		"车", "马", "炮",
		"相", "仕", "帅",
		"兵",
	}
	blackChessNames := []string {
		"車", "馬", "砲",
		"象", "士", "將",
		"卒",
	}
	result += "---------------------------------\n"
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			chess := cb[row][col]
			if chess._type == _CHESS_NULL {
				result += "　"
			} else {
				name := ""
				if chess.color == _COLOR_RED {
					name = redChessNames[chess._type - 1]
				} else {
					name = blackChessNames[chess._type - 1]
				}
				result += name
			}
		}
		result += "\n"
	}
	result += "---------------------------------\n"
	return result
}

func (cb chessBoard) dump() {
	fmt.Print(cb.dumpString())
}
