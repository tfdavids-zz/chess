package models

// InitialBoardState returns the initial state of a chess game
func InitialBoardState() *Board {
	pieces := [8][8]*Piece{
		[8]*Piece{&Piece{Rook, true}, &Piece{Knight, true}, &Piece{Bishop, true}, &Piece{King, true}, &Piece{Queen, true}, &Piece{Bishop, true}, &Piece{Knight, true}, &Piece{Rook, true}},
		[8]*Piece{&Piece{Pawn, true}, &Piece{Pawn, true}, &Piece{Pawn, true}, &Piece{Pawn, true}, &Piece{Pawn, true}, &Piece{Pawn, true}, &Piece{Pawn, true}, &Piece{Pawn, true}},
		[8]*Piece{nil, nil, nil, nil, nil, nil, nil, nil},
		[8]*Piece{nil, nil, nil, nil, nil, nil, nil, nil},
		[8]*Piece{nil, nil, nil, nil, nil, nil, nil, nil},
		[8]*Piece{nil, nil, nil, nil, nil, nil, nil, nil},
		[8]*Piece{&Piece{Pawn, false}, &Piece{Pawn, false}, &Piece{Pawn, false}, &Piece{Pawn, false}, &Piece{Pawn, false}, &Piece{Pawn, false}, &Piece{Pawn, false}, &Piece{Pawn, false}},
		[8]*Piece{&Piece{Rook, false}, &Piece{Knight, false}, &Piece{Bishop, false}, &Piece{King, false}, &Piece{Queen, false}, &Piece{Bishop, false}, &Piece{Knight, false}, &Piece{Rook, false}},
	}

	b := Board{Pieces: pieces}

	return &b
}
