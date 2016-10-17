package chess

type ChessBoard [][]*Chess

func (cb ChessBoard) findTargetChessPosition(t ChessType, c ChessColor) (int, int) {
	for row := 0; row < BOARD_ROW; row++ {
		for col := 0; col < BOARD_COL; col++ {
			if cb[row][col].Type == t && cb[row][col].Color == c {
				return row, col
			}
		}
	}
	return -1, -1
}
