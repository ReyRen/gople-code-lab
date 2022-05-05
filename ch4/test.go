package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	ch1 := sha256.Sum256([]byte("x"))
	ch2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", ch1, ch2, ch1 == ch2, ch1)
}
