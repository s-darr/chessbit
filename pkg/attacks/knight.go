package attacks

import Bitboard "chessbit/pkg/bitboard"

// knight attacks table [square]

var KnightAttacks [64]uint64

// generate knight attacks

func MaskKnightAttacks(square int) uint64 {
	// results attack bitboard
	var attacks uint64 = 0

	// piece bitboard
	var bitboard uint64 = 0

	// set piece on bitboard
	Bitboard.SetBit(&bitboard, square)

	// generate knight attacks

	if (bitboard>>17)&NotHFile != 0 {
		attacks |= bitboard >> 17
	}
	if (bitboard>>15)&NotAFile != 0 {
		attacks |= bitboard >> 15
	}
	if (bitboard>>10)&NotHGFile != 0 {
		attacks |= bitboard >> 10
	}
	if (bitboard>>6)&NotABFile != 0 {
		attacks |= bitboard >> 6
	}
	if (bitboard<<17)&NotAFile != 0 {
		attacks |= bitboard << 17
	}
	if (bitboard<<15)&NotHFile != 0 {
		attacks |= bitboard << 15
	}
	if (bitboard<<10)&NotABFile != 0 {
		attacks |= bitboard << 10
	}
	if (bitboard<<6)&NotHGFile != 0 {
		attacks |= bitboard << 6
	}

	return attacks

}
