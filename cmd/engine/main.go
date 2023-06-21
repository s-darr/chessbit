package main

import (
	"chessbit/pkg/attacks"
	b "chessbit/pkg/bitboard"
	"fmt"
	"math/bits"
)

func main() {

	// init leaper pieces attack

	attacks.InitLeapersAttacks()

	var attackMask uint64 = attacks.MaskRookAttacks(b.A1)

    for x:=0; x<4096; x++{

        var occupancy uint64 = b.SetOccupancy(x,bits.OnesCount64(attackMask), attackMask)
        b.Print(occupancy)
        fmt.Scanln()
    }

    


}

