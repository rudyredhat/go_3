// strings.Fields(s) splits the string s around each instance of one or more consecutive white space characters

// strings.Split(s, sep) : works the same as Fields, but splits around a separator character or string sep (e.g.: ; or, or -).

// This results in a string with all the elements of the slice, separated by sep:
// Strings.Join(sl []string, sep string)
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hey my name is rudra"
	s1 := strings.Fields(str)
	fmt.Printf("Spilted in slice: %v\n", s1)
	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	// split with certain instructions
	str2 := "go1 | the abc fo go | 25"
	s12 := strings.Split(str2, "|")
	fmt.Printf("Spillted in slice: %v\n", s12)
	for _, val := range s12 {
		fmt.Printf("%s - ", val)
	}

	// join both the strings
	fmt.Println()
	str3 := strings.Join(s12, ";")
	fmt.Printf("s12 joined by ; => %s\n", str3)

}
