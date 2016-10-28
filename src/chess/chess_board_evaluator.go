package chess

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type chessBoardEvaluator struct {

}

var chessPieceValueArray = []int16 {
	0, 1000, 450, 450, 200, 200, 10000, 60,
}

var chessPiecePositionValueArray = [][][][]int16 {}

func newChessBoardEvaluator(pvPath string) *chessBoardEvaluator {
	evaluator := &chessBoardEvaluator{}
	evaluator.loadPositionValueConfig(pvPath)
	return evaluator
}

func (cbe *chessBoardEvaluator) loadPositionValueConfig(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("loadPositionValueConfig failed...")
	}
	err = json.Unmarshal(bytes, &chessPiecePositionValueArray)
	if err != nil {
		log.Fatalln("position value config is invalid...")
	}
}

func (cbe *chessBoardEvaluator) eval(chessBoard chessBoard) int16 {
	rv1, bv1 := cbe.evalPieceValue(chessBoard)
	rv2, bv2 := cbe.evalPositionValue(chessBoard)
	return (bv1 + bv2) - (rv1 + rv2)
}

func (cbe *chessBoardEvaluator) evalPieceValue(chessBoard chessBoard) (int16, int16) {
	var rResult int16
	var bResult int16
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			chess := chessBoard[row][col]
			if chess.color == _COLOR_RED {
				rResult += chessPieceValueArray[chess._type]
			} else if chess.color == _COLOR_BLACK {
				bResult += chessPieceValueArray[chess._type]
			}
		}
	}
	return rResult, bResult
}

func (cbe *chessBoardEvaluator) evalPositionValue(chessBoard chessBoard) (int16, int16) {
	var rResult int16
	var bResult int16
	for row := 0; row < _BOARD_ROW; row++ {
		for col := 0; col < _BOARD_COL; col++ {
			chess := chessBoard[row][col]
			if chess.color == _COLOR_RED {
				rResult += chessPiecePositionValueArray[chess.color][chess._type][row][col]
			} else if chess.color == _COLOR_BLACK {
				bResult += chessPiecePositionValueArray[chess.color][chess._type][row][col]
			}
		}
	}
	return rResult, bResult
}
