// addr are stored in special data type called as pointer
// var intP *int
// Then the following is true: intP = &i1 , intP points to i1.
// So for any variable var the following is true:	var == *(&var)
package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int //declare pointer
	intP = &i1    //accessing the memory add
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}

// An integer: 5, its location in memory: 0xc000014100
// The value at memory location 0xc000014100 is 5
