// functions in Go are defined so that they return 2 values with successful execution: the value
// and true, and with unsuccessful execution: a 0 (or nil value) and false
// a program should test for every occurring error and behave accordingly
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC" // is taken as "123" the pass
	var an int
	var err error
	an, err = strconv.Atoi(orig) // str to int conv
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return // either we can return err to a calling function - val, err := pack1.Function1(param1)
		// prog to stop in case of err - os.Exit(1) - incase of return
	}
	fmt.Printf("The integer is %d\n", an)
}

// orig ABC is not an integer - exiting with error
