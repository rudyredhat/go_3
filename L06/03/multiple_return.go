// Named return variables
// takes 1 int parameter and ret 2 int param

package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}
func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

// take 1 input and ret 2 val - with unamed ret var
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// take 1 input and ret 2 val - with named ret var
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

// num = 10, 2x num = 20, 3x num = 30
// num = 10, 2x num = 20, 3x num = 30
