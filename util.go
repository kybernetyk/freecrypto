// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto'.

package freecrypto


// Turns from bits to bytes of different sizes.

func BitsToBytes(n uint) uint {
	return n >> 3 // n / 8
}

func BitsToWords16(n uint) uint {
	return n >> 4 // n / 16
}

func BitsToWords32(n uint) uint {
	return n >> 5 // n / 32
}

func BitsToWords64(n uint) uint {
	return n >> 6 // n / 64
}

// Turns from bytes of different sizes to bits.

func BytesToBits(n uint) uint {
	return n << 3 // * 8
}

