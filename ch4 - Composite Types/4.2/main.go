package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hashfun = flag.String("hashfunc", "SHA256", "Select hashing function: SHA256, SHA384, SHA512")

func main() {
	flag.Parse()
	fmt.Println(flag.Arg(0))
	if *hashfun == "SHA256" {
		fmt.Println(sha256.Sum256([]byte(flag.Arg(0))))
	} else if *hashfun == "SHA512" {
		fmt.Println(sha512.Sum512([]byte(flag.Arg(0))))
	} else {
		fmt.Println("Error: wrong sha")
	}
}
