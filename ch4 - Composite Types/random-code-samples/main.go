package main

import (
	"fmt"
)

func main() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element

	// Print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// Initiate an array with array literals
	var q [3]int = [3]int{1, 2, 3}
	var p [3]int = [3]int{1, 2}
	fmt.Println(p[2]) // "0"
	fmt.Println(q[2]) // "3"

	// Simpler definition with array literals to determine size (type) of array
	r := [...]int{1, 2, 3}
	fmt.Printf("%T\n", r) // Type verb to show [3]int

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "\u20AC"}
	fmt.Println(EUR, symbol[EUR])

	// Defining array with 9 zeros and 1 -1
	t := [...]int{9: -1}

	for _, v := range t {
		fmt.Println(v)
	}

	// If an array element is comparable then the array is comparable as well
	ba := [2]int{1, 2}
	bb := [...]int{1, 2}
	bc := [2]int{1, 3}
	fmt.Println(ba == bb, bb == bc, bc == bb)
	fmt.Println(ba)
	s := []int{5, 6, 7, 8, 9}
	k := s[2:]

	for _, v := range k {
		fmt.Println(v)
	}
}
