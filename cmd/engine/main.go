package main

import (
	"chessbit/pkg/attacks"
	b "chessbit/pkg/bitboard"
)

func main() {

	// init leaper pieces attack

	attacks.InitLeapersAttacks()

	// init occupancy bitboard
	var block uint64 = 0
	b.SetBit(&block, b.D7)
	b.SetBit(&block, b.D2)
	b.SetBit(&block, b.B4)
	b.SetBit(&block, b.G4)

	b.Print(block)

	b.Print(attacks.RookAttacksOnTheFly(b.D4, block))

}
