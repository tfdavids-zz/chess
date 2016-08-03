package models

import "fmt"

// CanMove indicates whether the given piece is allowed to move from the given square
// to the given square.
func (b *Board) CanMove(p *Piece, from Square, to Square) (bool, error) {
	switch p.Type {
	case Pawn:
		return b.pawnCanMove(from, to)
	case Rook:
		return b.rookCanMove(from, to)
	case Knight:
		return b.knightCanMove(from, to)
	case Bishop:
		return b.bishopCanMove(from, to)
	case Queen:
		return b.queenCanMove(from, to)
	case King:
		return b.kingCanMove(from, to)
	default:
		return false, fmt.Errorf("Invalid piece type")
	}
}

func (b *Board) pawnCanMove(from Square, to Square) (bool, error) {
	// if square isn't on board return false
	// if move leaves friendly king in check return false
	// if square isn't one or two ahead of the pawn or to the side return false
	// if square is to the side and isn't a capture return false
	// if square is two ahead and pawn isn't on second rank return false
	// if square is blocked by friendly piece return false
	// if square is two ahead and blocked by piece return false
	// if square is on last rank and no promotion, return false
	return true, nil
}

func (b *Board) rookCanMove(from Square, to Square) (bool, error) {
	// if square isn't on board return false
	// if move leaves friendly king in check return false
	// if square isn't in same rank or file return false
	// if square is blocked by friendly piece return false
	// if squares in between are blocked by piece return false
	return true, nil
}

func (b *Board) knightCanMove(from Square, to Square) (bool, error) {
	// if square isn't on board return false
	// if move leaves friendly king in check return false
	// if square isn't a knight's move away return false
	// if square is blocked by a friendly piece return false
	return true, nil
}

func (b *Board) bishopCanMove(from Square, to Square) (bool, error) {
	// if square isn't on board return false
	// if move leaves friendly king in check return false
	// if square isn't in same diagonal return false
	// if square is blocked by friendly piece return false
	// if squares in between are blocked by piece return false
	return true, nil
}

func (b *Board) queenCanMove(from Square, to Square) (bool, error) {
	// if square isn't on board return false
	// if move leaves friendly king in check return false
	// if square isn't in same rank, file, or diagonal return false
	// if square is blocked by friendly piece return false
	// if squares in between are blocked by piece return false
	return true, nil
}

func (b *Board) kingCanMove(from Square, to Square) (bool, error) {
	// if square isn't on board return false
	// if square isn't adjacent return false
	// if square is blocked by friendly piece return false
	// if square is threatened by enemy piece return false
	return true, nil
}

// Next returns the next board state after the indicated move is made
func (b *Board) Next(from Square, to Square) (*Board, error) {
	next := Board{Pieces: b.Pieces}

	next.Pieces[to.File][to.Rank] = next.Pieces[from.File][from.Rank]
	next.Pieces[from.File][from.Rank] = nil

	return &next, nil
}
