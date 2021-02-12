// Index returns the index of the first instance of str in s, or -1 if str is not present in s:
// strings.Index(s, str string) int
// LastIndex returns the index of the last instance of str in s, or -1 if str is not present in s:
// strings.LastIndex(s, str string) int
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi I'm Rudra, Hi."
	fmt.Printf("The position of \"Rudra\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Rudra"))
	fmt.Printf("The position of first \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of last \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))
	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
}

// The position of "Rudra" is: 7
// The position of first "Hi" is: 0
// The position of last "Hi" is: 14
// The position of "Burger" is: -1
