package chess

import "fmt"

type chess struct {
	_type chessType
	color chessColor
}

func (c *chess) string() string {
	return fmt.Sprintf("%d%d", c._type, c.color)
}