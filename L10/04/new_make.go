// illustrate the difference in behavior for maps and the possible errors
package main
type Foo map[string]string //ch-8, map=key:value, make & new - 142
type Bar struct{
	thingOne string
	thingTwo int
}
func main() {
	// OK:
	y := new(Bar)
	(*y).thingOne = "hello" // struct as pointer
	(*y).thingTwo = 1
	// not OK:
	z := make(Bar) // compile error: cannot make type Bar
	z.thingOne = "hello"
	z.thingTwo = 1
	// OK:
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"
	// not OK:
	u := new(Foo)
	(*u)["x"] = "goodbye" // !! panic !!: runtime error: assignment to entry
	in nil map
	(*u)["y"] = "world"
}
// To try to make() a struct variable is not so bad, the compiler gets the error; but newing a map and
// trying to fill it with data gives a runtime error! new(Foo) is a pointer to a nil, not yet allocated,
// map: so be very cautious with this!