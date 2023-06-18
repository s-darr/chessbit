package main

import (
	"chessbit/pkg/bitboard"
)

func main() {

	// init leaper pieces attack

	bitboard.InitLeapersAttacks()

	bitboard.Print(bitboard.MaskPawnAttacks(bitboard.Black, bitboard.H4))

}
