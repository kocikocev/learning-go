// Non empty is example of in-place slice algorithm that removes adjacent duplicates
package main

import (
	"fmt"
	"unicode"
)

// Non empty returns a slice holding only non empty strings
// The underlying array is modified during the call

func nonduplicate(strings []byte) []byte {
	index := 0
	for i := 0; i < len(strings); i++ {
		if (i == 0 || strings[index-1] == ' ') && unicode.IsSpace(rune(strings[i])) {
			continue
		} else {
			if unicode.IsSpace(rune(strings[i])) {
				strings[index] = ' '
			} else {
				strings[index] = strings[i]
			}

			index++
		}
	}
	return strings[:index]

}

func main() {
	x := []byte{'1', '\t', '\t', 2}
	x = nonduplicate(x)
	for _, v := range x {
		fmt.Println(v)
	}
}
