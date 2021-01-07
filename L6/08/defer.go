// defer keyword allows us to postpone the execution of a statement or a function until the end
// of the enclosing (calling) function
// defer - generally all the current function work is done and then the defer fucn will execute
// It can be helpful to keep the code cleaner and so often shorter.
package main

import "fmt"

func main() {
	Function1()
}
func Function1() {
	fmt.Printf("In F1\n")
	defer Function2()
	fmt.Printf("In F1 end\n")
}
func Function2() {
	fmt.Printf("Defered until the end of calling func")
}

// In F1
// In F1 end
// Defered until the end of calling func
