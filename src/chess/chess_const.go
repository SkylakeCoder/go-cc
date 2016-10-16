package chess

type ChessType byte
type ChessColor byte

const (
	BOARD_ROW int = 10
	BOARD_COL int = 9
)

const (
	CHESS_NULL ChessType = iota
	CHESS_CAR
	CHESS_HOUSE
	CHESS_CANNON
	CHESS_ELEPHANT
	CHESS_GUARD
	CHESS_KING
	CHESS_PAWN
)

const (
	COLOR_NULL ChessColor = iota
	COLOR_RED
	COLOR_BLACK
)