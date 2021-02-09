// A method and the type on which it acts must be defined in the same package,
// that’s why you cannot define methods on type int, float or the like.
// Trying to define a method on an int type gives the compiler error:
// --- work around ---
// But there is a way around: you can define an alias for that type (int, float, ...), and then define
// a method for that type. Or embed the type as an anonymous type in a new struct

package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time //anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}
func main() {
	m := myTime{time.Now()}
	//calling existing String method on anonymous Time field
	fmt.Println("Full time now:", m.String())
	//calling myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())
}

// Full time now: 2021-02-08 11:59:26.078748028 +0530 IST m=+0.000139652
// First 3 chars: 202
