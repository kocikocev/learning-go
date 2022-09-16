// Charcount computes counts of Unicode characters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[rune]int) // count of Unicode characters
	invalid := 0                 // count of invalid UTF-8 characters
	letters := 0                 // count of letters
	digits := 0                  // count of digits

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++

		if unicode.IsLetter(r) {
			letters++
		}

		if unicode.IsDigit(r) {
			digits++
		}

	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if letters > 0 {
		fmt.Printf("\n%d letters\n", letters)
	}

	if digits > 0 {
		fmt.Printf("\n%d digits\n", digits)
	}
}
