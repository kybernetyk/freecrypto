// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto/aes'.

/* AES' Galois field */

package aes


const _IRREDUCIBLE_POLYNOMIAL = 0x11b


/* Addition */
func gadd(x, y uint, z ...uint) uint {
	x |= y // There are 2 operators, like minimum.

	if len(z) != 0 {
		for _, value := range z {
			x |= value
		}
	}
	return x
}

/* Subtraction */
func gsub(x, y uint, z ...uint) uint {
	return gadd(x, y, z)
}

/* Multiplication */
func gmul(x, y byte) byte {
	var highBit, product byte
	_x := uint(x)

	for i := 0; i < 8; i++ {
		if (y & 1) == 1 {
			product ^= byte(_x)
		}
		highBit = byte(_x) & 0x80
		_x <<= 1
		if highBit == 0x80 {
			_x ^= _IRREDUCIBLE_POLYNOMIAL
		}
		y >>= 1
	}
	return product
}

