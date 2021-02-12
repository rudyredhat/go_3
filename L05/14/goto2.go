// use of goto and showing incorrect way to define var inside goto and labels
// use only forward labels
package main

import "fmt"

func main() {
	a := 1
	goto TARGET // compile error:
	// goto TARGET jumps over declaration of b at goto2.go:8
	b := 9
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
