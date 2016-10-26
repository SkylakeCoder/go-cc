package chess

type chessType byte
type chessColor byte

const (
	_BOARD_ROW int = 10
	_BOARD_COL int = 9
)

const (
	_CHESS_NULL chessType = iota
	_CHESS_CAR
	_CHESS_HORSE
	_CHESS_CANNON
	_CHESS_ELEPHANT
	_CHESS_GUARD
	_CHESS_KING
	_CHESS_PAWN
)

const (
	_COLOR_NULL chessColor = iota
	_COLOR_RED
	_COLOR_BLACK
)

var _CHESS_TYPE_ARRAY = []string { "0", "1", "2", "3", "4", "5", "6", "7" }
var _CHESS_COLOR_ARRAY = []string { "0", "1", "2" }

const (
	_MIN_VALUE = -10000000
	_MAX_VALUE = 10000000
)