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
			fmt.Printf(" %d", (bitboard>>square)&1)
		}

		// Print new line every rank
		fmt.Println()
	}

	// Print board files
	fmt.Println("\n     a b c d e f g h")

	// Print bitboard as unsigned decimal number
	fmt.Printf("\n     Bitboard: %d\n\n", bitboard)
}