package attacks

import Bitboard "chessbit/pkg/bitboard"

/*******************************\
=================================

ATTACKS

=================================
\*********************************/

/*

    not A file
8  0 1 1 1 1 1 1 1
7  0 1 1 1 1 1 1 1
6  0 1 1 1 1 1 1 1
5  0 1 1 1 1 1 1 1
4  0 1 1 1 1 1 1 1
3  0 1 1 1 1 1 1 1
2  0 1 1 1 1 1 1 1
1  0 1 1 1 1 1 1 1

a b c d e f g h
\*********************************/

// not A file constant (decimal)
const NotAFile uint64 = 18374403900871474942

// not H file constant
const NotHFile uint64 = 9187201950435737471

// not H or G file constant
const NotHGFile uint64 = 4557430888798830399

// not A or B file constant
const NotABFile uint64 = 18229723555195321596

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
