package attacks

// rook attacks table [square]

var RookAttacks [64]uint64

// generate rook attacks (relevant occupancy squares)

func MaskRookAttacks(square int) uint64 {
	// results attack bitboard
	var attacks uint64 = 0

	// init ranks & files
	var r, f int

	// init target rank & files
	tr := square / 8
	tf := square % 8

	// mask relevant rook occupancy bits

	for r = tr + 1; r <= 6; r++ {
		attacks |= (1 << (r*8 + tf))
	}

	for r = tr - 1; r >= 1; r-- {
		attacks |= (1 << (r*8 + tf))
	}
	for f = tf + 1; f <= 6; f++ {
		attacks |= (1 << (tr*8 + f))
	}
	for f = tf - 1; f >= 1; f-- {
		attacks |= (1 << (tr*8 + f))
	}

	return attacks

}
