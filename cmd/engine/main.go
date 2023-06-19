package main

import (
	"chessbit/pkg/attacks"
	"chessbit/pkg/bitboard"
)

func main() {

	// init leaper pieces attack

	attacks.InitLeapersAttacks()

	//for square := 0; square < 64; square++ {
	//	bitboard.Print(attacks.MaskRookAttacks(square))
	var a uint64 = 0
	bitboard.Print(attacks.BishopAttacksOnTheFly(bitboard.D4, a))

}
