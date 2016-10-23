package chess

type chess struct {
	_type chessType
	color chessColor
}

func (c *chess) string() string {
	return _CHESS_TYPE_ARRAY[c._type] + _CHESS_COLOR_ARRAY[c.color]
}