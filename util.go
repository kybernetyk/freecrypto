// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto'.

package freecrypto

import (
	"encoding/binary"
)


// Turns from bits.

func BitToByte(n uint) uint {
	return n >> 3 // n / 8
}

func BitToWord16(n uint) uint {
	return n >> 4 // n / 16
}

func BitToWord32(n uint) uint {
	return n >> 5 // n / 32
}

func BitToWord64(n uint) uint {
	return n >> 6 // n / 64
}

// Turns to bits.

func ByteToBit(n uint) uint {
	return n << 3 // * 8
}


func IsPowerOf2(n uint) bool {
	return n > 0 && (n & (n-1)) == 0
}

