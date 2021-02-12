// ... type
// If the last parameter of a function is followed by ...type , this indicates that the function can deal
// with a variable number of parameters of that type, possibly also 0: a so called variadic function
// If the parameters are stored in an array arr, the function can be called with the parameter arr...
package main

import "fmt"

func main() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	// calling Min() with the ...
	x = Min(arr...)
	fmt.Printf("The minimum in the array arr is: %d\n", x)
}

// receive set of parameter using ...
func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

// The minimum is: 0
// The minimum in the array arr is: 1
