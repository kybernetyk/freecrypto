// To the extent possible under law, Authors have waived all copyright and
// related or neighboring rights to 'freecrypto'.

package rand

import (
	"fmt"
	"testing"
)


func TestChoice(t *testing.T) {
	var found bool

	chars := []byte{'a', 'b', 'c', '1', '2', '3'}
	dest := make([]byte, 16)
	Choice(dest, chars)

	for _, vDest := range dest {
		found = false

		for _, vChar := range chars {
			if vDest == vChar {
				found = true
				break
			}
		}

		if !found {
			t.Error("byte should be in list of characters:", string(vDest))
		}
	}

	fmt.Println("Characters to choice: " + string(chars))
	fmt.Println("Random choice: " + string(dest))
}

