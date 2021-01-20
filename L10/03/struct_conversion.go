// When we have a struct type and
// define an alias type for it, both types have the same underlying type and can be converted into
// one another as illustrated in Listing 10.3, but also note the compile-error cases which denote
// impossible assignments or conversions:
package main

import "fmt"

type number struct {
	f float32
}
type nr number // alias type
func main() {
	a := number{5.0}
	b := nr{5.0}
	// var i float32 = b compile-error: cannot use b (type nr) as type float32 in assignment

	// var i = float32(b) compile-error: cannot convert b (type nr) to type float32

	// var c number = b compile-error: cannot use b (type nr) as type number in assignment

	// needs a conversion:

	var c = number(b)

	fmt.Println(a, b, c)

}

//{5} {5} {5}
