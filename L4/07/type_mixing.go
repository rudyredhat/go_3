// type mixing is not possible in go
// page 77 lisiting 4.8
package main
func main()  {
	var a int
	var b int32
	a = 15
	b = a + a // compile error
	b = b + 5
}