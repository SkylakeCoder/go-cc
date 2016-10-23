package chess

import (
	"fmt"
)

type chessBoard [][]*chess

func (cb chessBoard) findTargetChessPosition(t chessType, c chessColor) []int {
	result := []int {}
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			if cb[row][col]._type == t && cb[row][col].color == c {
				result = append(result, row, col)
			}
		}
	}
	return result
}

func (cb chessBoard) clone() chessBoard {
	new := chessBoard{}
	for row := 0; row < _BOARD_ROW; row++ {
		cols := []*chess{}
		for col := 0; col < _BOARD_COL; col++ {
			oldChess := cb[row][col]
			cols = append(cols, &chess{
				_type: oldChess._type,
				color: oldChess.color,
			})
		}
		new = append(new, cols)
	}
	return new
}

func (cb chessBoard) visit(row, col int) *chess {
	if row < 0 || row >= _BOARD_ROW {
		return nil
	}
	if col < 0 || col >= _BOARD_COL {
		return nil
	}
	return cb[row][col]
}

func (cb chessBoard) getChessColor(row, col int) (chessColor, bool) {
	if row < 0 || row >= _BOARD_ROW {
		return _COLOR_NULL, false
	}
	if col < 0 || col >= _BOARD_COL {
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
