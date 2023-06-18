package main

import (
	"chessbit/pkg/attacks"
	"chessbit/pkg/bitboard"
)

func main() {

	// init leaper pieces attack

	attacks.InitLeapersAttacks()

	for square := 0; square < 64; square++ {
		bitboard.Print(attacks.KingAttacks[square])
	}

}
