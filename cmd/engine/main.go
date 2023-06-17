package main

import (
	"chessbit/pkg/bitboard"
)




func main() {
    
    var testBitboard uint64

    bitboard.SetBit(&testBitboard, bitboard.E2);
    bitboard.SetBit(&testBitboard, bitboard.C3);
    bitboard.SetBit(&testBitboard, bitboard.F2);
    bitboard.PrintBitboard(testBitboard)
   
    bitboard.PopBit(&testBitboard, bitboard.E2);

    
    bitboard.PrintBitboard(testBitboard)



	

}