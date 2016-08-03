package models

// PieceType is the enum for the type of a piece
type PieceType int

// Enums for different piece types
const (
	Pawn PieceType = iota
	Rook
	Knight
	Bishop
	Queen
	King
)

// PieceNotations represents the letters each piece is represented by
var PieceNotations = [6]string{"", "R", "N", "B", "Q", "K"}

// Square represents the rank and file of a square
type Square struct {
	Rank int
	File int
}

// Piece represents a piece, independent of location
type Piece struct {
	Type  PieceType
	White bool
}

// Board represents all the pieces and their positions
type Board struct {
	Pieces [8][8]*Piece
}

// Game is a sequence of board states
type Game struct {
	BoardStates []*Board
}
