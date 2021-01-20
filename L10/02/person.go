// example Listing 10.2—person.go shows a struct Person, a method upPerson which
// has a parameter of type *Person (so that the object itself can be changed!) and 3 different ways of
// calling this method:
// struct as a val
// struct as a pointer
// struct as a literal
package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func main() {
	// 1- struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward" // this is also valid
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.
		lastName)

	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.
		lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.
		lastName)
}

// The name of the person is CHRIS WOODWARD
// The name of the person is CHRIS WOODWARD
// The name of the person is CHRIS WOODWARD
