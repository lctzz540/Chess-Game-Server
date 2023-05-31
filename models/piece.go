package models

import (
	"strconv"
)

type Piece struct {
	initSquare string
	VectorMove [][]int
	Name       string
	IsWhite    bool
	Coordinate []int
}

func SquareToCoordinate(square string) []int {
	var coordinate []int
	var row = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for index, value := range row {
		if value == string(square[0]) {
			coordinate = append(coordinate, index)
		}
	}
	num, _ := strconv.Atoi(string(square[1]))
	coordinate = append(coordinate, num-1)
	return coordinate
}

func King(isWhite bool) *Piece {
	var square string
	if isWhite {
		square = "e5"
	} else {
		square = "e8"
	}
	coordinate := SquareToCoordinate(square)

	return &Piece{
		initSquare: square,
		VectorMove: [][]int{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}},
		Name:       "king",
		IsWhite:    isWhite,
		Coordinate: coordinate,
	}
}

func Queen(isWhite bool) *Piece {
	var square string
	if isWhite {
		square = "d5"
	} else {
		square = "d8"
	}
	coordinate := SquareToCoordinate(square)
	return &Piece{
		initSquare: square,
		VectorMove: [][]int{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}},
		Name:       "queen",
		IsWhite:    isWhite,
		Coordinate: coordinate,
	}
}
func Bishop(isWhite bool, isLeft bool) *Piece {
	var square string
	if isWhite && isLeft {
		square = "c1"
	} else {
		if isWhite && !isLeft {
			square = "f1"
		} else {
			if !isWhite && isLeft {
				square = "c8"
			} else {
				square = "f8"
			}
		}
	}
	coordinate := SquareToCoordinate(square)
	return &Piece{
		initSquare: square,
		VectorMove: [][]int{{-1, 1}, {1, 1}, {-1, -1}, {1, 1}},
		Name:       "bishop",
		IsWhite:    isWhite,
		Coordinate: coordinate,
	}
}

func Knight(isWhite bool, isLeft bool) *Piece {
	var square string
	if isWhite && isLeft {
		square = "b1"
	} else {
		if isWhite && !isLeft {
			square = "g1"
		} else {
			if !isWhite && isLeft {
				square = "b8"
			} else {
				square = "g8"
			}
		}
	}
	coordinate := SquareToCoordinate(square)
	return &Piece{
		initSquare: square,
		VectorMove: [][]int{{-2, 1}, {1, -2}, {2, 1}, {2, -1}},
		Name:       "knight",
		IsWhite:    isWhite,
		Coordinate: coordinate,
	}
}
func Rock(isWhite bool, isLeft bool) *Piece {
	var square string
	if isWhite && isLeft {
		square = "a1"
	} else {
		if isWhite && !isLeft {
			square = "h1"
		} else {
			if !isWhite && isLeft {
				square = "a8"
			} else {
				square = "h8"
			}
		}
	}
	coordinate := SquareToCoordinate(square)
	return &Piece{
		initSquare: square,
		VectorMove: [][]int{{1, 0}, {0, 1}, {-1, 0}, {-0, -1}},
		Name:       "rock",
		IsWhite:    isWhite,
		Coordinate: coordinate,
	}
}
func Pawn(isWhite bool, square string) *Piece {
	coordinate := SquareToCoordinate(square)
	return &Piece{
		initSquare: square,
		VectorMove: [][]int{{1, 1}},
		Name:       "pawn",
		IsWhite:    isWhite,
		Coordinate: coordinate,
	}
}
