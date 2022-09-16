package main

import (
	"fmt"
	"sort"
)

func main() {
	// using make to create a map
	ages := make(map[string]int) // mapping from strings to ints

	fmt.Println(ages)

	// literal definition of map
	agesLiteral := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	fmt.Println(agesLiteral)

	// Initialize equivalent to literal initialization
	agesInitialize := make(map[string]int)
	agesInitialize["alice"] = 31
	agesInitialize["charlie"] = 34

	fmt.Println(agesInitialize)

	// accessing through subscript notation
	ages["alice"] = 32
	fmt.Println(ages["alice"])

	// delete with built in functions
	delete(ages, "alice")

	// Operations are safe even if the element isn't in the map
	fmt.Println(ages["bob"])

	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages["bob"])

	ages["bob"] += 1
	ages["bob"]++

	fmt.Println(ages["bob"])

	// Map element is not a variable, and we cannot take its address
	// _ = &ages["bob"]

	// Iterating - NOTE: order is not guaranteed

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// The zero value for a maps is nil, that is, a reference to no hash table at all
	var agesZero map[string]int
	fmt.Println(agesZero == nil) // "true"
	fmt.Println(len(agesZero) == 0)

	// delete, len and range are safe to perform on a nil map, but
	// storing to a nil map causes panic

	// Panic --> agesZero["carol"] = 31

	// Subscripting yields a value, but to know if the element was present you do:
	ageValue, ok := ages["kate"]

	if !ok {
		fmt.Println(ageValue)
		/* "bob" is not a key in this map; */
	}

	// Often combined as
	if ageValue, ok := ages["kate"]; !ok {
		fmt.Println(ageValue)
	}

	// Maps can't be compared to each other, they can only be compared to nil
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))

	// If we need a set of values who cannot be a key for the map, because == does not work,
	// we can map them to strings as long as strings are equal if and only if values being mapped are equal
	var m = make(map[string]int)

	list := []string{"yo", "ho", "ho"}
	m[k(list)]++
	fmt.Println(k(list))
	fmt.Println(m)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

// Mapping slice to string
func k(list []string) string { return fmt.Sprintf("%q", list) }
