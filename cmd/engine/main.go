package main

import "chessbit/pkg/bitboard"




func main() {
    
    var testBitboard uint64


	testBitboard |= (1 << bitboard.E2)

    bitboard.PrintBitboard(testBitboard)



	
    

}