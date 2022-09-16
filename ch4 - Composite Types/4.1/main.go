package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	result := BitDiffenceCountSha256(sha256.Sum256([]byte("X")), sha256.Sum256([]byte("x")))
	fmt.Println(result)
}

func PopCount(x uint64) int {
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func BitDiffenceCountSha256(a [32]byte, b [32]byte) int {
	var result int
	for i := 0; i < 32; i++ {
		// Using bitwise formula for exclusive OR
		c := (a[i] | b[i]) & ^(a[i] & b[i])
		result += PopCount(uint64(c))
	}
	return result
}
