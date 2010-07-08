// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto'.

package util

import (
	"fmt"
	"testing"
)

var bits = []uint{128, 192, 256}


func TestEntropy(t *testing.T) {
	result := []float64{16, 24, 32} // for domain of 256 values

	for i, b := range bits {
		if Entropy(b, 256) != result[i] {
			t.Errorf("error using %d bits\n", b)
		}
	}
}

func TestShowEntropy(t *testing.T) {
	var domain = []byte(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-")

	fmt.Printf("Entropy for a domain of %d values:\n\n", len(domain))

	for _, b := range bits {
		fmt.Printf("%v (bits) = %v (chars.)\n",
			b, Entropy(b, uint(len(domain))))
	}
	println()
}

