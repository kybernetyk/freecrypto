// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto/twofish'.

package twofish

import (
//"bytes"
)


/* Computes (c * x^4) mod (x^4 + (a + 1/a) * x^3 + a * x^2 + (a + 1/a) * x + 1)
over GF(256)
*/
func mod(c uint) uint {
	const MODULUS uint = 0x14d

	// c2 = (c << 1) ^ ((c & 0x80) ? MODULUS : 0)
	if c & 0x80 {
		v := MODULUS
	} else {
		v := 0
	}
	c2 := (c << 1) ^ v

	// c1 = c2 ^ (c >> 1) ^ ((c & 1) ? (MODULUS >> 1) : 0)
	if c & 1 {
		v = MODULUS >> 1
	} else {
		v = 0
	}
	c1 := c2 ^ (c >> 1) ^ v

	return c | (c1 << 8) | (c2 << 16) | (c1 << 24)
}

/* Compute `RS(12,8)` code with the above polynomial as generator.
This is equivalent to multiplying by the RS matrix.
*/
func reedSolomon(high, low uint32) uint32 {
	for i := 0; i < 8; i++ {
		high = mod(high>>24) ^ (high << 8) ^ (low >> 24)
		low <<= 8
	}
	return high
}

func h0(x, *key uint32, keyLen uint) uint32 {
	x = x | (x << 8) | (x << 16) | (x << 24)

	switch keyLen {
	case 4:
		x = Q(1, 0, 0, 1, x) ^ key[6]
	case 3:
		x = Q(1, 1, 0, 0, x) ^ key[4]
	case 2:
		x = Q(0, 1, 0, 1, x) ^ key[2]
		x = Q(0, 0, 1, 1, x) ^ key[0]
	}
	return x
}

// #define Q(a, b, c, d, t) q[a][GETBYTE(t,0)] ^ (q[b][GETBYTE(t,1)] << 8) ^ (q[c][GETBYTE(t,2)] << 16) ^ (q[d][GETBYTE(t,3)] << 24)
func Q(a, b, c, d, t byte) {
	q[a][GETBYTE(t,0)] ^ (q[b][GETBYTE(t,1)] << 8) ^ (q[c][GETBYTE(t,2)] << 16) ^ (q[d][GETBYTE(t,3)] << 24)
}

