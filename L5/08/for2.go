// for loop in  the form of while
// initialize the val first
// iterate through the for with the val with a given condition
// and decrement the value to the condition
// then print the value, else if the decrement is at last then the -1 will not be printed
package main

import "fmt"

func main() {
	var i int = 5
	for i >= 0 {
		//		i = i - 1
		// will print the above from 4 to -1
		// else we can have
		// if i < 0{
		// print
		// break
		// }
		fmt.Printf("The variable i is now: %d\n", i)
		i = i - 1

	}
}

// The variable i is now: 5
// The variable i is now: 4
// The variable i is now: 3
// The variable i is now: 2
// The variable i is now: 1
// The variable i is now: 0
