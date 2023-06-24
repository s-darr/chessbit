package attacks

import Bitboard "chessbit/pkg/bitboard"

/*******************************\
=================================

ATTACKS

=================================
\*********************************/

// pawn attacks table [side][square]

var PawnAttacks [2][64]uint64

func MaskPawnAttacks(side int, square int) uint64 {

	// results attack bitboard
	var attacks uint64 = 0

	// piece bitboard
	var bitboard uint64 = 0

	// set piece on bitboard
	Bitboard.SetBit(&bitboard, square)

	// white pawns or black (0 for white, 1 for black)
	if side == 0 {
		attacks |= bitboard >> 7
		// generate pawn attacks
		if (bitboard>>7)&NotAFile != 0 {
			attacks |= bitboard >> 7
		}
		if (bitboard>>9)&NotHFile != 0 {
			attacks |= bitboard >> 9
		}
	} else {
		if (bitboard<<7)&NotHFile != 0 {
			attacks |= bitboard << 7
		}
		if (bitboard<<9)&NotAFile != 0 {
			attacks |= bitboard << 9
		}

	}

	// return attack map
	return attacks

}

// initialise leaper pieces attack (does not need to go in pawn.go)

func InitLeapersAttacks() {
	// loop over 64 board squares
	for square := 0; square < 64; square++ {
		// init pawn attacks
		PawnAttacks[White][square] = MaskPawnAttacks(White, square)
		PawnAttacks[Black][square] = MaskPawnAttacks(Black, square)

		// init knight attacks
		KnightAttacks[square] = MaskKnightAttacks(square)

		// init king attacks
		KingAttacks[square] = MaskKingAttacks(square)
	}
}
