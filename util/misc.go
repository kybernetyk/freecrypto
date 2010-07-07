// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto'.

package util

import (
	"math"
)


/* Calculates the entropy level according to the desired security level in bits
and the domain of possible values (the number of values).
*/
func Entropy(bitsLevel, values uint16) float64 {
	return float64(bitsLevel) / math.Log2(float64(values))
}

