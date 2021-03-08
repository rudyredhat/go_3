// interfaces are set of methods
// does not contain any code
// cannot contain var
// naming = name+er , eg Logger, Reader etc
// contains 0-3 methods max
// 7.2.4 - new() and make() diff
package main

import "fmt"

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

// a method which has a parameter of type *Square (so that the object itself can be changed!) - P 228
// P-240, last 2 para and P-241 - method explained
// P-241, code 10.10 - eg for the method
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func main() {
	sq1 := new(Square)
	sq1.side = 5
	areaIntf := sq1 //polymorphism
	fmt.Printf("the square has area: %f\n", areaIntf.Area())
}

//the square has area: 25.000000
