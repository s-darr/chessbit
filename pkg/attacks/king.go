package attacks

import Bitboard "chessbit/pkg/bitboard"

// king attacks table [square]

var KingAttacks [64]uint64

// generate king attacks

func MaskKingAttacks(square int) uint64 {
	// results attack bitboard
	var attacks uint64 = 0

	// piece bitboard
	var bitboard uint64 = 0

	// set piece on bitboard
	Bitboard.SetBit(&bitboard, square)

	// generate king attacks

	if (bitboard >> 8) != 0 {
		attacks |= bitboard >> 8
	}
	if (bitboard>>9)&NotHFile != 0 {
		attacks |= bitboard >> 9
	}
	if (bitboard>>7)&NotAFile != 0 {
		attacks |= bitboard >> 7
	}
	if (bitboard>>1)&NotHFile != 0 {
		attacks |= bitboard >> 1
	}

	if (bitboard << 8) != 0 {
		attacks |= bitboard << 8
	}
	if (bitboard<<9)&NotAFile != 0 {
		attacks |= bitboard << 9
	}
	if (bitboard<<7)&NotHFile != 0 {
		attacks |= bitboard << 7
	}
	if (bitboard<<1)&NotAFile != 0 {
		attacks |= bitboard << 1
	}

	return attacks

}
