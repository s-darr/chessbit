package bitboard

import (
	"fmt"
	"math/bits"
)

func Print(bitboard uint64) {
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

// function to get the index of the least significant bit
func LSBIndex(bitboard uint64) int {
	if bitboard == 0 {
		return -1
	}
	lsb := bitboard & -bitboard // trick to isolate the lsb
	return bits.TrailingZeros64(lsb)
}

// set occupancies
func SetOccupancy(index int, bitsInMask int, attackMask uint64) uint64 {

	// occupancy map
	var occupancy uint64 = 0

	// loop over the range of bits within attack mask
	for count:=0; count < bitsInMask; count++ {
		// Get ls1b index of attacks mask
		square  := LSBIndex(attackMask)

		// pop LS1B in attack map
		PopBit(&attackMask, square)

		// make sure occupancy is in board
		if (index & (1 << count) != 0){
			occupancy |= (1 << square)


		}

	}

	//return occupancy map
	return occupancy

}