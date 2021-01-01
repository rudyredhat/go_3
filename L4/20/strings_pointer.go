// assigning a new value to *p changes the value of the variable itself (here a string).
package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  //prints addr
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}

// Here is the pointer p: 0xc000010200
// Here is the string *p: ciao
// Here is the string s: ciao
