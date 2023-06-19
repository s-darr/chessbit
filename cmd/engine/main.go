package main

import (
	"chessbit/pkg/attacks"
	b "chessbit/pkg/bitboard"
	"fmt"
)

func main() {

	// init leaper pieces attack

	attacks.InitLeapersAttacks()

	var block uint64 = 0
	b.SetBit(&block, b.D7)
	//b.SetBit(&block, b.D2)
	//b.SetBit(&block, b.B4)
	//b.SetBit(&block, b.G4)

	b.Print(block)
	fmt.Printf("index: %d    coordinate: %s\n", b.LSBIndex(block), b.SquareToCoordinate[b.LSBIndex((block))])

	//b.Print(attacks.RookAttacksOnTheFly(b.D4, block))

}
