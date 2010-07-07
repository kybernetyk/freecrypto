// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto'.

package util

import (
	"testing"
)


func TestEntropy(t *testing.T) {
	bits := []uint16{80, 112, 128, 192, 256}
	result := []float64{10, 14, 16, 24, 32} // for domain of 256 values

	for i, b := range bits {
		if Entropy(b, 256) != result[i] {
			t.Errorf("error using %d bits\n", b)
		}
	}
}

