package chess

type ChessMovementGenerator struct {

}

func NewChessMovementGenerator() *ChessMovementGenerator {
	return &ChessMovementGenerator {}
}

func (cmg *ChessMovementGenerator) GenerateMoves(chessBoard ChessBoard) []ChessBoard {
	moves := []ChessBoard {}
	return moves
}

func (cmg *ChessMovementGenerator) generateCarMoves() []ChessBoard {
	return nil
}