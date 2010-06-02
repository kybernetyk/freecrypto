// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto/aes'.

package aes

import (
	"testing"
)


func TestIrreduciblePolynomial(t *testing.T) {
	// x^8 + x^4 + x^3 + x + 1
	if gadd(1<<8, 1<<4, 1<<3, 1<<1, 1<<0) != _IRREDUCIBLE_POLYNOMIAL {
		t.Error("wrong value in constant for irreducible polynomial")
	}
}

func TestGenerator(t *testing.T) {
	// x + 1
	if gadd(1<<1, 1<<0) != _GENERATOR {
		t.Error("wrong value in constant for generator")
	}
}

