package bitboard

import "fmt"





func PrintBitboard(bitboard uint64) {
	fmt.Println()

	// Loop over board ranks
	for rank := 0; rank < 8; rank++ {
		// Loop over board files
		for file := 0; file < 8; file++ {
			// Convert file & rank into square index
			square := rank*8 + file

			// Print ranks
			if file == 0 {
				fmt.Printf("  %d ", 8-rank)
			}

			// Print bit state (either 1 or 0)
			if GetBit(bitboard, square) != 0 {
				fmt.Printf(" %d", 1)
			} else {
				fmt.Printf(" %d", 0)
			}
		} 

		// Print new line every rank
		fmt.Println()
	}

	// Print board files
	fmt.Println("\n     a b c d e f g h")

	// Print bitboard as unsigned decimal number
	fmt.Printf("\n     Bitboard: %d\n\n", bitboard)
}

func SetBit(bitboard *uint64, square int) {
	*bitboard |= (1 << square)

}

func GetBit(bitboard uint64, square int) uint64 {
	return bitboard & (1 << square)
}

func PopBit(bitboard *uint64, square int) {
    if GetBit(*bitboard, square) != 0 {
        *bitboard ^= (1 << square)
    }
}

/*******************************\
=================================

ATTACKS

=================================
\*********************************/

// not a file constant

var NotAFile uint64 = 0




// pawn attacks table [side][square]

var pawnAttacks [2][64] uint64

func MaskPawnAttacks(square int, side int) uint64 {

	// results attack bitboard
	var attacks uint64 = 0

	// piece bitboard
	var bitboard uint64 = 0

	// set piece on board
	SetBit(&bitboard,square)
	

	// white pawns or black
	if side == 1 {
		attacks |= bitboard >> 7

	}
	


	// return attack map
	return attacks



}

