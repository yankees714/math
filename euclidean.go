package main

import (
	"flag"
	"fmt"
)

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	aPtr := flag.Int("a", 1, "an int")
	bPtr := flag.Int("b", 1, "an int")

	flag.Parse()

	fmt.Println(gcd(*aPtr, *bPtr))
}
