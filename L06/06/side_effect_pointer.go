// passing a pointer to a func conserves memory
// no copy of val is made
// var and obj can be changed inside the func and does not have to be returned
// program, where reply, a pointer to an integer, is being changed in the function itself.
package main

import "fmt"

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Printf("Multiply:", *reply)
}

// Multiply:%!(EXTRA int=50)
