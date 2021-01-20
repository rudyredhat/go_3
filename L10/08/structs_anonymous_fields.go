// Sometimes it can be useful to have structs which contain one or more anonymous (or embedded) fields,
// that is fields with no explicit name. Only the type of such a field is mandatory and the type is then also
// its name. Such an anonymous field can also be itself a struct: structs can contain embedded structs.
// Go composition is of favoured over inheritance.
package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS // anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// with a struct-literal:
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is: ", outer2)
}

// outer.b is: 6
// outer.c is: 7.500000
// outer.int is: 60
// outer.in1 is: 5
// outer.in2 is: 10
// outer2 is:  {6 7.5 60 {5 10}}

// To store data in an anonymous field or get access to the data we use the name of the data type:
// outer.int; a consequence is that we can only have one anonymous field of each data type in a
// struct.
