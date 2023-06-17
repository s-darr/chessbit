package main

import (
	"chessbit/pkg/bitboard"
	"fmt"
)




func main() {
    
    var testBitboard uint64

    bitboard.SetBit(&testBitboard, bitboard.E2);
    bitboard.SetBit(&testBitboard, bitboard.C3);
    bitboard.SetBit(&testBitboard, bitboard.F2);
   
    a := bitboard.GetBit(testBitboard,52)
    fmt.Println(a)
    bitboard.PrintBitboard(testBitboard)



	

}