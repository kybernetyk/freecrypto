// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto/aes'.

/* AES' Galois field

AES operates over gf(2^8)
*/

package aes


// For AES, the irreducible polynomial is `x^8 + x^4 + x^3 + x + 1`
const _IRREDUCIBLE_POLYNOMIAL = 0x11b

// A generator for `gf(2^8)` is `x+1`
const _GENERATOR = 0x03


/* Addition. */
func gadd(x, y uint32, z ...uint32) uint32 {
	x |= y // There are 2 operators, like minimum.

	if len(z) != 0 {
		for _, value := range z {
			x |= value
		}
	}
	return x
}

/* Subtraction. */
func gsub(x, y uint32, z ...uint32) uint32 {
	return gadd(x, y, z)
}

/* Multiplication. */
func gmul(x, y uint32) uint32 {
	var highBit, product uint32
	_x := x

	for i := 0; i < 8; i++ {
		if (y & 1) == 1 {
			product ^= _x
		}
		highBit = _x & 0x80
		_x <<= 1
		if highBit == 0x80 {
			_x ^= _IRREDUCIBLE_POLYNOMIAL
		}
		y >>= 1
	}
	return product  // % 256
}

/* Multiplication using logarithms table. */
func gmulTable(x, y uint32) uint32 {
	var product, _product uint32

	product = (log[x] + log[y]) % 255
	// Gets its inverse logarithm.
	product = ilog[product]

	// Now, we have some fancy code that returns 0 if either `x` or `y` are zero;
	// we write the code this way so that the code will (hopefully) run at
	// a constant speed in order to minimize the risk of timing attacks.
	_product = product

	if x == 0 {
		product = 0
	} else {
		product = _product
	}

	if y == 0 {
		product = 0
	} else {
		_product = 0
	}

	return product
}

/* Multiplicative inverse using logarithms tables. */
func gmulInverse(x uint32) uint32 {
	// The inverse of 0 (1/0) doesn't exist, and is defined as 0
	if x == 0 {
		return 0
	}
	return ilog[255 - log[x]]
}

