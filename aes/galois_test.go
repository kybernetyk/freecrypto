// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto/aes'.

package aes

import (
	"testing"
)


/* For the AES algorith, the irreducible polynomial is
    `x^8 + x^4 + x^3 + x + 1`
*/
func TestIrreduciblePolynomial(t *testing.T) {
	if _IRREDUCIBLE_POLYNOMIAL != gadd(1<<8, 1<<4, 1<<3, 1<<1, 1<<0) {
		t.Error("wrong value in irreducible polynomial")
	}
}

