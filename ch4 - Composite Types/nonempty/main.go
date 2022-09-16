// Non empty is example of in-place slice algorithm
package main

// Non empty returns a slice holding only non empty strings
// The underlying array is modified during the call

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
