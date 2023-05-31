package utils

import (
	"github.com/lctzz540/chessaiserver/models"
)

func InitChessBoard() [][]models.Piece {
	var pieceListInit []models.Piece
	const white = true
	const left = true
	var row = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for i := 1; i <= 8; i++ {
		pieceListInit = append(pieceListInit, *models.Pawn(white, row[i-1]+"2"))
		pieceListInit = append(pieceListInit, *models.Pawn(!white, row[i-1]+"7"))
	}
	pieceListInit = append(pieceListInit, *models.King(white))
	pieceListInit = append(pieceListInit, *models.King(!white))

	pieceListInit = append(pieceListInit, *models.Queen(white))
	pieceListInit = append(pieceListInit, *models.Queen(!white))

	pieceListInit = append(pieceListInit, *models.Bishop(white, left))
	pieceListInit = append(pieceListInit, *models.Bishop(white, !left))
	pieceListInit = append(pieceListInit, *models.Bishop(!white, left))
	pieceListInit = append(pieceListInit, *models.Bishop(!white, !left))

	pieceListInit = append(pieceListInit, *models.Knight(white, left))
	pieceListInit = append(pieceListInit, *models.Knight(white, !left))
	pieceListInit = append(pieceListInit, *models.Knight(!white, left))
	pieceListInit = append(pieceListInit, *models.Knight(!white, !left))

	pieceListInit = append(pieceListInit, *models.Rock(white, left))
	pieceListInit = append(pieceListInit, *models.Rock(white, !left))
	pieceListInit = append(pieceListInit, *models.Rock(!white, left))
	pieceListInit = append(pieceListInit, *models.Rock(!white, !left))

	chessBoard := make([][]models.Piece, 8)
	for i := 0; i < 8; i++ {
		chessBoard[i] = make([]models.Piece, 8)
	}

	for _, piece := range pieceListInit {
		pieceCoordinate := piece.Coordinate
		chessBoard[pieceCoordinate[0]][pieceCoordinate[1]] = piece
	}

	return chessBoard
}
