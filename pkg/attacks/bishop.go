package attacks

// Bishop attacks table [square]

var BishopAttacks [64]uint64

// generate bishop attacks (relevant occupancy squares)

func MaskBishopAttacks(square int) uint64 {
	// results attack bitboard
	var attacks uint64 = 0

	// init ranks & files
	var r, f int

	// init target rank & files
	tr := square / 8
	tf := square % 8

	// mask relevant bishop occupancy bits
	for r, f = tr+1, tf+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		attacks |= (1 << (r*8 + f))
	}
	for r, f = tr-1, tf+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		attacks |= (1 << (r*8 + f))
	}
	for r, f = tr+1, tf-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		attacks |= (1 << (r*8 + f))
	}
	for r, f = tr-1, tf-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		attacks |= (1 << (r*8 + f))
	}

	return attacks

}