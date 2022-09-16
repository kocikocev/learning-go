// Non empty is example of in-place slice algorithm that removes adjacent duplicates
package main

import "fmt"

// Non empty returns a slice holding only non empty strings
// The underlying array is modified during the call

func nonduplicate(strings []string) []string {
	if len(strings) == 1 || len(strings) == 0 {
		return strings
	}

	index := 1
	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[i-1] {
			strings[index] = strings[i]
			index++
		}
	}
	return strings[:index]
}

func main() {
	x := []string{"1", "\u4e16", "世"}
	x = nonduplicate(x)
	for _, v := range x {
		fmt.Println(v)
	}
}
