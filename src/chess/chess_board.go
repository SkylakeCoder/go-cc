package chess

import (
	"fmt"
)

type ChessBoard [][]*Chess

func (cb ChessBoard) findTargetChessPosition(t ChessType, c ChessColor) []int {
	result := []int {}
	for row := 0; row < BOARD_ROW; row++ {
		for col := 0; col < BOARD_COL; col++ {
			if cb[row][col].Type == t && cb[row][col].Color == c {
				result = append(result, row, col)
			}
		}
	}
	return result
}

func (cb ChessBoard) clone() ChessBoard {
	new := ChessBoard {}
	for row := 0; row < BOARD_ROW; row++ {
		cols := []*Chess {}
		for col := 0; col < BOARD_COL; col++ {
			oldChess := cb[row][col]
			cols = append(cols, &Chess {
				Type: oldChess.Type,
				Color: oldChess.Color,
			})
		}
		new = append(new, cols)
	}
	return new
}

func (cb ChessBoard) dump() {
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
	for row := 0; row < BOARD_ROW; row++ {
		for col := 0; col < BOARD_COL; col++ {
			chess := cb[row][col]
			if chess.Type == CHESS_NULL {
				fmt.Print("　")
			} else {
				name := ""
				if chess.Color == COLOR_RED {
					name = redChessNames[chess.Type - 1]
				} else {
					name = blackChessNames[chess.Type - 1]
				}
				fmt.Print(name)
			}
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------")
}
