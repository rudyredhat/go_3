// HasPrefix tests whether the string s begins with prefix:
// strings.HasPrefix(s, prefix string) bool
// HasSuffix tests whether the string s end with suffix:
// strings.HasSuffix(s, suffix string) bool

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	// This illustrates also the use of the escape character \ to output a literal "with \", and the use of 2
	// substitutions in a format-string.
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
