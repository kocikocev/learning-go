package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	x := "世Hello"
	var reversed []byte
	for i := 0; i < len(x); {
		r, size := utf8.DecodeRuneInString(x[i:])
		fmt.Printf("%d\t%d\n", i, r)
		i += size
		y := []byte(string(r))
		reverse(y)
		fmt.Println(y)
		for _, v := range y {
			reversed = append(reversed, v)
		}
	}

	reverse(reversed)
	fmt.Println(string(reversed))
}
