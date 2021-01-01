// Repeat returns a new string consisting of count copies of the string s: int) string
// repeat the set of strings, with how many times you want to repeat
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = " Hi there"
	var str1 string
	str1 = strings.Repeat(str, 3)
	fmt.Printf("The new repeated string is: %s", str1)
}
