package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		wordCount[input.Text()]++
	}

	for word, n := range wordCount {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, word)
		}
	}
}
