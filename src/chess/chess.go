package chess

import "fmt"

type Chess struct {
	Type ChessType
	Color ChessColor
}

func (c *Chess) String() string {
	return fmt.Sprintf("%d%d", c.Type, c.Color)
}