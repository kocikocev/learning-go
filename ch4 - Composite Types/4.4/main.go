// Exercise 4.4: Write a version of rotate that operates in a single pass

package main

import "fmt"

func rotate(x []int, n int) []int {
	if n < 0 {
		n = n%len(x) + len(x)
	}
	rotatedSlice := make([]int, len(x))
	for i, v := range x {
		newIndex := (i + n) % len(x)
		rotatedSlice[newIndex] = v
	}
	return rotatedSlice
}

func main() {
	x := []int{1, 2, 3}
	x = rotate(x, 2)

	for i, v := range x {
		fmt.Printf("%d\t%d\n", i, v)
	}
}
