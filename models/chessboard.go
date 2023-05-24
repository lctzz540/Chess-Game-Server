package models

func InitChessBoard() []Piece {
	var chessBoard []Piece
	const white = true
	const left = true
	var row = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for i := 1; i <= 8; i++ {
		chessBoard = append(chessBoard, *Pawn(white, row[i-1]+"2"))
		chessBoard = append(chessBoard, *Pawn(!white, row[i-1]+"7"))
	}
	chessBoard = append(chessBoard, *King(white))
	chessBoard = append(chessBoard, *King(!white))

	chessBoard = append(chessBoard, *Queen(white))
	chessBoard = append(chessBoard, *Queen(!white))

	chessBoard = append(chessBoard, *Bishop(white, left))
	chessBoard = append(chessBoard, *Bishop(white, !left))
	chessBoard = append(chessBoard, *Bishop(!white, left))
	chessBoard = append(chessBoard, *Bishop(!white, !left))

	chessBoard = append(chessBoard, *Knight(white, left))
	chessBoard = append(chessBoard, *Knight(white, !left))
	chessBoard = append(chessBoard, *Knight(!white, left))
	chessBoard = append(chessBoard, *Knight(!white, !left))

	chessBoard = append(chessBoard, *Rock(white, left))
	chessBoard = append(chessBoard, *Rock(white, !left))
	chessBoard = append(chessBoard, *Rock(!white, left))
	chessBoard = append(chessBoard, *Rock(!white, !left))

	return chessBoard
}
