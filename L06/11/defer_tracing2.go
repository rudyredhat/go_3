package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	// function inside function
	// so inside function gets printed
	// and the outside fucntion is defered
	defer un(trace("a"))

	fmt.Println("in a")
}
func b() {

	defer un(trace("b"))

	fmt.Println("in b")

	a()
}
func main() {

	b()
}

// entering: b
// in b
// entering: a
// in a
// leaving: a
// leaving: b
