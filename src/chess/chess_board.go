package chess

import (
	"fmt"
)

type chessBoard [][]*chess

func (cb chessBoard) findTargetChessPosition(t chessType, c chessColor) []int8 {
	result := make([]int8, 10)
	result = result[:0]
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			if cb[row][col]._type == t && cb[row][col].color == c {
				result = append(result, int8(row), int8(col))
			}
		}
	}
	return result
}

func (cb chessBoard) clone() chessBoard {
	newBoard := make(chessBoard, 10)
	newBoard = newBoard[:0]
	for row := 0; row < _BOARD_ROW; row++ {
		cols := make([]*chess, 9)
		cols = cols[:0]
		for col := 0; col < _BOARD_COL; col++ {
			oldChess := cb[row][col]
			cols = append(cols, &chess{
				_type: oldChess._type,
				color: oldChess.color,
			})
		}
		newBoard = append(newBoard, cols)
	}
	return newBoard
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

func (cb chessBoard) dump() {
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
	fmt.Println("---------------------------------")
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			chess := cb[row][col]
			if chess._type == _CHESS_NULL {
				fmt.Print("　")
			} else {
				name := ""
				if chess.color == _COLOR_RED {
					name = redChessNames[chess._type - 1]
				} else {
					name = blackChessNames[chess._type - 1]
				}
				fmt.Print(name)
			}
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------")
}
