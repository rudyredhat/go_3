// A nil pointer dereference, like in the following 2 lines (see program testcrash.go), is illegal and
// makes a program crash:

package main

func main() {
	var p *int = nil
	*p = 0
}
