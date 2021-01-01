// In this program we disregard the possible conversion-error with the blank identifier _:
// anInt, _ = strconv.Atoi(origStr)
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// str to int
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The inetger is: %d\n", an)
	an = an + 5
	// int to str
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}

// The size of ints is: 64
// The inetger is: 666
// The new string is: 671
