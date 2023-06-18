package main

import (
	"chessbit/pkg/bitboard"
)




func main() {

    for rank := 0; rank < 8; rank++ {
        for file:=0; file < 8; file++ {
            square := rank * 8 + file

            if file != 0 {
                bitboard.SetBit(&bitboard.NotAFile, square)

            }
            
        }
    }
    
   // bitboard.PrintBitboard(bitboard.MaskPawnAttacks(bitboard.E4,bitboard.White))
    bitboard.PrintBitboard(bitboard.NotAFile)
    



	

}