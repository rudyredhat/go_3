// program simple_function.go the function takes 3 int parameters a,b and c and returns an int
// call by val - orig val not changed
// call by add - orig val is changed - pass the memory add
package main

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	// var i1 int = MultiPly3Nums(2, 5, 6)
	// fmt.Printf("Multiply 2 * 5 * 6 = %d\n", i1)
}
func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

//Multiply 2 * 5 * 6 = 60
