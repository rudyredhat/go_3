// aliasing types
// a type can be given a another name, so that it can be used in the code
// page 84 listing 4.11
package main

import "fmt"

type TZ int

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d", c)
}
