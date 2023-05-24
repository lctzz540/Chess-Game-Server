package models

type Piece struct {
	Coordinate string
	VectorMove [][]int
	Name       string
	IsWhite    bool
}

func King(isWhite bool) *Piece {
	var coordinate string
	if isWhite {
		coordinate = "e5"
	} else {
		coordinate = "e8"
	}
	return &Piece{
		Coordinate: coordinate,
		VectorMove: [][]int{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}},
		Name:       "king",
		IsWhite:    isWhite,
	}
}

func Queen(isWhite bool) *Piece {
	var coordinate string
	if isWhite {
		coordinate = "d5"
	} else {
		coordinate = "d8"
	}
	return &Piece{
		Coordinate: coordinate,
		VectorMove: [][]int{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}},
		Name:       "queen",
		IsWhite:    isWhite,
	}
}
func Bishop(isWhite bool, isLeft bool) *Piece {
	var coordinate string
	if isWhite && isLeft {
		coordinate = "c1"
	} else {
		if isWhite && !isLeft {
			coordinate = "f1"
		} else {
			if !isWhite && isLeft {
				coordinate = "c8"
			} else {
				coordinate = "f8"
			}
		}
	}
	return &Piece{
		Coordinate: coordinate,
		VectorMove: [][]int{{-1, 1}, {1, 1}, {-1, -1}, {1, 1}},
		Name:       "bishop",
		IsWhite:    isWhite,
	}
}

func Knight(isWhite bool, isLeft bool) *Piece {
	var coordinate string
	if isWhite && isLeft {
		coordinate = "b1"
	} else {
		if isWhite && !isLeft {
			coordinate = "g1"
		} else {
			if !isWhite && isLeft {
				coordinate = "b8"
			} else {
				coordinate = "g8"
			}
		}
	}
	return &Piece{
		Coordinate: coordinate,
		VectorMove: [][]int{{-2, 1}, {1, -2}, {2, 1}, {2, -1}},
		Name:       "knight",
		IsWhite:    isWhite,
	}
}
func Rock(isWhite bool, isLeft bool) *Piece {
	var coordinate string
	if isWhite && isLeft {
		coordinate = "a1"
	} else {
		if isWhite && !isLeft {
			coordinate = "h1"
		} else {
			if !isWhite && isLeft {
				coordinate = "a8"
			} else {
				coordinate = "h8"
			}
		}
	}
	return &Piece{
		Coordinate: coordinate,
		VectorMove: [][]int{{1, 0}, {0, 1}, {-1, 0}, {-0, -1}},
		Name:       "rock",
		IsWhite:    isWhite,
	}
}
func Pawn(isWhite bool, coordinate string) *Piece {
	return &Piece{
		Coordinate: coordinate,
		VectorMove: [][]int{{1, 1}},
		Name:       "pawn",
		IsWhite:    isWhite,
	}
}
