// ch 14 - mutexclevenodd.go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 6, odd(6))
	// 18 is odd: is false
}
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}
func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}
func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

// 16 is even: is true
// 17 is odd: is true
// 6 is odd: is false

// --------------------------------------------------------------------------------------------
// ch - 15 - function as parameteres to other func called as callback
// f1(1, f2)
package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
}
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

// check how it is called over here
func callback(y int, f func(int, int)) {
	f(y, 2)
	// this becomes Add(1, 2)
}

//The sum of 1 and 2 is: 3
// ------------------------------------------------------------------------------------------------
// ch - 16 
// closures -  sometimes not giving a name to a func make its anonymous func
// The first ( ) is the parameter-list and follows immediately after the keyword func because there is
// no function-name. The { } comprise the function body, and the last pair of ( ) represent the call of
// the function.
package main

import "fmt"

func main() {
	f()
}
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

// 0  - g is of type func(int) and has value 0x4ba440
// 1  - g is of type func(int) and has value 0x4ba440
// 2  - g is of type func(int) and has value 0x4ba440
// 3  - g is of type func(int) and has value 0x4ba440
// ------------------------------------------------------------------------------------------------
// ch - 17 - anonymous func with ret and defer
package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(f())
}

// 2 
// because of the ret++, which is executed after return 1.
// ------------------------------------------------------------------------------------------------
// ch - 18
// Lambda functions are also called closures, refer to var in surrounding func
// A closure is a function that captures some external state—for example, the state of the function in which it is created.
// Applying closures: a function returning another function
package main

import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add32()
	fmt.Printf("Call Add32 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}
func Add32() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int {

	return func(b int) int {

		return a + b
	}
}

// Call Add32 for 3 gives: 5
// The result is: 5
// ------------------------------------------------------------------------------------------------
// ch - 19
// (nearly) the same function used in a slightly different way:
package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

// 1 - 21 - 321
// ------------------------------------------------------------------------------------------------
// ch - 20
// Debugging with closures
// use of runtme and caller can be implemented in the closure fucntion
// ------------------------------------------------------------------------------------------------
// ch - 21
// timing a function
// start := time.Now() -- current
// longCalculation()
// end := time.Now()
// delta := end.Sub(start) -- .sub for the diff between them
// fmt.Printf("longCalculation took this amount of time: %s\n", delta)
// ------------------------------------------------------------------------------------------------
// ch - 22
// Instead cache the calculated value in memory, which is called memoization. eg - fibonacci
