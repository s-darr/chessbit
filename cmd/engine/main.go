package main

import (
	"chessbit/pkg/attacks"
	"fmt"
)

var state uint = 3850519659

func GetRandomNumber(state *uint) uint {
	*state ^= *state << 13
	*state ^= *state >> 17
	*state ^= *state << 5
	return *state
}

func main() {

	// init leaper pieces attack

	attacks.InitLeapersAttacks()
	// for each square count the relevant mask occupancy bits for bishop attaks
	//for rank := 0; rank < 8; rank++ {
	//	for file := 0; file < 8; file++ {
	//		square := rank*8 + file

	//		fmt.Printf("%d, ", bits.OnesCount64(attacks.MaskRookAttacks(square)))
	//	}
	//	fmt.Printf("\n")

	//}
	fmt.Printf("Before xorshift: %d\n", state)
	state := GetRandomNumber(&state)
	fmt.Printf("After xorshift: %d\n", state)

}
